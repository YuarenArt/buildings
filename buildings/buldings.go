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

func MakeBuilding() Building {

	member1 := familyComponent.FamilyMember{Name: "John", Age: 35, Gender: "Male"}
	member2 := familyComponent.FamilyMember{Name: "Emily", Age: 30, Gender: "Female"}

	family1 := family.Family{
		Surname: "Smith",
		Members: []familyComponent.FamilyMember{member1, member2},
		Components: []family.FamilyComponent{
			familyComponent.Budget{
				Money:  100000,
				Credit: 0,
			},
		},
	}

	item1 := furnitureComponents.Item{Name: "Стул", Material: "Дерево"}
	item2 := furnitureComponents.Item{Name: "Стол", Material: "Стекло"}
	item3 := furnitureComponents.Item{Name: "Диван", Material: "Ткань"}

	furniture1 := furniture.Furniture{Items: []furniture.FurnitureComponents{item1, item2, item3}}

	shape1 := shapeComponents.TypeOfShape{Type: "Прямоугльник"}
	shape2 := shapeComponents.TypeOfShape{Type: "Круг"}
	shape3 := shapeComponents.TypeOfShape{Type: "Треугольник"}

	shapes := shape.Shape{Components: []shape.ShapeComponents{shape1, shape2, shape3}}

	component1 := wallComponents.Window{Form: "Круглый", Size: "Большой"}
	component2 := wallComponents.Door{Material: "Дерево", Size: "Большой"}

	wall1 := wall.Wall{
		Material:   "Кирпич",
		Components: []wall.WallComponents{component1, component2},
	}

	buildingTmp := Building{Components: []BuildingComponent{family1, furniture1, shapes, wall1}}
	return buildingTmp
}
