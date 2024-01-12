package buldings

import (
	shapes "structs/building/shapes"
	shapeComponents "structs/building/shapes/shapeComponents"
	wallComponents "structs/building/walls/wallComponents"
)

type WallInterface interface {
	Info() string
}

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

func CreateWall() WallInterface {

	shapeTypeRectangle := shapeComponents.TypeOfShape{Type: "Прямоугольник"}
	shapeForWindow1 := shapes.Shape{Components: []shapes.ShapeComponents{shapeTypeRectangle}}

	component1 := wallComponents.Window{Shape: shapeForWindow1, Size: "Большой"}
	component2 := wallComponents.Door{Material: "Дерево", Size: "Большой"}

	return Wall{
		Material:   "Кирпич",
		Components: []WallComponents{component1, component2},
	}
}
