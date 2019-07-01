package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"

	"github.com/goburrow/serial"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("few arguments: need device path and file path")
	}

	fi, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	port, err := serial.Open(&serial.Config{Address: os.Args[1]})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	reader := bufio.NewReaderSize(fi, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fi.Seek(0, 0)
		} else if err != nil {
			log.Fatal(err)
		}

		port.Write(line)
		port.Write([]byte("\n"))

		time.Sleep(time.Second)
	}
}
