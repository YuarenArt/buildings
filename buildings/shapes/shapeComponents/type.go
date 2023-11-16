package buildings

type TypeOfShape struct {
	Type string
}

func (t TypeOfShape) Info() string {
	return "Тип формы: " + t.Type + "\n"
}
