package gtree

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
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
					if strings.HasPrefix(l, "-") {
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