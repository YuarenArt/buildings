package buldings

import "fmt"

type Door struct {
	Material string
	Size     string
}

func (d Door) Info() string {
	return fmt.Sprintf("Материал двери: %s, размер: %s", d.Material, d.Size)
}
