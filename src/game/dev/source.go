package dev

type staticSource struct { polarity PolarityType }
func (self staticSource) Polarity() PolarityType { return self.polarity }
func NewStaticSource(polarity PolarityType) Polarized {
	return staticSource{ polarity }
}
