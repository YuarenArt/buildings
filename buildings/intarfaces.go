package buildings

type BuildingComponent interface {
	Info() string
}

type FamilyInterface interface {
	Info() string
}

type FurnitureInterface interface {
	Info() string
}

type ShapeInterface interface {
	Info() string
}

type WallInterface interface {
	Info() string
}
