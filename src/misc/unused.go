//go:build nope

// // draw mouse pointer. the pointer is special in the sense that it's
// // a pixel-art asset but we draw it on the hiResScreen to stay above text
// opts = &ebiten.DrawImageOptions{}
// opts.GeoM.Translate(-2, -2)
// opts.GeoM.Scale(xScale, yScale)
// cx, cy := ebiten.CursorPosition()
// opts.GeoM.Translate(float64(cx), float64(cy))
// hiResScreen.DrawImage(graphics.Pointer, opts)
