package water

import (
	"patterns/pkg/cap"
)

type WaterWork interface {
	AddWater(cap *cap.Cap)
}

type water struct {
	volume int
}

func (m *water) AddWater(cap *cap.Cap) {
	cap.Volume += m.volume
}

func New(vol int) (s *water) {
	return &water{volume: vol}
}
