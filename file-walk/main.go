package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ext  string
	size int64
	list bool
	del  bool
	wLog io.Writer
}

func main() {
	root := flag.String("root", ".", "Root directory")
	list := flag.Bool("list", false, "List Files")
	ext := flag.String("ext", "", "File")
	size := flag.Int64("size", 0, "Minium file")
	del := flag.Bool("del", false, "delete the matched file")
	logFile := flag.String("log", "", "Log deletes to this file")
	flag.Parse()
	var (
		f   = os.Stdout
		err error
	)
	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}
	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
		del:  *del,
		wLog: f,
	}
	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)

	}

}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE", log.LstdFlags)
	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			// check the error while executing the walk
			if err != nil {
				return err
			}
			// check whether it needs to be skipped or not
			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}
			// if list is specified don't do anything else
			if cfg.list {
				return listFile(path, out)
			}
			if cfg.del {
				return delFile(path, delLogger)
			}
			return listFile(path, out)
		})
}
