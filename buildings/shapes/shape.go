package buildings

type ShapeComponents interface {
	Info() string
}

type Shape struct {
	Components []ShapeComponents
}

func (s Shape) Info() string {
	info := "Информация о формах:\n"
	for _, component := range s.Components {
		info += component.Info()
	}
	return info + "\n"
}
