package main

import (
	"fmt"
	"sync"

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

	wg := &sync.WaitGroup{}
	for _, c := range creatures {
		wg.Add(1)
		go active(c, wg)
	}
	wg.Wait()

	fmt.Println("end")
}

func active(creature animal.Creature, wg *sync.WaitGroup) {
	defer wg.Done()

	creature.Run()
	creature.Cry()
}
