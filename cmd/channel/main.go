package main

import (
	"fmt"
)

func go1(ch chan int, params ...int) {
	sum := 0
	for _, i := range params {
		sum += i
	}
	ch <- sum
}

func go2(ch chan int, is []int) {
	sum := 0
	for _, i := range is {
		sum += i
		ch <- sum
	}
	close(ch) // NOTE: go2から、呼び出しもとでforで受信するため、channelをここで閉じる。
}

func main() {
	//try1()
	//try2()
	//try3()
	//try4()
	//try5()
	//try6()
	//try7()
	try8()
}

func try1() {
	ch := make(chan int)
	go go1(ch, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	go go1(ch, 11, 12, 13, 14, 15, 16, 17)

	sum1 := <-ch
	fmt.Println(sum1)

	sum2 := <-ch
	fmt.Println(sum2)

	/*
		sum3 := <-ch
		fmt.Println(sum3)
	*/
}

// buffered channels
func try2() {
	ch := make(chan int, 2)
	fmt.Println(len(ch))
	ch <- 100
	_ = <-ch
	ch <- 200
	ch <- 300
	fmt.Println(len(ch))
	fmt.Println()

	close(ch) // NOTE: channelの終了。これがないと、下のforで存在しない３つ目のchannelを取り出そうとしてエラー
	// NOTE: 予めバッファを指定していてもcloseしないと、3つ目取りに行ってしまう
	for c := range ch {
		fmt.Println(c)
	}
	//ch <- 400 // NOTE: chは閉じられているのでエラー
}

func try3() {
	ch := make(chan int)
	is := []int{1, 1, 1, 1, 1, 1}
	go go2(ch, is)
	for c := range ch {
		fmt.Println(c)
	}
}

// Producer/Consumer
func try4() {
	run()
}

// fan-out fan-in
func try5() {
	runFan()
}

// select
func try6() {
	runSel()
}

func try7() {
	runTic()
}

func try8() {
	runMu()
}
