package level

import "image"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/art/graphics"
import "github.com/tinne26/bindless/src/art/palette"
import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/sound"

type Abilities struct {
	// we use -1 for unavailable. max value is always 4
	Dock int8
	Rewire int8
	Switch int8
	Spectre int8
	Selected uint8 // 0 for none, 1 for dock, 2 rewire, etc.
	Hovered uint8 // 0 for none, 1 for dock, 2 rewire, etc.
	BlinkLeft uint8
	CantSelectLeft uint8
	CantSelectTarget uint8
	key1Pressed bool
	key2Pressed bool
	key3Pressed bool
	key4Pressed bool
}

func (self *Abilities) ConsumeCharge() {
	if self.Selected == 0 {
		panic("can't consume charge for unavailable ability")
	}

	switch self.Selected {
	case 1: self.Dock    -= 1
	case 2: self.Rewire  -= 1
	case 3: self.Switch  -= 1
	case 4: self.Spectre -= 1
	}
}

func (self *Abilities) Update(click bool, logCursorX, logCursorY int) {
	if self.BlinkLeft > 0 { self.BlinkLeft -= 1 }
	if self.CantSelectLeft > 0 { self.CantSelectLeft -= 1 }

	x, y := 10, 360 - 10 - 29
	cx, cy := logCursorX, logCursorY

	target := uint8(0)
	var checkTarget = func(index uint8, charges int8) {
		if charges != -1 {
			if pointWithinRect(cx, cy, x, y, x + 29, y + 29) {
				target = index
			}
			x += 29 + 3
		}
	}

	checkTarget(1, self.Dock)
	checkTarget(2, self.Rewire)
	checkTarget(3, self.Switch)
	checkTarget(4, self.Spectre)

	// added numbers for easier controls
	press1 := ebiten.IsKeyPressed(ebiten.Key1)
	if !press1 { self.key1Pressed = false }
	if press1 && !self.key1Pressed && self.Dock > 0  {
		self.key1Pressed = true
		target = 1
		click = true
	}
	press2 := ebiten.IsKeyPressed(ebiten.Key2)
	if !press2 { self.key2Pressed = false }
	if press2 && !self.key2Pressed && self.Rewire > 0  {
		self.key2Pressed = true
		target = 2
		click = true
	}
	press3 := ebiten.IsKeyPressed(ebiten.Key3)
	if !press3 { self.key3Pressed = false }
	if press3 && !self.key3Pressed && self.Switch > 0  {
		self.key3Pressed = true
		target = 3
		click = true
	}
	press4 := ebiten.IsKeyPressed(ebiten.Key4)
	if !press4 { self.key4Pressed = false }
	if press4 && !self.key4Pressed && self.Spectre > 0 {
		self.key4Pressed = true
		target = 4
		click = true
	}

	// update selected state
	if target != 0 {
		if self.Hovered != target {
			self.Hovered = target
			if !click { sound.PlaySFX(sound.SfxLoudNav) }
		}
		if click {
			if target == self.Selected {
				self.Selected = 0
				sound.PlaySFX(sound.SfxClick)
			} else if self.HasChargesLeft(target) {
				self.Selected = target
				sound.PlaySFX(sound.SfxClick)
			} else {
				sound.PlaySFX(sound.SfxNope)
				self.CantSelectLeft = 6
				self.CantSelectTarget = target
			}
			return
		}
	} else {
		self.Hovered = 0
	}

	if self.Selected != 0 && !self.HasChargesLeft(self.Selected) {
		self.Selected = 0
	}
}

func (self *Abilities) HasChargesLeft(abilityId uint8) bool {
	switch abilityId {
	case 1: return self.Dock    > 0
	case 2: return self.Rewire  > 0
	case 3: return self.Switch  > 0
	case 4: return self.Spectre > 0
	default:
		panic("invalid abilityId")
	}
}

func (self *Abilities) Draw(screen *ebiten.Image) {
	x, y := 10, 360 - 10 - 29

	if self.Dock != -1 {
		self.drawAbility(screen, x, y, 1, self.Dock, graphics.IconDock, graphics.HudDock)
		x += 29 + 3
	}
	if self.Rewire != -1 {
		self.drawAbility(screen, x, y, 2, self.Rewire, graphics.IconRewire, graphics.HudRewire)
		x += 29 + 3
	}
	if self.Switch != -1 {
		self.drawAbility(screen, x, y, 3, self.Switch, graphics.IconSwitch, graphics.HudSwitch)
		x += 29 + 3
	}
	if self.Spectre != -1 {
		self.drawAbility(screen, x, y, 4, self.Spectre, graphics.IconSpectre, graphics.HudSpectre)
		x += 29 + 3
	}
}

func (self *Abilities) DrawWordHint(screen *ebiten.Image, col, row int16, word *ebiten.Image) {
	x, y := iso.TileCoords(col, row)

	w, h := word.Size()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x - w - 4), float64(y - h - 2))
	opts.ColorM.ScaleWithColor(palette.Background)
	screen.DrawImage(word, opts) // top shadow
	opts.GeoM.Translate(0, 2)
	screen.DrawImage(word, opts) // bottom shadow
	opts.GeoM.Translate(-1, -1)
	screen.DrawImage(word, opts) // left shadow
	opts.GeoM.Translate(2, 0)
	screen.DrawImage(word, opts) // right shadow
	opts.GeoM.Translate(-1, 0)
	opts.ColorM.Reset()
	screen.DrawImage(word, opts)
	opts.GeoM.Translate(float64(w), float64(h - 2))
	screen.DrawImage(graphics.HudMsgTail, opts)

}

func (self *Abilities) drawAbility(screen *ebiten.Image, x, y int, selectionIndex uint8, charges int8, icon, word *ebiten.Image) {
	drawWord  := false
	drawColor := palette.AbilityDefault

	// determine style based on selected / hovering icon
	if self.Selected == selectionIndex {
		drawWord = true
		if self.BlinkLeft == 0 {
			drawColor = palette.AbilitySelected
		}
	} else if self.BlinkLeft > 0 && self.Selected == 0 {
		drawColor = palette.AbilitySelected
	} else if self.CantSelectLeft > 0 && self.CantSelectTarget == selectionIndex {
		drawColor = palette.AbilitySelected
	}
	if self.Hovered == selectionIndex { drawWord = true }

	// draw frame
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	opts.ColorM.ScaleWithColor(drawColor)
	screen.DrawImage(graphics.AbilityFrame, opts)

	// draw word
	if drawWord {
		opts.GeoM.Translate(0, -6)
		screen.DrawImage(word, opts)
	}

	// draw charges
	opts.ColorM.Reset()
	opts.GeoM.Reset()
	opts.ColorM.Scale(1.0, 1.0, 1.0, 0.37)
	opts.GeoM.Translate(float64(x + 4), float64(y + 29 - 7))
	screen.DrawImage(graphics.AbilityCharges, opts)
	opts.ColorM.Reset()
	opts.ColorM.ScaleWithColor(palette.AbilityDefault)
	if charges > 0 {
		rect := image.Rect(0, 0, 5*int(charges), 6)
		sub := graphics.AbilityCharges.SubImage(rect).(*ebiten.Image)
		screen.DrawImage(sub, opts)
	}

	// draw icon
	opts.GeoM.Translate(-3, -21)
	screen.DrawImage(icon, opts)
}

func pointWithinRect(x, y, rox, roy, rfx, rfy int) bool {
	return x >= rox && x <= rfx && y >= roy && y <= rfy
}
