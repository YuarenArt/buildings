package buildings

import "fmt"

type Item struct {
	Name     string
	Material string
	Count    int
}

func (i Item) Info() string {
	return fmt.Sprintf("Название: %s, Материал: %s, Количество: %d\n", i.Name, i.Material, i.Count)
}
