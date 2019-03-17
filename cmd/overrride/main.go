package main

import (
	"github.com/ddddddO/work/cmd/overrride/animal"
)

func main() {
	taro := animal.NewDog("TARO")
	jiro := animal.NewDog("JIRO")

	creatures := []animal.Creature{
		taro,
		jiro,
	}

	for _, c := range creatures {
		active(c)
	}

}

func active(creature animal.Creature) {
	creature.Run()
	creature.Cry()
}
