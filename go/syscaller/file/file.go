package file

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type FileSyscaller struct {
	f *os.File
}

const fileName = "results/file/test.txt"

func Gen() FileSyscaller {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return FileSyscaller{
		f: f,
	}
}

func (fsc FileSyscaller) Write() {
	defer fsc.Close()

	now := time.Now().String()
	_, err := fsc.f.Write([]byte(now + "\n"))
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	uuid := uuid.New().String()
	_, err = fsc.f.Write([]byte(uuid + "\n"))
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

func (fsc FileSyscaller) Read() {
	defer fsc.Close()

	buf := make([]byte, 1024)
	_, err := fsc.f.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}

func (fsc FileSyscaller) Close() {
	fsc.f.Close()
}
