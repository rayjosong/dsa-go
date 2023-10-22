package main

import "testing"

func TestLinkedList(t *testing.T) {
	myList := NewLinkedList()

	myList.Append(1)
	if myList.Len() != 1 {
		t.Errorf("Expected length 1, got %d", myList.Len())
	}

	myList.Insert(2, 1)
	if myList.Len() != 2 {
		t.Errorf("Expected length 2, got %d", myList.Len())
	}

	err := myList.Delete(1)
	if err != nil {
		t.Errorf("Error deleting: %s", err)
	}

	if myList.Len() != 1 {
		t.Errorf("Expected length 1, got %d", myList.Len())
	}
}
