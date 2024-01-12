package layout

import (
	"fmt"
	"structs/building/layout/room"
)

type LayoutInterface interface {
	Info() string
}

type Layout struct {
	Count int
	Rooms []room.Room
}

func (l Layout) Info() string {
	info := fmt.Sprintf("Количество комнат: %d\n", l.Count)

	for _, room := range l.Rooms {
		info += room.Info()
	}
	return info + "\n"
}

func CreateLayout() LayoutInterface {

	room1 := room.Room{Type: "Гостинная", Square: 25}
	room2 := room.Room{Type: "Кухня", Square: 20}
	room3 := room.Room{Type: "Ванная", Square: 10}
	room4 := room.Room{Type: "Спальня", Square: 15}

	rooms := []room.Room{room1, room2, room3, room4}
	countOfRooms := len(rooms)

	return Layout{
		Count: countOfRooms,
		Rooms: rooms,
	}
}
