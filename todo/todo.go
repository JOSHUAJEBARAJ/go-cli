package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type todo struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List is the name type of todo slices
type List []todo

func (l *List) Add(task string) {
	t := todo{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)

}
func (l *List) Complete(i int) error {
	// check the user input
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("Item doesnt exist")
	}
	ls := *l
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("Item doesnt exist")
	}
	// add the before value and add the after value

	*l = append(ls[:i-1], ls[i:]...)
	return nil
}
func (l *List) String() string {
	formated := ""
	for k, t := range *l {
		prefix := ""
		if t.Done {
			prefix = "X"
		}
		formated += fmt.Sprintf("%s %d: %s\n", prefix, k+1, t.Task)
	}
	return formated
}
func (l *List) Save(file string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, js, 0644)
}

func (l *List) Get(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		// check if the file is present or not
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(content, l)
}
