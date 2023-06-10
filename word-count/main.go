package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func open_file(filename string) (*os.File, error) {

	newFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return newFile, nil
}

func LineCounter(r io.Reader) (int, error) {
	// create a buffer
	buf := make([]byte, 32*1024)
	count := 0
	// create a line sepeator
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
func main() {
	var lFlag = flag.String("l", "", "use to find the line count")
	flag.Parse()
	file, err := open_file(*lFlag)
	if err != nil {
		log.Fatal(err)
	}
	lc, err := LineCounter(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lc, *lFlag)
}
