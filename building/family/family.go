package buildings

import (
	"fmt"
	familyComponent "structs/building/family/familyComponent"
)

type FamilyInterface interface {
	Info() string
}

type FamilyComponent interface {
	Info() string
}

type Family struct {
	Surname string
	Members []familyComponent.FamilyMember
	Budget  familyComponent.Budget
}

func (f Family) Info() string {
	info := fmt.Sprintf("Семья: %s\n", f.Surname)
	for _, member := range f.Members {
		info += member.Info()
	}
	info += f.Budget.Info()
	return info + "\n"
}

func CreateFamily() FamilyInterface {
	member1 := familyComponent.FamilyMember{Name: "John", Age: 35, Gender: "Male"}
	member2 := familyComponent.FamilyMember{Name: "Emily", Age: 30, Gender: "Female"}

	return Family{
		Surname: "Smith",
		Members: []familyComponent.FamilyMember{member1, member2},
		Budget:  familyComponent.Budget{Money: 1000000, Credit: 0},
	}
}
