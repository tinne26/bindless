package graphics

import "github.com/hajimehoshi/ebiten/v2"

// hey, if it works it works.

// background
var BackDecorations [9]*ebiten.Image

// misc
//var Pointer *ebiten.Image
var TileMask *ebiten.Image
var TileBottomLeft *ebiten.Image
var TileBottomRight *ebiten.Image
var CutTileHighlight *ebiten.Image

// hud
var AbilityFrame *ebiten.Image
var AbilityCharges *ebiten.Image
var IconDock *ebiten.Image
var IconRewire *ebiten.Image
var IconSwitch *ebiten.Image
var IconSpectre *ebiten.Image
var HudDock *ebiten.Image
var HudUndock *ebiten.Image
var HudRewire *ebiten.Image
var HudSwitch *ebiten.Image
var HudSpectre *ebiten.Image
var HudMsgTail *ebiten.Image

// magnets!
var MagnetSmall *ebiten.Image
var MagnetSmallFill *ebiten.Image
var MagnetSmallHalo *ebiten.Image
var MagnetSmallShadowAni *ebiten.Image
var MagnetLarge *ebiten.Image
var MagnetLargeFill *ebiten.Image
var MagnetLargeHalo *ebiten.Image
var MagnetLargeFloor *ebiten.Image

// circuit related devices!
var FieldShape *ebiten.Image
var FieldShadow *ebiten.Image
var DockShape *ebiten.Image
var DockShadow *ebiten.Image
var DockFill *ebiten.Image

// wire circuits!
var WireNW2NE [2]*ebiten.Image // first element is the wire, second the shadow
var WireNW2SW [2]*ebiten.Image
var WireNW2SE [2]*ebiten.Image
var WireNE2SW [2]*ebiten.Image
var WireNE2SE [2]*ebiten.Image
var WireSW2SE [2]*ebiten.Image
var Wire3Point *ebiten.Image
