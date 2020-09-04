package milk

import (
	"patterns/pkg/cap"
)

type MilkWork interface {
	AddMilk(cap *cap.Cap)
}

type milk struct {
	volume int
}

func (m *milk) AddMilk(cap *cap.Cap) {
	cap.Volume += m.volume
	cap.IsMilk = true
}

func New(vol int) (s *milk) {
	return &milk{volume: vol}
}
