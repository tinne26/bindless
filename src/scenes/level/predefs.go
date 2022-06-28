package level

import "bindless/src/game/iso"
import "bindless/src/game/dev"

// hardcoded game levels, veeeery raw

type levelKey int
const (
	CleanerTestDock   levelKey = 0
	CleanerTestRewire levelKey = 1
	CleanerTestReal   levelKey = 2
	ResearchLabDoor   levelKey = 3
	ResearchLabGuard1 levelKey = 4
	ResearchLabGuard2 levelKey = 5
	SwitchTest        levelKey = 6
	FinalGuard        levelKey = 7
	FinalDoor         levelKey = 8
)

const cleanerTestDockCol  , cleanerTestDockRow   int16 = 19, 15
const cleanerTestRewireCol, cleanerTestRewireRow int16 = 19, 16
const cleanerTestRealCol  , cleanerTestRealRow   int16 = 19, 19
const researchLabDoorCol  , researchLabDoorRow   int16 = 18, 19
const researchLabGuard1Col, researchLabGuard1Row int16 = 18, 19
const researchLabGuard2Col, researchLabGuard2Row int16 = 18, 18
const switchTestCol       , switchTestRow        int16 = 18, 19
const finalGuardCol       , finalGuardRow        int16 = 18, 19
const finalDoorCol        , finalDoorRow         int16 = 18, 19

func makeLevelSurface(key levelKey) iso.Map[struct{}] {
	surface := iso.NewMap[struct{}]()
	switch key {
	case CleanerTestDock:
		col, row := cleanerTestDockCol, cleanerTestDockRow
		surface.SetArea(col, row, 1, 4, struct{}{})
		row += 3
		surface.Set(col + 2, row, struct{}{})
		surface.SetArea(col - 2, row, 2, 1, struct{}{})
		col -= 2
		surface.Set(col, row - 2, struct{}{})
		surface.Set(col, row + 1, struct{}{})
		row += 2
		surface.SetArea(col, row, 3, 1, struct{}{})
		surface.Set(col - 2, row, struct{}{})
		col, row = cleanerTestDockCol, cleanerTestDockRow
		surface.SetArea(col - 5, row + 1, 2, 3, struct{}{})
		surface.Delete(col - 5, row + 1)
		surface.SetArea(col + 2, row + 5, 1, 3, struct{}{})
		surface.SetArea(col, row + 7, 2, 1, struct{}{})
	case CleanerTestRewire:
		col, row := cleanerTestRewireCol, cleanerTestRewireRow
		surface.SetArea(col - 2, row - 1, 5, 6, struct{}{})
		surface.Delete(col + 1, row)
		surface.Delete(col - 1, row + 1)
		surface.DeleteArea(col - 2, row - 1, 2, 2)
		surface.SetArea(col + 4, row - 2, 2, 2, struct{}{})
		surface.Set(col + 5, row, struct{}{})
		surface.SetArea(col - 3, row + 3, 2, 3, struct{}{})
		surface.SetArea(col + 3, row + 1, 1, 3, struct{}{})
		surface.SetArea(col, row + 6, 2, 1, struct{}{})
		surface.SetArea(col - 5, row + 4, 1, 3, struct{}{})
		surface.SetArea(col - 5, row + 7, 2, 1, struct{}{})
	case CleanerTestReal:
		col, row := cleanerTestRealCol, cleanerTestRealRow
		surface.SetArea(col - 3, row - 2, 6, 6, struct{}{})
		surface.Delete(col - 1, row)
		surface.Delete(col + 1, row)
		surface.Delete(col + 1, row + 1)
		surface.Set(col - 4, row + 1, struct{}{})
		surface.Set(col, row - 3, struct{}{})
		surface.DeleteArea(col + 1, row + 2, 2, 2)
		surface.Delete(col, row + 2)
	case ResearchLabDoor:
		col, row := researchLabDoorCol, researchLabDoorRow
		surface.SetArea(col - 4, row - 4, 9, 9, struct{}{})
		surface.DeleteArea(col - 2, row, 5, 1)
	case ResearchLabGuard1:
		col, row := researchLabGuard1Col, researchLabGuard1Row
		surface.SetArea(col - 4, row - 4, 9, 8, struct{}{})
		surface.DeleteArea(col - 4, row - 4, 2, 2)
		surface.DeleteArea(col - 4, row + 1, 2, 3)
		surface.DeleteArea(col + 0, row + 3, 5, 1)
		surface.DeleteArea(col + 4, row - 3, 1, 4)
		surface.Delete(col + 0, row - 4)
		surface.Delete(col + 0, row - 3)
		surface.Delete(col - 1, row - 3)

		surface.SetArea(col - 5, row + 2, 2, 2, struct{}{})
		surface.Set(col - 4, row + 5, struct{}{})
		surface.SetArea(col + 5, row - 2, 2, 2, struct{}{})
		surface.Set(col + 6, row - 3, struct{}{})
	case ResearchLabGuard2:
		col, row := researchLabGuard2Col, researchLabGuard2Row
		surface.SetArea(col - 4, row - 3, 11, 9, struct{}{})
		surface.SetArea(col - 2, row - 4, 5, 1, struct{}{})
		surface.SetArea(col - 6, row - 2, 1, 6, struct{}{})
		surface.SetArea(col - 7, row - 1, 1, 4, struct{}{})
		surface.Set(col + 8, row + 2, struct{}{})
		surface.DeleteArea(col - 1, row - 1, 1, 2)
		surface.DeleteArea(col - 1, row + 1, 2, 1)
		surface.DeleteArea(col, row + 2, 3, 1)
		surface.DeleteArea(col + 4, row - 4, 3, 3)
		surface.Delete(col + 6, row - 1)
		surface.Delete(col - 4, row + 5)
		surface.Delete(col + 6, row + 5)
		surface.Set(col + 5, row - 3, struct{}{})
	case SwitchTest:
		col, row := switchTestCol, switchTestRow
		surface.SetArea(col - 2, row - 2, 5, 5, struct{}{})
		surface.SetArea(col - 4, row - 1, 1, 3, struct{}{})
		surface.SetArea(col + 4, row - 1, 1, 3, struct{}{})
		surface.SetArea(col - 1, row - 4, 3, 1, struct{}{})
		surface.SetArea(col - 1, row + 4, 3, 1, struct{}{})
		surface.Set(col - 5, row, struct{}{})
		surface.Set(col + 5, row, struct{}{})
		surface.Set(col, row - 5, struct{}{})
		surface.Set(col, row + 5, struct{}{})

		surface.Set(col + 4, row + 4, struct{}{})
		surface.Set(col - 4, row - 4, struct{}{})
	case FinalGuard:
		col, row := finalGuardCol, finalGuardRow
		surface.SetArea(col - 2, row - 5, 6, 9, struct{}{})
		surface.Set(col, row + 4, struct{}{})
		surface.SetArea(col - 4, row - 2, 2, 3, struct{}{})
		surface.DeleteArea(col + 2, row - 5, 1, 8)
		surface.DeleteArea(col - 1, row - 5, 3, 1)
		surface.Delete(col - 1, row - 4)
		surface.Delete(col + 1, row - 4)
		surface.DeleteArea(col - 3, row - 2, 2, 2)
		surface.DeleteArea(col - 2, row + 3, 2, 2)
		surface.DeleteArea(col + 1, row - 1, 1, 2)
		surface.Delete(col + 1, row + 2)

		surface.Set(col + 2, row + 4, struct{}{})
		surface.Set(col + 4, row - 1, struct{}{})
		surface.Set(col + 2, row - 5, struct{}{})
		surface.Set(col + 3, row - 6, struct{}{})
	case FinalDoor:
		col, row := finalDoorCol, finalDoorRow
		surface.SetArea(col - 3, row - 4, 8, 9, struct{}{})
		surface.SetArea(col, row - 6, 3, 1, struct{}{})
		surface.Set(col + 1, row - 7, struct{}{})
		surface.Set(col - 4, row - 4, struct{}{})
		surface.Set(col + 5, row - 4, struct{}{})
		surface.Delete(col + 1, row - 3)
		surface.Delete(col + 3, row - 4)
		surface.Delete(col - 1, row - 4)
		surface.DeleteArea(col - 3, row - 2, 8, 1)
		surface.DeleteArea(col - 3, row - 2, 8, 1)

		surface.DeleteArea(col - 3, row - 1, 2, 1)
		surface.DeleteArea(col - 3, row + 1, 2, 1)
		surface.DeleteArea(col + 3, row - 1, 2, 1)
		surface.DeleteArea(col + 3, row + 1, 2, 1)
		surface.DeleteArea(col - 2, row + 3, 1, 2)
		surface.DeleteArea(col + 3, row + 3, 1, 2)
	default:
		panic("undefined level surface")
	}

	return surface
}

