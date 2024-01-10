package buildings

type FurnitureComponentsInterface interface {
	Info() string
}

type Furniture struct {
	Items []FurnitureComponentsInterface
}

func (f Furniture) Info() string {
	info := "Информация о мебели:\n"
	for _, item := range f.Items {
		info += item.Info()
	}
	return info + "\n"
}
