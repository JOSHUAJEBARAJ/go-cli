package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	op := flag.String("op", "sum", "operation to be executed")
	col := flag.Int("col", 1, "Column on which operation to be executed")
	flag.Parse()

	if err := run(flag.Args(), *op, *col, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filenames []string, op string, col int, out io.Writer) error {
	var opFunc statsFunc
	if len(filenames) == 0 {
		return ErrNoFiles
	}
	if col < 1 {
		return fmt.Errorf("%w:%d", ErrInvalidColum, col)
	}
	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}
	consolidate := make([]float64, 0)
	for _, fname := range filenames {
		f, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("cannot open file:%w", err)
		}
		data, err := csv2float(f, col)
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
		consolidate = append(consolidate, data...)
	}
	_, err := fmt.Println(out, opFunc(consolidate))
	return err
}
