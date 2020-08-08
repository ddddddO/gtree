package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

const (
	duration = 1 * time.Second
	fileName = "test.txt"
)

func main() {
	interval := time.NewTicker(duration)
	end := time.NewTicker(duration * 5)

END:
	for {
		select {
		case <-interval.C:
			func() {
				write(fileName)
				read(fileName)
			}()
		case <-end.C:
			break END
		}
	}
}

func write(fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	now := time.Now().String()
	_, err = f.Write([]byte(now + "\n"))
	if err != nil {
		log.Fatal(err)
	}

	uuid := uuid.New().String()
	_, err = f.Write([]byte(uuid + "\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func read(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, info.Size())
	_, err = f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}
