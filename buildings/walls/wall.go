package buldings

type WallComponents interface {
	Info() string
}

type Wall struct {
	Material   string
	Components []WallComponents
}

func (w Wall) Info() string {
	info := "Материал стен: " + w.Material + "\n"
	info += "Дополнительные компоненты стен:\n"
	for _, component := range w.Components {
		info += component.Info() + "\n"
	}
	return info
}
