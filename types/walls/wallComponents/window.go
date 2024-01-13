package buldings

import (
	"fmt"
	shapes "structs/types/shapes"
)

type Window struct {
	Shape shapes.Shape
	Size  string
	Count int
}

func (w Window) Info() string {
	return fmt.Sprintf("Тип окна: \n%s, размер: %s, количество: %d", w.Shape.Info(), w.Size, w.Count)
}
