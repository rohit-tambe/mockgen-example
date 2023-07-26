package party

import "fmt"

type Visitor struct {
	Name    string
	Surname string
}

func (v Visitor) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Surname)
}
