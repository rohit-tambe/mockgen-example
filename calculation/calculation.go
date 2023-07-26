package calculation

type Calculator interface {
	Add(int, int) (int, error)
	Substraction(int, int) (int, error)
	Multiplication(int, int) (int, error)
	Addition(value interface{}) error
}

type Calc struct{}

func (c *Calc) Add(a, b int) (int, error) {
	return (a + b), nil
}
