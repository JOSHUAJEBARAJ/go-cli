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
func main() {
	var lFlag = flag.String("l", "", "use to find the line count")
	flag.Parse()
	file, err := open_file(*lFlag)
	checkError(err)
	lc, err := lineCounter(file)
	checkError(err)
	// moving the cursor to the begining
	file.Seek(0, io.SeekStart)
	fmt.Println(lc, *lFlag)
	wc, err := wordCounter(file)
	checkError(err)
	println(wc)

}

func checkError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
