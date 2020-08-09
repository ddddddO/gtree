package syscaller

import (
	"github.com/ddddddO/work/syscaller/file"
)

func Run(sc syscaller) {
	sc.Write()

	if _, ok := sc.(file.FileSyscaller); ok {
		sc = file.Gen()
	}

	sc.Read()
}
