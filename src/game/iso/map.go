package iso

// generics were very useful for this!

type Map[T any] struct  {
	tiles map[int32]T
}

func NewMap[T any]() Map[T] {
	return Map[T]{
		tiles: make(map[int32]T),
	}
}

func (self *Map[T]) Size() int {
	return len(self.tiles)
}

func (self *Map[T]) Set(x, y int16, value T) {
	key := TileIndexToKey(x, y)
	self.tiles[key] = value
}

func (self *Map[T]) Get(x, y int16) (T, bool) {
	key := TileIndexToKey(x, y)
	value, found := self.tiles[key]
	return value, found
}

func (self *Map[T]) Delete(x, y int16) {
	key := TileIndexToKey(x, y)
	delete(self.tiles, key)
}

func (self *Map[T]) Each(fn func(x, y int16, value T)) {
	for key, value := range self.tiles {
		x := int16((key >> 16) & 0xFFFF)
		y := int16(key & 0xFFFF)
		fn(x, y, value)
	}
}

func (self *Map[T]) SetArea(x, y int16, w, h int, value T) {
	for j := int16(0); j < int16(h); j++ {
		for i := int16(0); i < int16(w); i++ {
			self.Set(x + i, y + j, value)
		}
	}
}

func (self *Map[T]) DeleteArea(x, y int16, w, h int) {
	for j := int16(0); j < int16(h); j++ {
		for i := int16(0); i < int16(w); i++ {
			self.Delete(x + i, y + j)
		}
	}
}
