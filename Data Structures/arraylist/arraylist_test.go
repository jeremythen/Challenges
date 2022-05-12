package arraylist_test

import (
	"fmt"
	"testing"

	"example.com/challenge/data-structure/arraylist"
)

func TestArrayLstCreation(t *testing.T) {
	myList := arraylist.CreateList()
	for _, value := range myList.Print() {
		if value != nil {
			t.Fatalf("New list with non-nil element")
		}
	}
	if myList.Size() != 0 {
		fmt.Println(myList.Size())
		t.Fatalf("New list with non-zero length.")
	}
}

func TestArrayListAdd(t *testing.T) {
	myList := arraylist.CreateList()
	myList.Add("a")
	if myList.IsEmpty() {
		t.Fatal("List empty after add")
	}
	if myList.Print()[0] != "a" {
		t.Fatalf("First element does not match added element")
	}
}

func TestArrayListGet(t *testing.T) {
	myList := arraylist.CreateList()
	myList.Add("a")
	myList.Add("b")
	myList.Add("c")
	myList.Add("d")
	value1, error1 := myList.Get(2)
	if value1 != "c" || error1 != nil {
		t.Fatal("Method get should return corresponding value with valid index")
	}
	value2, error2 := myList.Get(10)
	if value2 != nil || error2 == nil {
		t.Fatal("Method get should return nil and error when out of bounds")
	}
}

func TestArrayListRemoveIndex(t *testing.T) {
	myList := arraylist.CreateList()
	myList.Add("a")
	myList.Add("b")
	myList.Add("c")
	myList.Add("d")
	removed1, error1 := myList.RemoveIndex(2)
	if !removed1 || error1 != nil {
		t.Fatal("Method remove should return true with valid index")
	}
	removed2, error2 := myList.RemoveIndex(10)
	if removed2 || error2 == nil {
		t.Fatal("Method remove should return false and error when out of bounds")
	}
}

func TestArrayListRemoveElement(t *testing.T) {
	myList := arraylist.CreateList()
	myList.Add("a")
	myList.Add("b")
	myList.Add("c")
	myList.Add("d")
	value1 := myList.RemoveElement("b")
	if !value1 {
		t.Fatal("Method remove should return true with valid element")
	}
	value2 := myList.RemoveElement("hello")
	if value2 {
		t.Fatal("Method remove should return false when not found")
	}
}

func TestArrayListContains(t *testing.T) {
	myList := arraylist.CreateList()
	myList.Add("a")
	myList.Add("b")
	myList.Add("c")
	myList.Add("d")
	value1 := myList.Contains("b")
	if !value1 {
		t.Fatal("Method remove should return true with valid element")
	}
	value2 := myList.Contains("hello")
	if value2 {
		t.Fatal("Method remove should return false when not found")
	}
}

func TestArrayListSize(t *testing.T) {
	myList := arraylist.CreateList()
	if myList.Size() != 0 {
		t.Fatal("List initialized to non-zero size")
	}
	myList.Add("a")
	myList.Add("b")
	myList.Add("c")
	myList.Add("d")
	if myList.Size() != 4 {
		t.Fatal("List size should be 4")
	}
}

func TestArrayListIsEmpty(t *testing.T) {
	myList := arraylist.CreateList()
	if !myList.IsEmpty() {
		t.Fatal("List initialized to non-zero size")
	}
	myList.Add("a")
	if myList.IsEmpty() {
		t.Fatal("List should not be empty after adding elements")
	}
}

func TestArrayListReverse(t *testing.T) {
	myList := arraylist.CreateList()
	myList.Add("a")
	myList.Add("b")
	myList.Add("c")
	myList.Add("d")
	myList.Reverse()
	if myList.Print()[0] != "d" {
		t.Fatal("List is not reversed")
	}
	if myList.Print()[myList.Size()-1] != "a" {
		t.Fatal("List is not reversed")
	}
}

func TestArrayListString(t *testing.T) {
	myList := arraylist.CreateList()
	myList.Add("a")
	myList.Add("b")
	myList.Add("c")
	myList.Add("d")
	if myList.String() != "a, b, c, d" {
		t.Fatal("List string is incorrect")
	}
}
