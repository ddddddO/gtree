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
	strc := make(chan string)
	errc := make(chan error)

	go func() {
		defer func() {
			close(strc)
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
						if !(len(ret) == 0) {
							strc <- ret
						}

						ret = ""
						ret += fmt.Sprintln(l)
						continue
					}
					ret += fmt.Sprintln(l)
				}

				strc <- ret // 最後のRoot送出
				errc <- sc.Err()
				return
			}
		}
	}()

	return strc, errc
}
