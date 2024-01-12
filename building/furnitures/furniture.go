package buildings

import furnitureComponents "structs/building/furnitures/furnitureComponents"

type FurnitureComponentsInterface interface {
	Info() string
}

type FurnitureInterface interface {
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

func CreateFurniture() FurnitureInterface {
	item1 := furnitureComponents.Item{Name: "Стул", Material: "Дерево"}
	item2 := furnitureComponents.Item{Name: "Стол", Material: "Стекло"}
	item3 := furnitureComponents.Item{Name: "Диван", Material: "Ткань"}

	return Furniture{Items: []FurnitureComponentsInterface{item1, item2, item3}}
}
