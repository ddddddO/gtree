package main

import (
	"fmt"
)

type NumError struct {
	Name string
	No   int
}

func (e *NumError) Error() string {
	return fmt.Sprintf("NUBER ERROR, Name: %s, No: %d", e.Name, e.No)
}

func main() {
	is := []int{0, 1, 2, 3}
	for _, i := range is {
		if err := echo(i); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("No: %d\n", i)
		}
	}

}

func echo(i int) error {
	switch i {
	case 0:
		return nil
	case 1:
		return &NumError{Name: "ONE", No: 1}
	case 2:
		return &NumError{Name: "TWO", No: 2}
	default:
		return nil
	}
}
