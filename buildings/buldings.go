package buildings

import (
	"fmt"
	shapes "structs/buildings/shapes"
	shapeComponents "structs/buildings/shapes/shapeComponents"

	family "structs/buildings/family"
	familyComponent "structs/buildings/family/familyComponent"

	furniture "structs/buildings/furnitures"
	furnitureComponents "structs/buildings/furnitures/furnitureComponents"

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
		fmt.Printf("%s\n", component.Info())
	}
}

func MakeBuilding() Building {

	family1 := createFamily()
	furniture1 := createFurniture()
	wall1 := createWall()

	buildingTmp := Building{Components: []BuildingComponent{family1, furniture1, wall1}}
	return buildingTmp
}

func createFamily() FamilyInterface {
	member1 := familyComponent.FamilyMember{Name: "John", Age: 35, Gender: "Male"}
	member2 := familyComponent.FamilyMember{Name: "Emily", Age: 30, Gender: "Female"}

	return family.Family{
		Surname: "Smith",
		Members: []familyComponent.FamilyMember{member1, member2},
		Budget:  familyComponent.Budget{Money: 1000000, Credit: 0},
	}
}

func createFurniture() FurnitureInterface {
	item1 := furnitureComponents.Item{Name: "Стул", Material: "Дерево"}
	item2 := furnitureComponents.Item{Name: "Стол", Material: "Стекло"}
	item3 := furnitureComponents.Item{Name: "Диван", Material: "Ткань"}

	return furniture.Furniture{Items: []furniture.FurnitureComponentsInterface{item1, item2, item3}}
}

func createWall() WallInterface {

	shapeTypeRectangle := shapeComponents.TypeOfShape{Type: "Прямоугольник"}
	shapeForWindow1 := shapes.Shape{Components: []shapes.ShapeComponents{shapeTypeRectangle}}

	component1 := wallComponents.Window{Shape: shapeForWindow1, Size: "Большой"}
	component2 := wallComponents.Door{Material: "Дерево", Size: "Большой"}

	return wall.Wall{
		Material:   "Кирпич",
		Components: []wall.WallComponents{component1, component2},
	}
}
