package main

import (
	"bufio"
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

func lineCounter(r io.Reader) (int, error) {
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

func wordCounter(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	count := 0
	// split through words
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count, nil
}

func charCounter(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	count := 0
	// split through words
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		count++
	}
	return count, nil
}
func main() {
	if len(os.Args) < 1 {
		fmt.Print("TBD")
	}
	var lFlag = flag.Bool("l", false, "use to find the line count")
	flag.Parse()
	_ = lFlag
	fmt.Println(os.Args[1])
}

func checkError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
