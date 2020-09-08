package milk

import (
	"patterns/pkg/model"
)

//Интерфейс взаимодействия с молоком
type MilkWork interface {
	//Добавить молоко
	AddMilk(cap *model.Cup)
}

type milk struct {
	volume int
}

func (m *milk) AddMilk(cap *model.Cup) {
	cap.Volume += m.volume
	cap.IsMilk = true
}

//Создать молоко
func NewMilk(vol int) (s *milk) {
	return &milk{volume: vol}
}
