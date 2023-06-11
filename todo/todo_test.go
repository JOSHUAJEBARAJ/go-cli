package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"joshua.com/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	l.Add("Learn go")
	if l[0].Task == "Learn Go" {
		t.Errorf("Expected Learn go , got %q instead", l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	l.Add("Learn go")

	if l[0].Done {
		t.Errorf("Task should not be completed")
	}
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("Task should be completed now")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	l.Add("Learn go")
	l.Delete(1)
	if len(l) != 0 {
		t.Errorf("Task should be deleted now")
	}

}

// test to save and get

func TestSaveGet(t *testing.T) {
	// from one list we adding task and saving it and from the other list we are loading and comparing values
	l1 := todo.List{}
	l2 := todo.List{}
	l1.Add("Learn go ")
	// create a temp file
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file")
	}
	defer os.Remove(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving file %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error reading from file")
	}
	if l1[0].Task != l2[0].Task {
		t.Errorf("Both should be matching")
	}

}
