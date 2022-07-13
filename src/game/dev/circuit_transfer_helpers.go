package dev

import imgcolor "image/color"
import "github.com/tinne26/bindless/src/art/palette"

// A TransferProc is a temporary struct to assist in the dock/undock
// synchronization process between two magnets and transfer docks.
type TransferProc struct {
	transfer *TransferSource
	magnet *FloatMagnet
}

func NewTransferProc(transfer *TransferSource, magnet *FloatMagnet) TransferProc {
	return TransferProc { transfer, magnet }
}

func (self TransferProc) OnDockChange(magnet *FloatMagnet) {
	self.transfer.OnDockChange(magnet)
	self.magnet.OnDockChange(magnet)
}

// TransferSource is a helper struct to allow some ephemeral graphical
// effects that require connecting both transferDocks from a single object.
type TransferSource struct {
	transferA *TransferDock
	transferB *TransferDock
	ephemerousPower uint16
	polarity PolarityType
	color imgcolor.RGBA
}

func newTransferSource(a, b *TransferDock) *TransferSource {
	return &TransferSource {
		transferA: a,
		transferB: b,
		polarity: PolarityNeutral,
		color: palette.PolarityNeutral,
	}
}

func (self *TransferSource) Output() (PolarityType, imgcolor.RGBA) {
	return self.polarity, self.color
}

func (self *TransferSource) Update() {
	if self.ephemerousPower > 0 {
		self.ephemerousPower -= 1
		if self.ephemerousPower == 0 {
			self.polarity = PolarityNeutral
			self.color = palette.PolarityNeutral
		} else if self.ephemerousPower % 2 == 0 {
			if self.ephemerousPower < 40 {
				color := self.polarity.Color()
				color.A = uint8((self.ephemerousPower*255)/40)
				color = palette.PreMultAlpha(color)
				self.color = palette.Mix(color, palette.PolarityNeutral)
			} else {
				self.color = self.polarity.Color()
			}
		}
	}
}

func (self *TransferSource) OnDockChange(magnet *FloatMagnet) {
	self.polarity = magnet.Polarity()
	self.ephemerousPower = 26*2
}
