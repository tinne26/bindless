package iso

//import "math"

// Given a tile column and row, returns the top left tile coordinates.
func TileCoords(col, row int16) (int, int) {
	x := (col*18 + row*18) - 363
	y := (row*9  - col*9 ) + 171
	return int(x), int(y)
}

func YCoord(col, row int16) int {
	return int(row*9  - col*9) + 171
}

// Given logical x and y, returns the tile (col, row).
func CoordsToTile(x, y int) (int16, int16) {
	diag := (y - 1)/18
	x += 363
	row := (x - 18)/36 + (diag - 9)
	col := (x - 18)/36 + (10 - diag)

	// handle corners
	xRem := (x - 18)%36
	yRem := (y - 1)%18
	if xRem < 15 {
		if yRem < (7 - xRem/2) {
			row -= 1
		} else if yRem > (8 + xRem/2) {
			col -= 1
		}
	} else if xRem > 17 {
		if yRem < (7 - (35 - xRem)/2) {
			col += 1
		} else if yRem > (7 + (35 - xRem)/2) {
			row += 1
		}
	}

	return int16(col), int16(row) // int16(col)
}

func TileIndexToKey(col, row int16) int32 {
	return (int32(col) << 16) | int32(row)
}
