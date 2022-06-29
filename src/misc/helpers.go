package misc

import "embed"
import "image/png"

import "github.com/hajimehoshi/ebiten/v2"

func LoadPNG(files *embed.FS, path string) (*ebiten.Image, error) {
	file, err := files.Open(path)
	if err != nil { return nil, err }
	img, err := png.Decode(file)
	if err != nil { return nil, err }
	return ebiten.NewImageFromImage(img), file.Close()
}

func ScaledFontSize(size float64, zoomLevel float64) int {
	scaledSize := int(size*zoomLevel)
	if scaledSize <= 0 { return 1 }
	return scaledSize
}

func MousePressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func SkipKeyPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyTab)
}
