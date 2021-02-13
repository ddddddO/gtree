package main

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

func jankenHandler(msg string) string {
	s := strings.ToLower(strings.TrimRight(msg, "\n"))
	var n int
	switch s {
	case "r", "rock":
		n = 0
	case "p", "paper":
		n = 1
	case "s", "scissors":
		n = 2
	default:
		n = 9
	}

	if n == 9 {
		return "INVALID!"
	}

	rand.Seed(time.Now().UnixNano()) // https://pinzolo.github.io/2017/03/28/golang-rand.html
	i := rand.Intn(3)                // 0 ~ 2

	var retMsg string
	switch math.Abs(float64(n) - float64(i)) {
	case 0.0:
		retMsg = "you win!!"
	case 1.0:
		retMsg = "you lose..."
	case 2.0:
		retMsg = "draw"
	}

	return retMsg
}
