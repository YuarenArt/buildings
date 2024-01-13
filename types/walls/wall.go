package buldings

import (
	shapes "structs/types/shapes"
	shapeComponents "structs/types/shapes/shapeComponents"
	wallComponents "structs/types/walls/wallComponents"
)

type WallInterface interface {
	Info() string
}

type WallComponentsInterface interface {
	Info() string
}

type Wall struct {
	Material   string
	Components []WallComponentsInterface
}

func (w Wall) Info() string {
	info := "Материал стен: " + w.Material + "\n"
	info += "Дополнительные компоненты стен:\n"
	for _, component := range w.Components {
		info += component.Info() + "\n"
	}
	return info
}

func CreateWall() WallInterface {

	shapeTypeRectangle := shapeComponents.TypeOfShape{Type: "Прямоугольник"}
	shapeForWindow1 := shapes.Shape{Components: []shapes.ShapeComponentsInterface{shapeTypeRectangle}}

	component1 := wallComponents.Window{Shape: shapeForWindow1, Size: "Большой"}
	component2 := wallComponents.Door{Material: "Дерево", Size: "Большой"}

	return Wall{
		Material:   "Кирпич",
		Components: []WallComponentsInterface{component1, component2},
	}
}
