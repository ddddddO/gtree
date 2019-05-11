package animal

import (
	"fmt"
)

type Dog struct {
	*Animal
	Favorite string
}

// ref: https://anatanoa.blogspot.com/2015/08/blog-post_18.html
func NewDog(name string) *Dog {
	return &Dog{
		Animal: &Animal{
			Name:  name,
			Voice: "Wan",
		},
		Favorite: "Doggiman",
	}

	/*	return &Dog{
			&Animal{
				Name:  name,
				Voice: "Wan",
			},
		}
	*/
}

func (d *Dog) Run() {
	//fmt.Println(d.Name + " is DOG." + d.Name + " is Run.")
	fmt.Println(d.Name + " is DOG." + d.Name + " is Run. He eats " + d.Favorite + " after running.")
}
