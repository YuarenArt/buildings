package buildings

import (
	"fmt"

	family "structs/buildings/family"
	familyComponent "structs/buildings/family/familyComponent"

	furniture "structs/buildings/furnitures"
	furnitureComponents "structs/buildings/furnitures/furnitureComponents"

	shape "structs/buildings/shapes"
	shapeComponents "structs/buildings/shapes/shapeComponents"

	wall "structs/buildings/walls"
	wallComponents "structs/buildings/walls/wallComponents"
)

type BuildingComponent interface {
	Info() string
}

type Building struct {
	Components []BuildingComponent
}

func (b Building) Info() {
	fmt.Printf("Информация о здании:\n\n")
	for _, component := range b.Components {
		fmt.Printf("%s", component.Info())
	}
}

func createFamily() family.Family {
	member1 := familyComponent.FamilyMember{Name: "John", Age: 35, Gender: "Male"}
	member2 := familyComponent.FamilyMember{Name: "Emily", Age: 30, Gender: "Female"}

	return family.Family{
		Surname:    "Smith",
		Members:    []familyComponent.FamilyMember{member1, member2},
		Components: []family.FamilyComponent{familyComponent.Budget{Money: 100000, Credit: 0}},
	}
}

func createFurniture() furniture.Furniture {
	item1 := furnitureComponents.Item{Name: "Стул", Material: "Дерево"}
	item2 := furnitureComponents.Item{Name: "Стол", Material: "Стекло"}
	item3 := furnitureComponents.Item{Name: "Диван", Material: "Ткань"}

	return furniture.Furniture{Items: []furniture.FurnitureComponents{item1, item2, item3}}
}

func createShapes() shape.Shape {
	shape1 := shapeComponents.TypeOfShape{Type: "Прямоугльник"}
	shape2 := shapeComponents.TypeOfShape{Type: "Круг"}
	shape3 := shapeComponents.TypeOfShape{Type: "Треугольник"}

	return shape.Shape{Components: []shape.ShapeComponents{shape1, shape2, shape3}}
}

func createWall() wall.Wall {
	component1 := wallComponents.Window{Form: "Круглый", Size: "Большой"}
	component2 := wallComponents.Door{Material: "Дерево", Size: "Большой"}

	return wall.Wall{
		Material:   "Кирпич",
		Components: []wall.WallComponents{component1, component2},
	}
}

func MakeBuilding() Building {
	family1 := createFamily()
	furniture1 := createFurniture()
	shapes := createShapes()
	wall1 := createWall()

	buildingTmp := Building{Components: []BuildingComponent{family1, furniture1, shapes, wall1}}
	return buildingTmp
}
