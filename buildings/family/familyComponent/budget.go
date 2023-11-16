package buildings

import "fmt"

type Budget struct {
	Money  int
	Credit int
}

func (b Budget) Info() string {
	return fmt.Sprintf("Деньги: %d, Кредит: %d\n", b.Money, b.Credit)
}
