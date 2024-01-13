package buildings

type ShapeInterface interface {
	Info() string
}

type ShapeComponentsInterface interface {
	Info() string
}

type Shape struct {
	Components []ShapeComponentsInterface
}

func (s Shape) Info() string {
	info := "Информация о форме:\n"
	for _, component := range s.Components {
		info += component.Info()
	}
	return info
}
