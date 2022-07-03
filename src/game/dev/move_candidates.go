package dev

import "math/rand"

import "github.com/tinne26/bindless/src/game/iso"

// Movement for magnets is quite tricky. They may move in four directions,
// but we have to prioritize closer sources of magnetism. There can be
// multiple candidate moves, and there can be lots of collisions and
// conflicts. Solving this nicely is a massive pain. So I don't.

// TODO: as exemplified in the door level, when one magnet is pulling and
//       another is pushing, the forces are not being added and the
//       simulation is not being faithful.

type CandidateMovesPack struct {
	Magnet *FloatMagnet
	NE int16 // north east move magnetism source distance, 0 if none
	NW int16 // north west...
	SE int16 // south east...
	SW int16 // south west...
}

// TODO: problem. if we end up moving nowhere, we may not be able to
//       solve the system, because we pre-allowed a move that may not
//       be viable anymore. I can do chain cancelling though, that does
//       make sense, seems viable.

// Returns true if other moves are still possible, or false if no
// moves remain after disabling the given one.
func (self *CandidateMovesPack) DisableMove(targetCol, targetRow int16) bool {
	col, row := self.Magnet.Column(), self.Magnet.Row()
	if absInt16(targetCol - col) + absInt16(targetRow - row) != 1 {
		return true // manhattan distance indicates that the positions are too far away
	}

	if self.NE > 0 {
		neCol, neRow := self.TileNE()
		if neCol == targetCol && neRow == targetRow {
			self.NE = 0
			return !self.Empty()
		}
	}
	if self.NW > 0 {
		nwCol, nwRow := self.TileNW()
		if nwCol == targetCol && nwRow == targetRow {
			self.NW = 0
			return !self.Empty()
		}
	}
	if self.SE > 0 {
		seCol, seRow := self.TileSE()
		if seCol == targetCol && seRow == targetRow {
			self.SE = 0
			return !self.Empty()
		}
	}
	if self.SW > 0 {
		swCol, swRow := self.TileSW()
		if swCol == targetCol && swRow == targetRow {
			self.SW = 0
			return !self.Empty()
		}
	}

	return true
}

func (self *CandidateMovesPack) Empty() bool {
	return self.NE + self.NW + self.SE + self.SW == 0
}

// must be called before applying inertia
func (self *CandidateMovesPack) CanUnlock() bool {
	return self.NE == 1 || self.NW == 1 || self.SE == 1 || self.SW == 1
}

func (self *CandidateMovesPack) TileNE() (int16, int16) {
	return self.Magnet.Column() + 1, self.Magnet.Row()
}

func (self *CandidateMovesPack) TileNW() (int16, int16) {
	return self.Magnet.Column(), self.Magnet.Row() - 1
}

func (self *CandidateMovesPack) TileSE() (int16, int16) {
	return self.Magnet.Column(), self.Magnet.Row() + 1
}

func (self *CandidateMovesPack) TileSW() (int16, int16) {
	return self.Magnet.Column() - 1, self.Magnet.Row()
}

func (self *CandidateMovesPack) CorrectOpposing() {
	if self.NE > 0 && self.SW > 0 {
		if self.NE == self.SW {
			self.NE = 0
			self.SW = 0
		} else if self.NE > self.SW {
			self.SW += 1
			self.NE = 0
		} else {
			self.NE += 1
			self.SW = 0
		}
	}
	if self.NW > 0 && self.SE > 0 {
		if self.NW == self.SE {
			self.NW = 0
			self.SE = 0
		} else if self.NW > self.SE {
			self.SE += 1
			self.NW  = 0
		} else {
			self.NW += 1
			self.SE = 0
		}
	}
}

// what originally were distances have now been hacked to also
// take into account the previous inertia of the magnet... absurd
func (self *CandidateMovesPack) ApplyInertia(state FloatMagnetState) {
	self.NE *= 3
	self.NW *= 3
	self.SE *= 3
	self.SW *= 3

	switch state {
	case StMoveNE:
		if self.NE != 0 { self.NE -= 2 }
		if self.SE != 0 { self.SE -= 1 }
		if self.NW != 0 { self.NW -= 1 }
	case StMoveNW:
		if self.NW != 0 { self.NW -= 2 }
		if self.NE != 0 { self.NE -= 1 }
		if self.SW != 0 { self.SW -= 1 }
	case StMoveSE:
		if self.SE != 0 { self.SE -= 2 }
		if self.NE != 0 { self.NE -= 1 }
		if self.SW != 0 { self.SW -= 1 }
	case StMoveSW:
		if self.SW != 0 { self.SW -= 2 }
		if self.NW != 0 { self.NW -= 1 }
		if self.SE != 0 { self.SE -= 1 }
	default:
		self.NE *= 3
		self.NW *= 3
		self.SE *= 3
		self.SW *= 3
	}
}

type MoveChoice struct {
	Magnet *FloatMagnet
	TargetColumn int16
	TargetRow int16
}

