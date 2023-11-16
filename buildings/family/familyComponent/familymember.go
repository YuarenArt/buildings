package buildings

import "fmt"

type FamilyMember struct {
	Name   string
	Age    int
	Gender string
}

func (f FamilyMember) Info() string {
	return fmt.Sprintf("Имя:%s, возраст: %d, пол: %s\n", f.Name, f.Age, f.Gender)
}
