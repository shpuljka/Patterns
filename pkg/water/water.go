package water

import (
	"patterns/pkg/cap"
)

//интрефейс взаимодействия с водой
type WaterWork interface {
	//Добавить воды
	AddWater(cap *cap.Cap)
}

type water struct {
	volume int
}

func (m *water) AddWater(cap *cap.Cap) {
	cap.Volume += m.volume
}

//Создать воду
func NewWater(vol int) (s *water) {
	return &water{volume: vol}
}
