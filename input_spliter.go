package gtree

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	md "github.com/ddddddO/gtree/markdown"
)

func split(ctx context.Context, r io.Reader) (<-chan string, <-chan error) {
	sc := bufio.NewScanner(r)
	blockc := make(chan string)
	errc := make(chan error)

	go func() {
		defer func() {
			close(blockc)
			close(errc)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				block := ""
				for sc.Scan() {
					l := sc.Text()
					if isRootBlockBeginning(l) {
						if len(block) != 0 {
							blockc <- block
						}
						block = ""
					}
					block += fmt.Sprintln(l)
				}
				if err := sc.Err(); err != nil {
					errc <- err
					return
				}
				blockc <- block // 最後のRootブロック送出
				return
			}
		}
	}()

	return blockc, errc
}

func isRootBlockBeginning(l string) bool {
	if len(l) == 0 {
		return false
	}
	return strings.ContainsRune(md.ListSymbolsLine, rune(l[0]))
}
