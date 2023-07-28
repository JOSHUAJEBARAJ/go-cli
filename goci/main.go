package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type executer interface {
	execute() (string, error)
}

func run(proj string, out io.Writer) error {

	if proj == "" {
		return fmt.Errorf("Project is required")
	}
	// args

	// args := []string{"build", ".", "errors"}

	// cmd := exec.Command("go", args...)
	// // set the working directory
	// cmd.Dir = proj
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("go build failed %s", err)
	// }
	// _, err := fmt.Fprintln(out, "Build was success")
	// return err

	pipeline := make([]executer, 3)
	pipeline[0] = newStep(
		"go build",
		"go",
		"Build was success\n",
		proj,
		[]string{"build", ".", "errors"},
	)

	for _, s := range pipeline {
		msg, err := s.execute()
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err
		}
	}
	return nil

}

func main() {
	proj := flag.String("p", "", "Project directory")
	flag.Parse()
	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
