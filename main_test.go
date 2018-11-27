package main

import (
	"testing"
)
var list LinkedList

func TestAppend(t *testing.T) {
	if !list.IsEmpty() {
		t.Errorf("List should be empyty")
	}

	list.Append(0)
	if size := list.Size(); size != 1 {
		t.Errorf("Wrong size. Expected 1 and got %d", size)
	}
	list.Append(1)
	list.Append(2)

	if size := list.Size(); size != 3 {
		t.Errorf("Wrong size. Expected 3 and got %d", size)
	}
}

func TestRemoveAt(t *testing.T) {

	_, err := list.RemoveAt(2)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	if size := list.Size(); size != 2 {
		t.Errorf("Wrong size. Expected 2 and got %d", size)
	}
}

func TestInsertAt(t *testing.T) {
	err := list.InsertAt(2, 2)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	if size := list.Size(); size != 3 {
		t.Errorf("Wrong size. Expected 3 and got %d", size)
	}
}

func TestIndexOf(t *testing.T) {
	if i := list.IndexOf(0); i != 0 {
		t.Errorf("expected position 0 but got %d", i)
	}

	if i := list.IndexOf(1); i != 1 {
		t.Errorf("expected position 1 but got %d", i)
	}
}
