package milk

import (
	"patterns/pkg/cap"
)

//Интерфейс взаимодействия с молоком
type MilkWork interface {
	//Добавить молоко
	AddMilk(cap *cap.Cap)
}

type milk struct {
	volume int
}

func (m *milk) AddMilk(cap *cap.Cap) {
	cap.Volume += m.volume
	cap.IsMilk = true
}

//Создать молоко
func NewMilk(vol int) (s *milk) {
	return &milk{volume: vol}
}
