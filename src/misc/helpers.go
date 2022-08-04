package misc

import "embed"
import "image"
import "image/png"
import "image/color"

import "github.com/hajimehoshi/ebiten/v2"

var miniMask *ebiten.Image
func init() {
	img := ebiten.NewImage(3, 3)
	img.Fill(color.White)
	miniMask = img.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
}

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

func DrawRect(target *ebiten.Image, rgba color.RGBA) {
	bounds := target.Bounds()
	r, g, b, a := rgba.RGBA()
	fr, fg, fb, fa := float32(r)/65535, float32(g)/65535, float32(b)/65535, float32(a)/65535

	vertices := make([]ebiten.Vertex, 4)
	vertices[0].DstX = float32(bounds.Min.X)
	vertices[0].DstY = float32(bounds.Min.Y)
	vertices[1].DstX = float32(bounds.Min.X) // bottom left
	vertices[1].DstY = float32(bounds.Max.Y)
	vertices[2].DstX = float32(bounds.Max.X) // top right
	vertices[2].DstY = float32(bounds.Min.Y)
	vertices[3].DstX = float32(bounds.Max.X) // bottom right
	vertices[3].DstY = float32(bounds.Max.Y)
	for i := 0; i < 4; i++ {
		vertices[i].SrcX = 1.0
		vertices[i].SrcY = 1.0
		vertices[i].ColorR = fr
		vertices[i].ColorG = fg
		vertices[i].ColorB = fb
		vertices[i].ColorA = fa
	}	
	target.DrawTriangles(vertices, []uint16{0, 2, 1, 1, 2, 3}, miniMask, nil)
}