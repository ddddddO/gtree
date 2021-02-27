package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/width"
)

var whiteList []string

func main() {
	whitelist := ""
	flag.StringVar(&whitelist, "wl", "", "半角に変換したくない単語を指定します(e.g. --wl=カエル,トラック)")
	flag.Parse()

	whiteList = strings.Split(whitelist, ",")

	t := ""
	_, err := fmt.Fscan(os.Stdin, &t)
	if err != nil {
		panic(err)
	}

	tmp := replaceByWhiteList(t)
	converted := width.Narrow.String(tmp)
	output := reverse(converted)

	fmt.Println(output)
}

func replaceByWhiteList(t string) string {
	replaced := t
	for i := range whiteList {
		i_str := strconv.Itoa(i)
		before := whiteList[i]
		after := "{placeholder_" + i_str + "}"

		replaced = strings.Replace(replaced, before, after, -1)
	}
	return replaced
}

func reverse(t string) string {
	reversed := t

	for i := range whiteList {
		i_str := strconv.Itoa(i)
		before := "{placeholder_" + i_str + "}"
		after := whiteList[i]

		reversed = strings.Replace(reversed, before, after, -1)
	}
	return reversed
}
