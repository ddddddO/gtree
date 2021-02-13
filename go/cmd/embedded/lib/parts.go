package lib

import "fmt"

type Parts interface {
	Fetch() error
	Update([]interface{})
}

func Run(p Parts) {
	if err := p.Fetch(); err != nil {
		panic(err)
	}

	s := []interface{}{"A", "B", "C"}
	p.Update(s)
}

type Default struct{}

func (d Default) Fetch() error {
	fmt.Println("Default Fetch")
	return nil
}

func (d Default) Update(vs []interface{}) {
	fmt.Println("Default Update")

	for _, v := range vs {
		fmt.Println(v)
	}
}
