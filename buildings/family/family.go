package buildings

import (
	"fmt"
	familyComponent "structs/buildings/family/familyComponent"
)

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
