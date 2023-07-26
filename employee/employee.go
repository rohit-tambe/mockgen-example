package employee

import (
	"github.com/rohit-tambe/mockgen-example/calculation"
)

type Employee struct {
	calculator calculation.Calculator
}

func New(c calculation.Calculator) Employee {
	return Employee{calculator: c}
}
func (e *Employee) Add(a int, b int) int {

	result, _ := e.calculator.Add(a, b)
	return result
}