func makeLevelAbilities(key levelKey) Abilities {
	switch key {
	case CleanerTestDock:
		return Abilities {
			Dock: 1,
			Rewire: 0,
			Switch: -1,
			Spectre: -1,
		}
	case CleanerTestRewire:
		return Abilities {
			Dock: 0,
			Rewire: 1,
			Switch: -1,
			Spectre: -1,
		}
	case CleanerTestReal:
		return Abilities {
			Dock: 4,
			Rewire: 2,
			Switch: -1,
			Spectre: -1,
		}
	case ResearchLabDoor:
		return Abilities {
			Dock: 4,
			Rewire: 0,
			Switch: -1,
			Spectre: -1,
		}
	case ResearchLabGuard1:
		return Abilities {
			Dock: 3,
			Rewire: 2,
			Switch: -1,
			Spectre: -1,
		}
	case ResearchLabGuard2:
		return Abilities {
			Dock: 3,
			Rewire: 4,
			Switch: -1,
			Spectre: -1,
		}
	case SwitchTest:
		return Abilities {
			Dock: 0,
			Rewire: 0,
			Switch: 1,
			Spectre: -1,
		}
	case FinalGuard:
		return Abilities {
			Dock: 3,
			Rewire: 4,
			Switch: 4,
			Spectre: -1,
		}
	case FinalDoor:
		return Abilities {
			Dock: 3,
			Rewire: 4,
			Switch: 2,
			Spectre: -1,
		}
	default:
		panic("undefined level abilities")
	}
}

func makeLevelWinPoints(key levelKey) []*dev.WinPoint {
	switch key {
	case CleanerTestDock:
		col, row := cleanerTestDockCol, cleanerTestDockRow + 5
		return []*dev.WinPoint{&dev.WinPoint{ Col: col, Row: row, Polarity: dev.PolarityPositive }}
	case CleanerTestRewire:
		col, row := cleanerTestRewireCol - 2, cleanerTestRewireRow + 1
		return []*dev.WinPoint{&dev.WinPoint{ Col: col, Row: row, Polarity: dev.PolarityNegative }}
	case CleanerTestReal:
		col, row := cleanerTestRealCol - 3, cleanerTestRealRow - 2
		return []*dev.WinPoint{&dev.WinPoint{ Col: col, Row: row, Polarity: dev.PolarityPositive }}
	case ResearchLabDoor:
		col, row := researchLabDoorCol, researchLabDoorRow
		return []*dev.WinPoint {
			&dev.WinPoint{ Col: col - 4, Row: row - 1, Polarity: dev.PolarityPositive },
			&dev.WinPoint{ Col: col - 4, Row: row + 1, Polarity: dev.PolarityNegative },
			&dev.WinPoint{ Col: col + 4, Row: row - 1, Polarity: dev.PolarityPositive },
			&dev.WinPoint{ Col: col + 4, Row: row + 1, Polarity: dev.PolarityNegative },
		}
	case ResearchLabGuard1:
		col, row := researchLabGuard1Col, researchLabGuard1Row
		return []*dev.WinPoint{&dev.WinPoint{ Col: col, Row: row, Polarity: dev.PolarityPositive }}
	case ResearchLabGuard2:
		col, row := researchLabGuard2Col, researchLabGuard2Row
		return []*dev.WinPoint{&dev.WinPoint{ Col: col, Row: row, Polarity: dev.PolarityNegative }}
	case SwitchTest:
		col, row := switchTestCol, switchTestRow
		return []*dev.WinPoint{&dev.WinPoint{ Col: col, Row: row, Polarity: dev.PolarityNegative }}
	case FinalGuard:
		col, row := finalGuardCol, finalGuardRow
		return []*dev.WinPoint{&dev.WinPoint{ Col: col, Row: row - 4, Polarity: dev.PolarityPositive }}
	case FinalDoor:
		col, row := finalDoorCol, finalDoorRow
		return []*dev.WinPoint{&dev.WinPoint{ Col: col + 3, Row: row - 3, Polarity: dev.PolarityNegative }}
	default:
		panic("undefined level winPoints")
	}
}

