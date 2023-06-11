package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"joshua.com/todo"
)

//const todoFilename = ".todo.json"
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, ""), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be empty")
	}
	return s.Text(), nil
}
func main() {
	var todoFilename = ".todo.json"
	if os.Getenv("TODO_FILENAME") != "" {
		todoFilename = os.Getenv("TODO_FILENAME")
	}
	// task := flag.String("task", "", "Task to be included")
	task := flag.Bool("task", false, "Task to be included")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("Complete", 0, "item to be completed")
	flag.Usage = func() {
		fmt.Println("This program was created by Joshua Jebaraj")
		flag.PrintDefaults()
	}
	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case *list:
		// for _, item := range *l {
		// 	if !item.Done {
		// 		fmt.Println(item.Task)
		// 	}
		fmt.Println(l)

	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)

		}
		if err := l.Save(todoFilename); err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
	case *task:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		if err := l.Save(todoFilename); err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
	}
}
