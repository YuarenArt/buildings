package building

import (
	"fmt"
	family "structs/types/family"
	furniture "structs/types/furnitures"
	"structs/types/layout"
	wall "structs/types/walls"
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

func CreateBuilding() Building {

	family1 := family.MakeFamily()
	furniture1 := furniture.CreateFurniture()
	wall1 := wall.CreateWall()
	layout1 := layout.CreateLayout()

	buildingTmp := Building{Components: []BuildingComponent{layout1, family1, furniture1, wall1}}
	return buildingTmp
}
