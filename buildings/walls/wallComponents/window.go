package buldings

import "fmt"

type Window struct {
	Form string
	Size string
}

func (w Window) Info() string {
	return fmt.Sprintf("Тип окна: %s, размер: %s", w.Form, w.Size)
}
