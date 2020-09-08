package water

import (
	"patterns/pkg/model"
)

//интрефейс взаимодействия с водой
type WaterWork interface {
	//Добавить воды
	AddWater(cap *model.Cup)
}

type water struct {
	volume int
}

func (m *water) AddWater(cap *model.Cup) {
	cap.Volume += m.volume
}

//Создать воду
func NewWater(vol int) (s *water) {
	return &water{volume: vol}
}
