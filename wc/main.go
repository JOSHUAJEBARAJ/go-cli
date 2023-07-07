package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, l bool, b bool) int {
	// read the file from the input
	scanner := bufio.NewScanner(r)

	wc := 0
	// check if l flag is set
	if !l {
		scanner.Split(bufio.ScanWords)
	}
	if b {
		scanner.Split(bufio.ScanBytes)
	}
	for scanner.Scan() {
		wc++
	}
	return wc
}
func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count,bytes")
	flag.Parse()
	count := Count(os.Stdin, *lines, *bytes)
	fmt.Println(count)
}
