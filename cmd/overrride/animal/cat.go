package animal

import (
	"fmt"
)

type Cat struct {
	animal   *Animal
	Favorite string
}

func NewCat(name string) *Cat {
	return &Cat{
		animal: &Animal{
			Name:  name,
			Voice: "nyan",
		},
		Favorite: "Fish",
	}
}

func (c *Cat) Run() {
	fmt.Println(c.animal.Name + " is Cat." + c.animal.Name + " is Run")
}

// NOTE: dog.goと比較すること。(以下をコメントアウトして実行)
//       [*animal.Cat does not implement animal.Creature (missing Cry method)]
func (c *Cat) Cry() {
	fmt.Printf("%s %s! Please %s!!\n", c.animal.Voice, c.animal.Voice, c.Favorite)
}
