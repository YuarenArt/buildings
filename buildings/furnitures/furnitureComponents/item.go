package buildings

import "fmt"

type Item struct {
	Name     string
	Material string
}

func (i Item) Info() string {
	return fmt.Sprintf("Название: %s, Материал: %s\n", i.Name, i.Material)
}
