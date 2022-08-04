package graphics

import "embed"

import "github.com/tinne26/bindless/src/misc"

// loads most of the graphical assets of the game so they are
// directly available from package-level variables as ebiten images.
// this is probably illegal except on game-jam jurisdiction.
// you should use a cache or something like that instead.
func Load(files *embed.FS) error {
	var err error

	// load background decorations
	baseDir := "assets/graphics/background/"
	BackDecorations[0], err = misc.LoadPNG(files, baseDir + "small_top.png")
	if err != nil { return err }
	BackDecorations[1], err = misc.LoadPNG(files, baseDir + "small_right.png")
	if err != nil { return err }
	BackDecorations[2], err = misc.LoadPNG(files, baseDir + "small_left.png")
	if err != nil { return err }
	BackDecorations[3], err = misc.LoadPNG(files, baseDir + "medium_top.png")
	if err != nil { return err }
	BackDecorations[4], err = misc.LoadPNG(files, baseDir + "medium_right.png")
	if err != nil { return err }
	BackDecorations[5], err = misc.LoadPNG(files, baseDir + "medium_left.png")
	if err != nil { return err }
	BackDecorations[6], err = misc.LoadPNG(files, baseDir + "large_top.png")
	if err != nil { return err }
	BackDecorations[7], err = misc.LoadPNG(files, baseDir + "large_right.png")
	if err != nil { return err }
	BackDecorations[8], err = misc.LoadPNG(files, baseDir + "large_left.png")
	if err != nil { return err }

	// misc graphics
	baseDir = "assets/graphics/misc/"
	TileMask, err = misc.LoadPNG(files, baseDir + "tile_mask.png")
	if err != nil { return err }
	TileBottomLeft, err = misc.LoadPNG(files, baseDir + "tile_bottom_left.png")
	if err != nil { return err }
	TileBottomRight, err = misc.LoadPNG(files, baseDir + "tile_bottom_right.png")
	if err != nil { return err }
	CutTileHighlight, err = misc.LoadPNG(files, baseDir + "cut_tile_highlight.png")
	if err != nil { return err }
	// Pointer, err = misc.LoadPNG(files, baseDir + "pointer.png")
	// if err != nil { return err }

	// hud
	baseDir = "assets/graphics/hud/"
	AbilityFrame, err = misc.LoadPNG(files, baseDir + "ability_frame.png")
	if err != nil { return err }
	AbilityCharges, err = misc.LoadPNG(files, baseDir + "charges.png")
	if err != nil { return err }
	IconDock, err = misc.LoadPNG(files, baseDir + "icon_dock.png")
	if err != nil { return err }
	IconRewire, err = misc.LoadPNG(files, baseDir + "icon_rewire.png")
	if err != nil { return err }
	IconSwitch, err = misc.LoadPNG(files, baseDir + "icon_switch.png")
	if err != nil { return err }
	IconSpectre, err = misc.LoadPNG(files, baseDir + "icon_spectre.png")
	if err != nil { return err }
	IconMenu, err = misc.LoadPNG(files, baseDir + "icon_menu.png")
	if err != nil { return err }
	baseDir = "assets/graphics/hud/msg/"
	HudDock, err = misc.LoadPNG(files, baseDir + "dock.png")
	if err != nil { return err }
	HudUndock, err = misc.LoadPNG(files, baseDir + "undock.png")
	if err != nil { return err }
	HudRewire, err = misc.LoadPNG(files, baseDir + "rewire.png")
	if err != nil { return err }
	HudSwitch, err = misc.LoadPNG(files, baseDir + "switch.png")
	if err != nil { return err }
	HudSpectre, err = misc.LoadPNG(files, baseDir + "spectre.png")
	if err != nil { return err }
	HudMenu, err = misc.LoadPNG(files, baseDir + "menu.png")
	if err != nil { return err }
	HudMsgTail, err = misc.LoadPNG(files, baseDir + "tail.png")
	if err != nil { return err }

	// load magnets
	baseDir = "assets/graphics/magnets/"
	MagnetSmall, err = misc.LoadPNG(files, baseDir + "small.png")
	if err != nil { return err }
	MagnetSmallFill, err = misc.LoadPNG(files, baseDir + "small_fill.png")
	if err != nil { return err }
	MagnetSmallHalo, err = misc.LoadPNG(files, baseDir + "small_halo.png")
	if err != nil { return err }
	MagnetSmallShadowAni, err = misc.LoadPNG(files, baseDir + "small_shadow_animation.png")
	if err != nil { return err }
	MagnetLarge, err = misc.LoadPNG(files, baseDir + "large.png")
	if err != nil { return err }
	MagnetLargeFill, err = misc.LoadPNG(files, baseDir + "large_fill.png")
	if err != nil { return err }
	MagnetLargeHalo, err = misc.LoadPNG(files, baseDir + "large_halo.png")
	if err != nil { return err }
	MagnetLargeFloor, err = misc.LoadPNG(files, baseDir + "large_floor.png")
	if err != nil { return err }

	// load circuit-related devices
	baseDir = "assets/graphics/circuits/"
	FieldShape, err = misc.LoadPNG(files, baseDir + "field.png")
	if err != nil { return err }
	FieldShadow, err = misc.LoadPNG(files, baseDir + "field_shadows.png")
	if err != nil { return err }
	DockShape, err = misc.LoadPNG(files, baseDir + "dock.png")
	if err != nil { return err }
	DockShadow, err = misc.LoadPNG(files, baseDir + "dock_shadows.png")
	if err != nil { return err }
	DockFill, err = misc.LoadPNG(files, baseDir + "dock_fill.png")
	if err != nil { return err }
	TargetShape, err = misc.LoadPNG(files, baseDir + "target.png")
	if err != nil { return err }
	TargetShadow, err = misc.LoadPNG(files, baseDir + "target_shadows.png")
	if err != nil { return err }

	// create wire images manually
	loadWires()

	return nil
}
