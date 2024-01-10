package buldings

import "fmt"

type Door struct {
	Material string
	Size     string
	Count    int
}

func (d Door) Info() string {
	return fmt.Sprintf("Материал двери: %s, размер: %s, количество: %d\n", d.Material, d.Size, d.Count)
}
