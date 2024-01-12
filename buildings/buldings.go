package buildings

import (
	"fmt"
	family "structs/buildings/family"
	furniture "structs/buildings/furnitures"
	wall "structs/buildings/walls"
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

	family1 := family.CreateFamily()
	furniture1 := furniture.CreateFurniture()
	wall1 := wall.CreateWall()

	buildingTmp := Building{Components: []BuildingComponent{family1, furniture1, wall1}}
	return buildingTmp
}
