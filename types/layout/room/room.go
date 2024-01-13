package room

import "fmt"

type squareMeter int

type Room struct {
	Type   string
	Square squareMeter
}

func (r Room) Info() string {
	info := fmt.Sprintf("Тип комнаты: %s\n", r.Type)
	info += fmt.Sprintf("Площадь: %d\n", r.Square)
	return info
}