func makeLevelDevices(key levelKey) (iso.Map[drawable], iso.Map[dev.Magnet]) {
	circuits := iso.NewMap[drawable]()
	magnets := iso.NewMap[dev.Magnet]()

	switch key {
	case CleanerTestDock:
		col, row := cleanerTestDockCol, cleanerTestDockRow
		c1, r1 := col, row
		c2, r2 := col, row + 3
		a, b := dev.NewTransferDockPair(c1, r1, c2, r2)
		circuits.Set(c1, r1, a)
		circuits.Set(c2, r2, b)

		wire := dev.NewWire2(col, row + 1, dev.ConnNW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col, row + 1, wire)
		wire  = dev.NewWire2(col, row + 2, dev.ConnNW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col, row + 2, wire)

		fm1 := dev.NewFloatMagnet(col, row, dev.StFloating, dev.PolarityPositive)
		magnets.Set(col, row, fm1)
		fm2 := dev.NewFloatMagnet(col, row + 3, dev.StDocked, dev.PolarityNeutral)
		magnets.Set(col, row + 3, fm2)

		sm1 := dev.NewStaticMagnet(col + 2, row + 3, dev.PolarityPositive.AsFunc())
		magnets.Set(col + 2, row + 3, sm1)
		sm2 := dev.NewStaticMagnet(col - 2, row + 1, dev.PolarityPositive.AsFunc())
		magnets.Set(col - 2, row + 1, sm2)
		sm3 := dev.NewStaticMagnet(col - 4, row + 5, dev.PolarityPositive.AsFunc())
		magnets.Set(col - 4, row + 5, sm3)
	case CleanerTestRewire:
		col, row := cleanerTestRewireCol, cleanerTestRewireRow
		powDock1 := dev.NewPowerDock(col + 2, row - 1)
		circuits.Set(col + 2, row - 1, powDock1)
		wire := dev.NewWire2(col + 2, row, dev.ConnNW, dev.ConnSE, powDock1.Output)
		circuits.Set(col + 2, row, wire)
		wire  = dev.NewWire2(col + 2, row + 1, dev.ConnNW, dev.ConnSW, powDock1.Output)
		circuits.Set(col + 2, row + 1, wire)
		wire  = dev.NewWire2(col + 1, row + 1, dev.ConnNE, dev.ConnSW, powDock1.Output)
		circuits.Set(col + 1, row + 1, wire)
		wSwitch := dev.NewWireSwitch(col, row + 1, dev.ConnNE, dev.ConnNW, dev.ConnSE, powDock1.Output)
		circuits.Set(col, row + 1, wSwitch)

		wire = dev.NewWire2(col, row, dev.ConnNW, dev.ConnSE, wSwitch.OutA)
		circuits.Set(col, row, wire)

		wire = dev.NewWire2(col, row + 2, dev.ConnNW, dev.ConnSE, wSwitch.OutB)
		circuits.Set(col, row + 2, wire)

		powDock2 := dev.NewPowerDock(col - 2, row + 3)
		circuits.Set(col - 2, row + 3, powDock2)
		wire = dev.NewWire2(col - 2, row + 4, dev.ConnNW, dev.ConnNE, powDock2.Output)
		circuits.Set(col - 2, row + 4, wire)
		wire = dev.NewWire2(col - 1, row + 4, dev.ConnSW, dev.ConnNE, powDock2.Output)
		circuits.Set(col - 1, row + 4, wire)
		wire = dev.NewWire2(col, row + 4, dev.ConnSW, dev.ConnNE, powDock2.Output)
		circuits.Set(col, row + 4, wire)
		wire = dev.NewWire2(col + 1, row + 4, dev.ConnSW, dev.ConnNW, powDock2.Output)
		circuits.Set(col + 1, row + 4, wire)
		wire = dev.NewWire2(col + 1, row + 3, dev.ConnSE, dev.ConnNW, powDock2.Output)
		circuits.Set(col + 1, row + 3, wire)

		fm1 := dev.NewFloatMagnet(col, row, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col, row, fm1)
		fm2 := dev.NewFloatMagnet(col + 2, row - 1, dev.StDocked, dev.PolarityPositive)
		magnets.Set(col + 2, row - 1, fm2)
		powDock1.PreSetMagnet(fm2)
		fm3 := dev.NewFloatMagnet(col - 2, row + 3, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col - 2, row + 3, fm3)
		powDock2.PreSetMagnet(fm3)

		sm1 := dev.NewStaticMagnet(col, row - 1, wSwitch.OutA)
		magnets.Set(col, row - 1, sm1)
		sm2 := dev.NewStaticMagnet(col, row + 3, wSwitch.OutB)
		magnets.Set(col, row + 3, sm2)
		sm3 := dev.NewStaticMagnet(col + 1, row + 2, powDock2.Output)
		magnets.Set(col + 1, row + 2, sm3)
	case CleanerTestReal:
		col, row := cleanerTestRealCol, cleanerTestRealRow

		fm1 := dev.NewFloatMagnet(col - 3, row + 1, dev.StDocked, dev.PolarityNeutral)
		magnets.Set(col - 3, row + 1, fm1)
		fm2 := dev.NewFloatMagnet(col - 3, row + 3, dev.StDocked, dev.PolarityPositive)
		magnets.Set(col - 3, row + 3, fm2)
		fm3 := dev.NewFloatMagnet(col - 3, row - 1, dev.StFloating, dev.PolarityPositive)
		magnets.Set(col - 3, row - 1, fm3)
		fm4 := dev.NewFloatMagnet(col - 1, row - 1, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col - 1, row - 1, fm4)
		fm5 := dev.NewFloatMagnet(col + 2, row + 1, dev.StDocked, dev.PolarityNeutral)
		magnets.Set(col + 2, row + 1, fm5)
		fm6 := dev.NewFloatMagnet(col - 2, row - 1, dev.StDocked, dev.PolarityNeutral)
		magnets.Set(col - 2, row - 1, fm6)

		a, b := dev.NewTransferDockPair(col - 1, row - 1, col + 2, row + 1)
		circuits.Set(col - 1, row - 1, a)
		circuits.Set(col + 2, row + 1, b)
		a, b = dev.NewTransferDockPair(col - 3, row + 1, col - 3, row - 1)
		circuits.Set(col - 3, row + 1, a)
		circuits.Set(col - 3, row - 1, b)

		powDock1 := dev.NewPowerDock(col - 3, row + 3)
		circuits.Set(col - 3, row + 3, powDock1)
		powDock1.PreSetMagnet(fm2)
		fm2.PreSetDockChangeHandler(powDock1)
		wSwitch := dev.NewWireSwitch(col - 2, row + 2, dev.ConnSE, dev.ConnNW, dev.ConnNE, powDock1.Output)
		circuits.Set(col - 2, row + 2, wSwitch)

		sm1 := dev.NewStaticMagnet(col - 4, row + 1, dev.PolarityPositive.AsFunc())
		magnets.Set(col - 4, row + 1, sm1)
		sm3 := dev.NewStaticMagnet(col + 1, row - 2, wSwitch.OutA)
		magnets.Set(col + 1, row - 2, sm3)
		sm4 := dev.NewStaticMagnet(col, row + 3, wSwitch.OutB)
		magnets.Set(col, row + 3, sm4)
		sm5 := dev.NewStaticMagnet(col, row - 3, dev.PolarityNegative.AsFunc())
		magnets.Set(col, row - 3, sm5)

		// transfer dock pair wires
		wire := dev.NewWire2(col - 3, row, dev.ConnNW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col - 3, row, wire)

		wire = dev.NewWire2(col, row - 1, dev.ConnSW, dev.ConnNE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col, row - 1, wire)
		wire = dev.NewWire2(col + 1, row - 1, dev.ConnSW, dev.ConnNE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 1, row - 1, wire)
		wire = dev.NewWire2(col + 2, row - 1, dev.ConnSW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 2, row - 1, wire)

		// switch wires
		wire = dev.NewWire2(col - 2, row + 3, dev.ConnSW, dev.ConnNW, powDock1.Output)
		circuits.Set(col - 2, row + 3, wire)
		wire = dev.NewWire2(col - 1, row + 2, dev.ConnSW, dev.ConnSE, wSwitch.OutB)
		circuits.Set(col - 1, row + 2, wire)
		wire = dev.NewWire2(col - 1, row + 3, dev.ConnNW, dev.ConnNE, wSwitch.OutB)
		circuits.Set(col - 1, row + 3, wire)

		wire = dev.NewWire2(col - 2, row + 1, dev.ConnSE, dev.ConnNW, wSwitch.OutA)
		circuits.Set(col - 2, row + 1, wire)
		wire = dev.NewWire2(col - 2, row, dev.ConnSE, dev.ConnNW, wSwitch.OutA)
		circuits.Set(col - 2, row, wire)
		wire = dev.NewWire2(col - 2, row - 1, dev.ConnSE, dev.ConnNW, wSwitch.OutA)
		circuits.Set(col - 2, row - 1, wire)
		wire = dev.NewWire2(col - 2, row - 2, dev.ConnSE, dev.ConnNE, wSwitch.OutA)
		circuits.Set(col - 2, row - 2, wire)
		wire = dev.NewWire2(col - 1, row - 2, dev.ConnSW, dev.ConnNE, wSwitch.OutA)
		circuits.Set(col - 1, row - 2, wire)
		wire = dev.NewWire2(col, row - 2, dev.ConnSW, dev.ConnNE, wSwitch.OutA)
		circuits.Set(col, row - 2, wire)
		wire = dev.NewWire2(col + 2, row, dev.ConnNW, dev.ConnSE, wSwitch.OutA)
		circuits.Set(col + 2, row, wire)
	case ResearchLabDoor:
		col, row := researchLabDoorCol, researchLabDoorRow

		powDockPositive1 := dev.NewPowerDock(col - 2, row - 1)
		circuits.Set(col - 2, row - 1, powDockPositive1)
		powDockPositive2 := dev.NewPowerDock(col - 0, row - 1)
		circuits.Set(col - 0, row - 1, powDockPositive2)
		powDockPositive3 := dev.NewPowerDock(col + 2, row - 1)
		circuits.Set(col + 2, row - 1, powDockPositive3)

		powDockNegative1 := dev.NewPowerDock(col - 2, row + 1)
		circuits.Set(col - 2, row + 1, powDockNegative1)
		powDockNegative2 := dev.NewPowerDock(col - 0, row + 1)
		circuits.Set(col - 0, row + 1, powDockNegative2)
		powDockNegative3 := dev.NewPowerDock(col + 2, row + 1)
		circuits.Set(col + 2, row + 1, powDockNegative3)

		smPositive1 := dev.NewStaticMagnet(col - 2, row - 4, powDockPositive1.Output)
		magnets.Set(col - 2, row - 4, smPositive1)
		smPositive2 := dev.NewStaticMagnet(col - 0, row - 4, powDockPositive2.Output)
		magnets.Set(col - 0, row - 4, smPositive2)
		smPositive3 := dev.NewStaticMagnet(col + 2, row - 4, powDockPositive3.Output)
		magnets.Set(col + 2, row - 4, smPositive3)

		smNegative1 := dev.NewStaticMagnet(col - 2, row + 4, powDockNegative1.Output)
		magnets.Set(col - 2, row + 4, smNegative1)
		smNegative2 := dev.NewStaticMagnet(col - 0, row + 4, powDockNegative2.Output)
		magnets.Set(col - 0, row + 4, smNegative2)
		smNegative3 := dev.NewStaticMagnet(col + 2, row + 4, powDockNegative3.Output)
		magnets.Set(col + 2, row + 4, smNegative3)

		fmPositive1 := dev.NewFloatMagnet(col - 2, row - 1, dev.StDocked, dev.PolarityPositive)
		magnets.Set(col - 2, row - 1, fmPositive1)
		fmPositive2 := dev.NewFloatMagnet(col - 0, row - 1, dev.StDocked, dev.PolarityPositive)
		magnets.Set(col - 0, row - 1, fmPositive2)
		fmPositive3 := dev.NewFloatMagnet(col + 2, row - 1, dev.StDocked, dev.PolarityPositive)
		magnets.Set(col + 2, row - 1, fmPositive3)
		fmNegative1 := dev.NewFloatMagnet(col - 2, row + 1, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col - 2, row + 1, fmNegative1)
		fmNegative2 := dev.NewFloatMagnet(col - 0, row + 1, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col - 0, row + 1, fmNegative2)
		fmNegative3 := dev.NewFloatMagnet(col + 2, row + 1, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col + 2, row + 1, fmNegative3)

		powDockPositive1.PreSetMagnet(fmPositive1)
		powDockPositive2.PreSetMagnet(fmPositive2)
		powDockPositive3.PreSetMagnet(fmPositive3)
		powDockNegative1.PreSetMagnet(fmNegative1)
		powDockNegative2.PreSetMagnet(fmNegative2)
		powDockNegative3.PreSetMagnet(fmNegative3)
		fmPositive1.PreSetDockChangeHandler(powDockPositive1)
		fmPositive2.PreSetDockChangeHandler(powDockPositive2)
		fmPositive3.PreSetDockChangeHandler(powDockPositive3)
		fmNegative1.PreSetDockChangeHandler(powDockNegative1)
		fmNegative2.PreSetDockChangeHandler(powDockNegative2)
		fmNegative3.PreSetDockChangeHandler(powDockNegative3)

		wire := dev.NewWire2(col - 2, row - 2, dev.ConnSE, dev.ConnNW, powDockPositive1.Output)
		circuits.Set(col - 2, row - 2, wire)
		wire  = dev.NewWire2(col - 0, row - 2, dev.ConnSE, dev.ConnNW, powDockPositive2.Output)
		circuits.Set(col - 0, row - 2, wire)
		wire  = dev.NewWire2(col + 2, row - 2, dev.ConnSE, dev.ConnNW, powDockPositive3.Output)
		circuits.Set(col + 2, row - 2, wire)
		wire  = dev.NewWire2(col - 2, row - 3, dev.ConnSE, dev.ConnNW, powDockPositive1.Output)
		circuits.Set(col - 2, row - 3, wire)
		wire  = dev.NewWire2(col - 0, row - 3, dev.ConnSE, dev.ConnNW, powDockPositive2.Output)
		circuits.Set(col - 0, row - 3, wire)
		wire  = dev.NewWire2(col + 2, row - 3, dev.ConnSE, dev.ConnNW, powDockPositive3.Output)
		circuits.Set(col + 2, row - 3, wire)

		wire  = dev.NewWire2(col - 2, row + 2, dev.ConnSE, dev.ConnNW, powDockNegative1.Output)
		circuits.Set(col - 2, row + 2, wire)
		wire  = dev.NewWire2(col - 0, row + 2, dev.ConnSE, dev.ConnNW, powDockNegative2.Output)
		circuits.Set(col - 0, row + 2, wire)
		wire  = dev.NewWire2(col + 2, row + 2, dev.ConnSE, dev.ConnNW, powDockNegative3.Output)
		circuits.Set(col + 2, row + 2, wire)
		wire  = dev.NewWire2(col - 2, row + 3, dev.ConnSE, dev.ConnNW, powDockNegative1.Output)
		circuits.Set(col - 2, row + 3, wire)
		wire  = dev.NewWire2(col - 0, row + 3, dev.ConnSE, dev.ConnNW, powDockNegative2.Output)
		circuits.Set(col - 0, row + 3, wire)
		wire  = dev.NewWire2(col + 2, row + 3, dev.ConnSE, dev.ConnNW, powDockNegative3.Output)
		circuits.Set(col + 2, row + 3, wire)
	case ResearchLabGuard1:
		col, row := researchLabGuard1Col, researchLabGuard1Row

		powDockNW := dev.NewPowerDock(col - 2, row - 4)
		circuits.Set(col - 2, row - 4, powDockNW)
		powDockSE := dev.NewPowerDock(col - 2, row + 3)
		circuits.Set(col - 2, row + 3, powDockSE)
		powDockCenter := dev.NewPowerDock(col, row + 1)
		circuits.Set(col, row + 1, powDockCenter)

		wSwitch := dev.NewWireSwitch(col - 1, row + 2, dev.ConnSE, dev.ConnNE, dev.ConnSW, powDockSE.Output)
		circuits.Set(col - 1, row + 2, wSwitch)

		smDockNW := dev.NewStaticMagnet(col - 3, row - 1, powDockNW.Output)
		magnets.Set(col - 3, row - 1, smDockNW)
		smDockSE_A := dev.NewStaticMagnet(col + 4, row + 1, wSwitch.OutA)
		magnets.Set(col + 4, row + 1, smDockSE_A)
		smDockSE_B := dev.NewStaticMagnet(col - 2, row + 1, wSwitch.OutB)
		magnets.Set(col - 2, row + 1, smDockSE_B)
		smDockCenter := dev.NewStaticMagnet(col - 3, row - 2, powDockCenter.Output)
		magnets.Set(col - 3, row - 2, smDockCenter)

		smTop := dev.NewStaticMagnet(col + 2, row - 2, dev.PolarityPositive.AsFunc())
		magnets.Set(col + 2, row - 2, smTop)

		a, b := dev.NewTransferDockPair(col + 1, row - 4, col + 3, row - 1)
		circuits.Set(col + 1, row - 4, a)
		circuits.Set(col + 3, row - 1, b)

		fmDead := dev.NewFloatMagnet(col + 2, row, dev.StDocked, dev.PolarityNeutral)
		magnets.Set(col + 2, row, fmDead)
		fmTransA := dev.NewFloatMagnet(col + 3, row - 1, dev.StFloating, dev.PolarityPositive)
		magnets.Set(col + 3, row - 1, fmTransA)
		fmTransB := dev.NewFloatMagnet(col + 1, row - 4, dev.StDocked, dev.PolarityNeutral)
		magnets.Set(col + 1, row - 4, fmTransB)

		fmMain := dev.NewFloatMagnet(col + 1, row - 3, dev.StFloating, dev.PolarityPositive)
		magnets.Set(col + 1, row - 3, fmMain)

		fmDockNW := dev.NewFloatMagnet(col - 1, row - 4, dev.StFloating, dev.PolarityPositive)
		magnets.Set(col - 1, row - 4, fmDockNW)
		fmDockSE := dev.NewFloatMagnet(col - 2, row + 3, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col - 2, row + 3, fmDockSE)
		powDockSE.PreSetMagnet(fmDockSE)
		fmDockSE.PreSetDockChangeHandler(powDockSE)

		fmDockCenter := dev.NewFloatMagnet(col - 1, row + 1, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col - 1, row + 1, fmDockCenter)

		// SE dock wires
		wire := dev.NewWire2(col - 1, row + 3, dev.ConnSW, dev.ConnNW, powDockSE.Output)
		circuits.Set(col - 1, row + 3, wire)
		wire  = dev.NewWire2(col + 0, row + 2, dev.ConnSW, dev.ConnNE, wSwitch.OutA)
		circuits.Set(col + 0, row + 2, wire)
		wire  = dev.NewWire2(col + 1, row + 2, dev.ConnSW, dev.ConnNE, wSwitch.OutA)
		circuits.Set(col + 1, row + 2, wire)
		wire  = dev.NewWire2(col + 2, row + 2, dev.ConnSW, dev.ConnNE, wSwitch.OutA)
		circuits.Set(col + 2, row + 2, wire)
		wire  = dev.NewWire2(col + 3, row + 2, dev.ConnSW, dev.ConnNE, wSwitch.OutA)
		circuits.Set(col + 3, row + 2, wire)
		wire  = dev.NewWire2(col + 4, row + 2, dev.ConnSW, dev.ConnNW, wSwitch.OutA)
		circuits.Set(col + 4, row + 2, wire)

		wire  = dev.NewWire2(col - 2, row + 2, dev.ConnNE, dev.ConnNW, wSwitch.OutB)
		circuits.Set(col - 2, row + 2, wire)

		// central dock wires
		wire  = dev.NewWire2(col - 1, row + 1, dev.ConnNE, dev.ConnNW, powDockCenter.Output)
		circuits.Set(col - 1, row + 1, wire)
		wire  = dev.NewWire2(col - 1, row + 0, dev.ConnSE, dev.ConnSW, powDockCenter.Output)
		circuits.Set(col - 1, row + 0, wire)
		wire  = dev.NewWire2(col - 2, row + 0, dev.ConnNE, dev.ConnSW, powDockCenter.Output)
		circuits.Set(col - 2, row + 0, wire)

		wire  = dev.NewWire2(col - 3, row + 0, dev.ConnNE, dev.ConnSW, powDockCenter.Output)
		circuits.Set(col - 3, row + 0, wire)
		wire  = dev.NewWire2(col - 4, row + 0, dev.ConnNE, dev.ConnNW, powDockCenter.Output)
		circuits.Set(col - 4, row + 0, wire)
		wire  = dev.NewWire2(col - 4, row - 1, dev.ConnSE, dev.ConnNW, powDockCenter.Output)
		circuits.Set(col - 4, row - 1, wire)
		wire  = dev.NewWire2(col - 4, row - 2, dev.ConnSE, dev.ConnNE, powDockCenter.Output)
		circuits.Set(col - 4, row - 2, wire)

		// transfer wires
		wire  = dev.NewWire2(col + 2, row - 4, dev.ConnSW, dev.ConnNE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 2, row - 4, wire)
		wire  = dev.NewWire2(col + 3, row - 4, dev.ConnSW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 3, row - 4, wire)
		wire  = dev.NewWire2(col + 3, row - 3, dev.ConnNW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 3, row - 3, wire)
		wire  = dev.NewWire2(col + 3, row - 2, dev.ConnNW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 3, row - 2, wire)

		// NW dock wires
		wire  = dev.NewWire2(col - 2, row - 3, dev.ConnNW, dev.ConnSE, powDockNW.Output)
		circuits.Set(col - 2, row - 3, wire)
		wire  = dev.NewWire2(col - 2, row - 2, dev.ConnNW, dev.ConnSE, powDockNW.Output)
		circuits.Set(col - 2, row - 2, wire)
		wire  = dev.NewWire2(col - 2, row - 1, dev.ConnNW, dev.ConnSW, powDockNW.Output)
		circuits.Set(col - 2, row - 1, wire)
	case ResearchLabGuard2:
		col, row := researchLabGuard2Col, researchLabGuard2Row

		powDockAlone := dev.NewPowerDock(col + 8, row + 2)
		circuits.Set(col + 8, row + 2, powDockAlone)
		powDockTop := dev.NewPowerDock(col + 6, row + 0)
		circuits.Set(col + 6, row + 0, powDockTop)
		powDockNE := dev.NewPowerDock(col + 6, row + 1)
		circuits.Set(col + 6, row + 1, powDockNE)
		powDockSwitchRight := dev.NewPowerDock(col + 3, row + 5)
		circuits.Set(col + 3, row + 5, powDockSwitchRight)
		powDockSwitchLeft := dev.NewPowerDock(col - 3, row - 2)
		circuits.Set(col - 3, row - 2, powDockSwitchLeft)

		wSwitchRight := dev.NewWireSwitch(col + 2, row + 4, dev.ConnSE, dev.ConnNE, dev.ConnSW, powDockSwitchRight.Output)
		circuits.Set(col + 2, row + 4, wSwitchRight)
		wSwitchLeft := dev.NewWireSwitch(col - 2, row - 2, dev.ConnSW, dev.ConnNE, dev.ConnSE, powDockSwitchLeft.Output)
		circuits.Set(col - 2, row - 2, wSwitchLeft)

		a, b := dev.NewTransferDockPair(col + 2, row - 2, col + 6, row + 2)
		circuits.Set(col + 2, row - 2, a)
		circuits.Set(col + 6, row + 2, b)

		smSwitchRightA := dev.NewStaticMagnet(col + 3, row + 4, wSwitchRight.OutA)
		magnets.Set(col + 3, row + 4, smSwitchRightA)
		smSwitchRightB := dev.NewStaticMagnet(col - 3, row + 3, wSwitchRight.OutB)
		magnets.Set(col - 3, row + 3, smSwitchRightB)
		smSwitchLeftA := dev.NewStaticMagnet(col + 0, row - 2, wSwitchLeft.OutA)
		magnets.Set(col + 0, row - 2, smSwitchLeftA)
		smSwitchLeftB := dev.NewStaticMagnet(col - 2, row + 0, wSwitchLeft.OutB)
		magnets.Set(col - 2, row + 0, smSwitchLeftB)

		smPowDockTop := dev.NewStaticMagnet(col + 4, row - 1, powDockTop.Output)
		magnets.Set(col + 4, row - 1, smPowDockTop)
		smPowDockNE := dev.NewStaticMagnet(col + 4, row + 3, powDockNE.Output)
		magnets.Set(col + 4, row + 3, smPowDockNE)

		smFixedNW := dev.NewStaticMagnet(col + 2, row - 3, dev.PolarityPositive.AsFunc())
		magnets.Set(col + 2, row - 3, smFixedNW)
		smFixedNW2 := dev.NewStaticMagnet(col + 3, row - 3, dev.PolarityPositive.AsFunc())
		magnets.Set(col + 3, row - 3, smFixedNW2)

		fmTransInactive := dev.NewFloatMagnet(col + 6, row + 2, dev.StDocked, dev.PolarityNeutral)
		magnets.Set(col + 6, row + 2, fmTransInactive)
		fmTransActive := dev.NewFloatMagnet(col + 1, row + 0, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col + 1, row + 0, fmTransActive)

		fmBlocker := dev.NewFloatMagnet(col + 5, row + 4, dev.StDocked, dev.PolarityNeutral)
		magnets.Set(col + 5, row + 4, fmBlocker)

		fmPowDockTop := dev.NewFloatMagnet(col + 6, row + 1, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col + 6, row + 1, fmPowDockTop)
		fmPowDockNE := dev.NewFloatMagnet(col + 6, row + 4, dev.StFloating, dev.PolarityPositive)
		magnets.Set(col + 6, row + 4, fmPowDockNE)

		fmPowDockAlone := dev.NewFloatMagnet(col + 8, row + 2, dev.StDocked, dev.PolarityPositive)
		magnets.Set(col + 8, row + 2, fmPowDockAlone)

		fmSwitchRight := dev.NewFloatMagnet(col + 3, row + 5, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col + 3, row + 5, fmSwitchRight)
		fmSwitchLeft := dev.NewFloatMagnet(col - 3, row - 2, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col - 3, row - 2, fmSwitchLeft)

		fmMain := dev.NewFloatMagnet(col - 2, row + 2, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col - 2, row + 2, fmMain)

		powDockSwitchLeft.PreSetMagnet(fmSwitchLeft)
		fmSwitchLeft.PreSetDockChangeHandler(powDockSwitchLeft)
		powDockSwitchRight.PreSetMagnet(fmSwitchRight)
		fmSwitchRight.PreSetDockChangeHandler(powDockSwitchRight)
		powDockAlone.PreSetMagnet(fmPowDockAlone)
		fmPowDockAlone.PreSetDockChangeHandler(powDockAlone)

		// left switch wires
		wire := dev.NewWire2(col - 1, row - 2, dev.ConnSW, dev.ConnNE, wSwitchLeft.OutA)
		circuits.Set(col - 1, row - 2, wire)
		wire  = dev.NewWire2(col - 2, row - 1, dev.ConnNW, dev.ConnSE, wSwitchLeft.OutB)
		circuits.Set(col - 2, row - 1, wire)

		// right switch wires
		wire  = dev.NewWire2(col + 2, row + 5, dev.ConnNW, dev.ConnNE, powDockSwitchRight.Output)
		circuits.Set(col + 2, row + 5, wire)
		wire  = dev.NewWire2(col + 1, row + 4, dev.ConnNE, dev.ConnSW, wSwitchRight.OutB)
		circuits.Set(col + 1, row + 4, wire)
		wire  = dev.NewWire2(col + 0, row + 4, dev.ConnNE, dev.ConnSW, wSwitchRight.OutB)
		circuits.Set(col + 0, row + 4, wire)
		wire  = dev.NewWire2(col - 1, row + 4, dev.ConnNE, dev.ConnSW, wSwitchRight.OutB)
		circuits.Set(col - 1, row + 4, wire)
		wire  = dev.NewWire2(col - 2, row + 4, dev.ConnNE, dev.ConnSW, wSwitchRight.OutB)
		circuits.Set(col - 2, row + 4, wire)
		wire  = dev.NewWire2(col - 3, row + 4, dev.ConnNE, dev.ConnNW, wSwitchRight.OutB)
		circuits.Set(col - 3, row + 4, wire)

		// top pow wires
		wire  = dev.NewWire2(col + 5, row + 0, dev.ConnNW, dev.ConnNE, powDockTop.Output)
		circuits.Set(col + 5, row + 0, wire)
		wire  = dev.NewWire2(col + 5, row - 1, dev.ConnSE, dev.ConnSW, powDockTop.Output)
		circuits.Set(col + 5, row - 1, wire)

		// NE pow wires
		wire  = dev.NewWire2(col + 5, row + 1, dev.ConnNE, dev.ConnSE, powDockNE.Output)
		circuits.Set(col + 5, row + 1, wire)
		wire  = dev.NewWire2(col + 5, row + 3, dev.ConnNW, dev.ConnSW, powDockNE.Output)
		circuits.Set(col + 5, row + 3, wire)

		// transfer wires
		wire  = dev.NewWire2(col + 2, row - 1, dev.ConnNW, dev.ConnNE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 2, row - 1, wire)
		wire  = dev.NewWire2(col + 3, row - 1, dev.ConnSW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 3, row - 1, wire)
		wire  = dev.NewWire2(col + 3, row + 0, dev.ConnNW, dev.ConnNE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 3, row + 0, wire)
		wire  = dev.NewWire2(col + 4, row + 0, dev.ConnSW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 4, row + 0, wire)
		wire  = dev.NewWire2(col + 4, row + 1, dev.ConnNW, dev.ConnSE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 4, row + 1, wire)
		wire  = dev.NewWire2(col + 4, row + 2, dev.ConnNW, dev.ConnNE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 4, row + 2, wire)
		wire  = dev.NewWire2(col + 5, row + 2, dev.ConnSW, dev.ConnNE, dev.PolarityNeutral.AsFunc())
		circuits.Set(col + 5, row + 2, wire)
	case SwitchTest:
		col, row := switchTestCol, switchTestRow
		smNW := dev.NewStaticMagnet(col + 1, row - 2, dev.PolarityNegative.AsFunc())
		magnets.Set(col + 1, row - 2, smNW)
		smNE := dev.NewStaticMagnet(col + 2, row + 1, dev.PolarityNegative.AsFunc())
		magnets.Set(col + 2, row + 1, smNE)
		smSW := dev.NewStaticMagnet(col - 2, row - 1, dev.PolarityNegative.AsFunc())
		magnets.Set(col - 2, row - 1, smSW)
		smSE := dev.NewStaticMagnet(col - 1, row + 2, dev.PolarityNegative.AsFunc())
		magnets.Set(col - 1, row + 2, smSE)
		smPush := dev.NewStaticMagnet(col + 0, row + 4, dev.PolarityNegative.AsFunc())
		magnets.Set(col + 0, row + 4, smPush)

		fmMain := dev.NewFloatMagnet(col - 1, row -  1, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col - 1, row - 1, fmMain)
		fmKey := dev.NewFloatMagnet(col - 2, row + 1, dev.StDocked, dev.PolarityPositive)
		magnets.Set(col - 2, row + 1, fmKey)

		powDock := dev.NewPowerDock(col - 2, row + 1) // no need to link, can't use Dock anyway
		circuits.Set(col - 2, row + 1, powDock)
	case FinalGuard:
		col, row := finalGuardCol, finalGuardRow

		powDockGoal := dev.NewPowerDock(col + 0, row - 3)
		circuits.Set(col + 0, row - 3, powDockGoal)
		powDockStart := dev.NewPowerDock(col + 0, row + 3)
		circuits.Set(col + 0, row + 3, powDockStart)
		powDockPerma := dev.NewPowerDock(col - 2, row + 2)
		circuits.Set(col - 2, row + 2, powDockPerma)

		wSwitch := dev.NewWireSwitch(col - 1, row + 1, dev.ConnSW, dev.ConnNE, dev.ConnSE, powDockPerma.Output)
		circuits.Set(col - 1, row + 1, wSwitch)

		smPowStart := dev.NewStaticMagnet(col + 3, row + 0, powDockStart.Output)
		magnets.Set(col + 3, row + 0, smPowStart)
		smStart := dev.NewStaticMagnet(col + 0, row + 4, dev.PolarityPositive.AsFunc())
		magnets.Set(col + 0, row + 4, smStart)
		smSidePush := dev.NewStaticMagnet(col - 4, row - 1, powDockGoal.Output)
		magnets.Set(col - 4, row - 1, smSidePush)
		smSwitchA := dev.NewStaticMagnet(col + 1, row + 1, wSwitch.OutA)
		magnets.Set(col + 1, row + 1, smSwitchA)
		smSwitchB := dev.NewStaticMagnet(col - 1, row + 2, wSwitch.OutB)
		magnets.Set(col - 1, row + 2, smSwitchB)
		smBounce := dev.NewStaticMagnet(col + 3, row - 5, dev.PolarityNegative.AsFunc())
		magnets.Set(col + 3, row - 5, smBounce)

		fmGoal   := dev.NewFloatMagnet(col + 0, row - 2, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col + 0, row - 2, fmGoal)
		fmBounce := dev.NewFloatMagnet(col + 3, row - 3, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col + 3, row - 3, fmBounce)
		fmStart  := dev.NewFloatMagnet(col + 0, row + 3, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col + 0, row + 3, fmStart)
		fmFloat  := dev.NewFloatMagnet(col - 1, row + 0, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col - 1, row + 0, fmFloat)
		fmPerma  := dev.NewFloatMagnet(col - 2, row + 2, dev.StDocked, dev.PolarityPositive)
		magnets.Set(col - 2, row + 2, fmPerma)

		powDockPerma.PreSetMagnet(fmPerma)
		fmPerma.PreSetDockChangeHandler(powDockPerma)
		powDockStart.PreSetMagnet(fmStart)
		fmStart.PreSetDockChangeHandler(powDockStart)

		wire := dev.NewWire2(col - 2, row + 1, dev.ConnSE, dev.ConnNE, powDockPerma.Output)
		circuits.Set(col - 2, row + 1, wire)
		wire  = dev.NewWire2(col + 0, row + 1, dev.ConnSW, dev.ConnNE, wSwitch.OutA)
		circuits.Set(col + 0, row + 1, wire)

		wire  = dev.NewWire2(col + 1, row + 3, dev.ConnSW, dev.ConnNE, powDockStart.Output)
		circuits.Set(col + 1, row + 3, wire)
		wire  = dev.NewWire2(col + 2, row + 3, dev.ConnSW, dev.ConnNE, powDockStart.Output)
		circuits.Set(col + 2, row + 3, wire)
		wire  = dev.NewWire2(col + 3, row + 3, dev.ConnSW, dev.ConnNW, powDockStart.Output)
		circuits.Set(col + 3, row + 3, wire)
		wire  = dev.NewWire2(col + 3, row + 2, dev.ConnSE, dev.ConnNW, powDockStart.Output)
		circuits.Set(col + 3, row + 2, wire)
		wire  = dev.NewWire2(col + 3, row + 1, dev.ConnSE, dev.ConnNW, powDockStart.Output)
		circuits.Set(col + 3, row + 1, wire)

		wire  = dev.NewWire2(col - 1, row - 3, dev.ConnNE, dev.ConnSE, powDockGoal.Output)
		circuits.Set(col - 1, row - 3, wire)
		wire  = dev.NewWire2(col - 1, row - 2, dev.ConnNW, dev.ConnSE, powDockGoal.Output)
		circuits.Set(col - 1, row - 2, wire)
		wire  = dev.NewWire2(col - 1, row - 1, dev.ConnNW, dev.ConnSE, powDockGoal.Output)
		circuits.Set(col - 1, row - 1, wire)
		wire  = dev.NewWire2(col - 1, row + 0, dev.ConnNW, dev.ConnSW, powDockGoal.Output)
		circuits.Set(col - 1, row + 0, wire)

		wire  = dev.NewWire2(col - 2, row + 0, dev.ConnNE, dev.ConnSW, powDockGoal.Output)
		circuits.Set(col - 2, row + 0, wire)
		wire  = dev.NewWire2(col - 3, row + 0, dev.ConnNE, dev.ConnSW, powDockGoal.Output)
		circuits.Set(col - 3, row + 0, wire)
		wire  = dev.NewWire2(col - 4, row + 0, dev.ConnNE, dev.ConnNW, powDockGoal.Output)
		circuits.Set(col - 4, row + 0, wire)
	case FinalDoor:
		col, row := finalDoorCol, finalDoorRow

		powDockTun1 := dev.NewPowerDock(col + 2, row + 4)
		circuits.Set(col + 2, row + 4, powDockTun1)
		powDockTun2 := dev.NewPowerDock(col + 0, row + 4)
		circuits.Set(col + 0, row + 4, powDockTun2)
		powDockBridgePush := dev.NewPowerDock(col - 4, row - 4)
		circuits.Set(col - 4, row - 4, powDockBridgePush)

		switch1 := dev.NewWireSwitch(col + 2, row + 2, dev.ConnSE, dev.ConnNW, dev.ConnSW, powDockTun1.Output)
		circuits.Set(col + 2, row + 2, switch1)
		switch2 := dev.NewWireSwitch(col + 0, row + 2, dev.ConnSE, dev.ConnNW, dev.ConnSW, powDockTun2.Output)
		circuits.Set(col + 0, row + 2, switch2)

		a, b := dev.NewTransferDockPair(col - 2, row + 0, col - 2, row - 3)
		circuits.Set(col - 2, row + 0, a)
		circuits.Set(col - 2, row - 3, b)

		smTun1A := dev.NewStaticMagnet(col + 2, row - 1, switch1.OutA)
		magnets.Set(col + 2, row - 1, smTun1A)
		smTun1B := dev.NewStaticMagnet(col + 1, row + 1, switch1.OutB)
		magnets.Set(col + 1, row + 1, smTun1B)
		smTun2A := dev.NewStaticMagnet(col + 0, row - 1, switch2.OutA)
		magnets.Set(col + 0, row - 1, smTun2A)
		smTun2B := dev.NewStaticMagnet(col - 1, row + 1, switch2.OutB)
		magnets.Set(col - 1, row + 1, smTun2B)

		smTunPush := dev.NewStaticMagnet(col + 4, row + 0, dev.PolarityNegative.AsFunc())
		magnets.Set(col + 4, row + 0, smTunPush)
		smTunPull := dev.NewStaticMagnet(col - 3, row + 0, dev.PolarityPositive.AsFunc())
		magnets.Set(col - 3, row + 0, smTunPull)

		smBridgePushA := dev.NewStaticMagnet(col - 3, row - 3, dev.PolarityNegative.AsFunc())
		magnets.Set(col - 3, row - 3, smBridgePushA)
		smBridgePushB := dev.NewStaticMagnet(col - 2, row - 4, powDockBridgePush.Output)
		magnets.Set(col - 2, row - 4, smBridgePushB)
		smBridgePullA := dev.NewStaticMagnet(col + 4, row - 3, dev.PolarityPositive.AsFunc())
		magnets.Set(col + 4, row - 3, smBridgePullA)

		fmMain := dev.NewFloatMagnet(col + 3, row + 0, dev.StFloating, dev.PolarityPositive)
		magnets.Set(col + 3, row + 0, fmMain)
		fmMain2 := dev.NewFloatMagnet(col - 2, row - 3, dev.StDocked, dev.PolarityNeutral)
		magnets.Set(col - 2, row - 3, fmMain2)
		fmBridgePush := dev.NewFloatMagnet(col - 4, row - 4, dev.StFloating, dev.PolarityNegative)
		magnets.Set(col - 4, row - 4, fmBridgePush)
		fmSwitch1 := dev.NewFloatMagnet(col + 2, row + 4, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col + 2, row + 4, fmSwitch1)
		fmSwitch2 := dev.NewFloatMagnet(col + 0, row + 4, dev.StDocked, dev.PolarityNegative)
		magnets.Set(col + 0, row + 4, fmSwitch2)

		powDockTun1.PreSetMagnet(fmSwitch1)
		fmSwitch1.PreSetDockChangeHandler(powDockTun1)
		powDockTun2.PreSetMagnet(fmSwitch2)
		fmSwitch2.PreSetDockChangeHandler(powDockTun2)

		wire := dev.NewWire2(col + 2, row + 3, dev.ConnSE, dev.ConnNW, powDockTun1.Output)
		circuits.Set(col + 2, row + 3, wire)
		wire  = dev.NewWire2(col + 0, row + 3, dev.ConnSE, dev.ConnNW, powDockTun2.Output)
		circuits.Set(col + 0, row + 3, wire)

		wire  = dev.NewWire2(col + 2, row + 1, dev.ConnSE, dev.ConnNW, switch1.OutA)
		circuits.Set(col + 2, row + 1, wire)
		wire  = dev.NewWire2(col + 2, row + 0, dev.ConnSE, dev.ConnNW, switch1.OutA)
		circuits.Set(col + 2, row + 0, wire)
		wire  = dev.NewWire2(col + 0, row + 1, dev.ConnSE, dev.ConnNW, switch2.OutA)
		circuits.Set(col + 0, row + 1, wire)
		wire  = dev.NewWire2(col + 0, row + 0, dev.ConnSE, dev.ConnNW, switch2.OutA)
		circuits.Set(col + 0, row + 0, wire)

		wire  = dev.NewWire2(col + 1, row + 2, dev.ConnNE, dev.ConnNW, switch1.OutB)
		circuits.Set(col + 1, row + 2, wire)
		wire  = dev.NewWire2(col - 1, row + 2, dev.ConnNE, dev.ConnNW, switch2.OutB)
		circuits.Set(col - 1, row + 2, wire)

		wire  = dev.NewWire2(col - 3, row - 4, dev.ConnSW, dev.ConnNE, powDockBridgePush.Output)
		circuits.Set(col - 3, row - 4, wire)

	// case TemplateLevelName:
	// 	col, row := templateLevelNameCol, templateLevelNameRow
	//
	// 	powDockTop := dev.NewPowerDock(col + 0, row + 0)
	// 	circuits.Set(col + 0, row + 0, powDockTop)
	// 	powDockNE  := dev.NewPowerDock(col + 0, row + 0)
	// 	circuits.Set(col + 0, row + 0, powDockNE)
	// 	powDockSWL := dev.NewPowerDock(col + 0, row + 0)
	// 	circuits.Set(col + 0, row + 0, powDockSWL)
	// 	powDockSWR := dev.NewPowerDock(col + 0, row + 0)
	// 	circuits.Set(col + 0, row + 0, powDockSWR)
	//
	// 	wSwitchTop := dev.NewWireSwitch(col + 0, row + 0, dev.ConnSE, dev.ConnNE, dev.ConnSW, powDockSE.Output)
	// 	circuits.Set(col + 0, row + 0, wSwitchTop)
	// 	wSwitchSWL := dev.NewWireSwitch(col + 0, row + 0, dev.ConnSE, dev.ConnNE, dev.ConnSW, powDockSE.Output)
	// 	circuits.Set(col + 0, row + 0, wSwitchSWL)
	// 	wSwitchSWR := dev.NewWireSwitch(col + 0, row + 0, dev.ConnSE, dev.ConnNE, dev.ConnSW, powDockSE.Output)
	// 	circuits.Set(col + 0, row + 0, wSwitchSWR)
	//
	// 	a, b := dev.NewTransferDockPair(col + 0, row + 0, col + 0, row + 0)
	// 	circuits.Set(col + 0, row + 0, a)
	// 	circuits.Set(col + 0, row + 0, b)
	//
	// 	smSwitchTopA := dev.NewStaticMagnet(col + 0, row + 0, wSwitchTop.OutA)
	// 	magnets.Set(col + 0, row + 0, smSwitchTopA)
	// 	smSwitchTopB := dev.NewStaticMagnet(col + 0, row + 0, wSwitchTop.OutB)
	// 	magnets.Set(col + 0, row + 0, smSwitchTopB)
	//
	// 	fmTransInactive := dev.NewFloatMagnet(col + 0, row + 0, dev.StDocked, dev.PolarityNeutral)
	// 	magnets.Set(col + 0, row + 0, fmTransInactive)
	// 	fmTransActive := dev.NewFloatMagnet(col + 0, row + 0, dev.StFloating, dev.PolarityNegative)
	// 	magnets.Set(col + 0, row + 0, fmTransA)
	default:
		panic("undefined level devices")
	}

	return circuits, magnets
}
