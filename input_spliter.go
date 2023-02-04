package gtree

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
)

func spliter(ctx context.Context, r io.Reader) (<-chan *bufio.Scanner, <-chan error) {
	sc := bufio.NewScanner(r)
	bsc := make(chan *bufio.Scanner)
	errc := make(chan error)

	go func() {
		defer func() {
			close(bsc)
			close(errc)
		}()

		ret := ""
		for {
			select {
			case <-ctx.Done():
				return
			default:
				for sc.Scan() {
					l := sc.Text()
					if strings.HasPrefix(l, "-") {
						// ここでrootGeneratorにret(空文字列でも)送って大丈夫だっけ。大丈夫だったらそうしたい
						if !(len(ret) == 0) {
							bsc <- bufio.NewScanner(strings.NewReader(ret))
						}

						ret = ""
						ret += fmt.Sprintln(l)
						continue
					}
					ret += fmt.Sprintln(l)
				}

				bsc <- bufio.NewScanner(strings.NewReader(ret)) // 最後のRoot送出

				if err := sc.Err(); err != nil {
					errc <- err
					return
				}
				return
			}
		}
	}()

	return bsc, errc
}