func SolveCandidateMovesSystem(unsolvedPacks []*CandidateMovesPack, magnets iso.Map[Magnet]) []MoveChoice {


	// shuffle unsolvedPacks
	rand.Shuffle(len(unsolvedPacks), func(i, j int) {
		unsolvedPacks[i], unsolvedPacks[j] = unsolvedPacks[j], unsolvedPacks[i]
	})

	// solution choices (may need to undo some at the
	// end if conflicts still happen)
	solutionChoices := make([]MoveChoice, 0, len(unsolvedPacks))

	// pre-filter invalid moves
	magnets.Each(
		func (col, row int16, magnet Magnet) {
			floatMagnet, isFloatMagnet := magnet.(*FloatMagnet)
			if isFloatMagnet {
				foundInUnsolved := false
				for _, pack := range unsolvedPacks {
					if pack.Magnet == floatMagnet {
						foundInUnsolved = true
						break
					}
				}
				if !foundInUnsolved {
					solutionChoices, unsolvedPacks = disableMoveInSystem(col, row, solutionChoices, unsolvedPacks)
				}
			}
		})

	// start solving candidate move packs from distance = 1
	distance := int16(1)
	repeats := 0
	for len(unsolvedPacks) > 0 {
		repeats += 1
		if repeats == 100 { panic("infinite loop") }

		// find solution for this distance level
	   // (at most 4 magnets can try to move to the same place at once)
		var targetPack *CandidateMovesPack
		var targetCol, targetRow int16
		for _, pack := range unsolvedPacks {
			// find first pack at this distance
			// TODO: checking the directions in a fixed order causes
			//       predictable predilections that are incorrect
			if pack.NW == distance {
				targetPack = pack
				targetCol, targetRow = pack.TileNW()
			} else if pack.SE == distance {
				targetPack = pack
				targetCol, targetRow = pack.TileSE()
			} else if pack.NE == distance {
				targetPack = pack
				targetCol, targetRow = pack.TileNE()
			} else if pack.SW == distance {
				targetPack = pack
				targetCol, targetRow = pack.TileSW()
			}
		}

		// if no target found at this distance, jump to next one
		if targetPack == nil { distance += 1 ; continue }

		// remove the pack from the unsolved list
		solutionChoices = append(solutionChoices, MoveChoice{ targetPack.Magnet, targetCol, targetRow })
		unsolvedPacks = removeElemFromSlice(targetPack, unsolvedPacks)

		// remove the option for other potentially colliding packs
		solutionChoices, unsolvedPacks = disableMoveInSystem(targetCol, targetRow, solutionChoices, unsolvedPacks)
		solutionChoices, unsolvedPacks = disableMoveFromToInSystem(targetCol, targetRow, targetPack.Magnet.Column(), targetPack.Magnet.Row(), solutionChoices, unsolvedPacks)
	}

	return solutionChoices
}

func disableMoveInSystem(targetCol, targetRow int16, solutionChoices []MoveChoice, unsolvedPacks []*CandidateMovesPack) ([]MoveChoice, []*CandidateMovesPack) {
	for _, pack := range unsolvedPacks {
		hasMovesLeft := pack.DisableMove(targetCol, targetRow)
		if !hasMovesLeft {
			unsolvedPacks = removeElemFromSlice(pack, unsolvedPacks)
			removedColumn, removedRow := pack.Magnet.Column(), pack.Magnet.Row()

			for i, choice := range solutionChoices {
				if choice.TargetColumn == removedColumn && choice.TargetRow == removedRow {
					// have to invalidate this solution
					solutionChoices[i] = solutionChoices[len(solutionChoices) - 1]
					solutionChoices = solutionChoices[0 : len(solutionChoices) - 1]
					removedColumn, removedRow = choice.Magnet.Column(), choice.Magnet.Row()
					solutionChoices, unsolvedPacks = disableMoveInSystem(removedColumn, removedRow, solutionChoices, unsolvedPacks)
					break
				}
			}
		}
	}
	return solutionChoices, unsolvedPacks
}

func disableMoveFromToInSystem(fromCol, fromRow, toCol, toRow int16, solutionChoices []MoveChoice, unsolvedPacks []*CandidateMovesPack) ([]MoveChoice, []*CandidateMovesPack) {
	for _, pack := range unsolvedPacks {
		if pack.Magnet.Column() == fromCol && pack.Magnet.Row() == fromRow {
			hasMovesLeft := pack.DisableMove(toCol, toRow)
			if !hasMovesLeft {
				unsolvedPacks = removeElemFromSlice(pack, unsolvedPacks)
				removedColumn, removedRow := pack.Magnet.Column(), pack.Magnet.Row()

				for i, choice := range solutionChoices {
					if choice.TargetColumn == removedColumn && choice.TargetRow == removedRow {
						// have to invalidate this solution
						solutionChoices[i] = solutionChoices[len(solutionChoices) - 1]
						solutionChoices = solutionChoices[0 : len(solutionChoices) - 1]
						removedColumn, removedRow = choice.Magnet.Column(), choice.Magnet.Row()
						solutionChoices, unsolvedPacks = disableMoveInSystem(removedColumn, removedRow, solutionChoices, unsolvedPacks)
						break
					}
				}
			}
		}
	}
	return solutionChoices, unsolvedPacks
}

func removeElemFromSlice[T comparable](elem T, slice []T) []T {
	for i, otherElem := range slice {
		if otherElem == elem {
			slice[i] = slice[len(slice) - 1]
			return slice[0 : len(slice) - 1]
		}
	}
	return slice
}

func absInt16(a int16) int16 {
	if a >= 0 { return a }
	return -a
}
