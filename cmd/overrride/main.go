package main

import (
	"github.com/ddddddO/work/cmd/overrride/animal"
)

func main() {
	taro := animal.NewDog("TARO")
	jiro := animal.NewDog("JIRO")
	mya := animal.NewCat("Mya-mya-")

	creatures := []animal.Creature{
		taro,
		jiro,
		mya,
	}

	for _, c := range creatures {
		active(c)
	}

}

func active(creature animal.Creature) {
	creature.Run()
	creature.Cry()
}
