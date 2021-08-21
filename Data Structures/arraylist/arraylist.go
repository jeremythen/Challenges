package arraylist

import "fmt"

type ArrayList struct {
	elements []interface{}
}

type ArrayListError struct {
	message string
}

func (e *ArrayListError) Error() string {
	return e.message
}

func CreateList() ArrayList {
	return ArrayList{make([]interface{}, 2)}
}

func (al *ArrayList) Add(element interface{}) {
	if element == nil {
		return
	}
	currentSize := al.Size()
	if currentSize == len(al.elements) {
		newElements := make([]interface{}, currentSize*2)
		for index, value := range al.elements {
			newElements[index] = value
		}
		al.elements = newElements
	}
	al.elements[currentSize] = element
}

func (al *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= al.Size() {
		return nil, &ArrayListError{"Invalid index"}
	}
	return al.elements[index], nil
}

func (al *ArrayList) RemoveIndex(index int) (bool, error) {
	if index < 0 || index >= al.Size() {
		return false, &ArrayListError{"Invalid index"}
	}
	for i, value := range al.elements {
		if i < index {
			al.elements[i] = value
		}
		if i > index {
			al.elements[i-1] = value
		}
	}
	al.elements[al.Size()-1] = nil
	return true, nil
}

func (al *ArrayList) RemoveElement(element interface{}) bool {
	for index, value := range al.elements {
		if value == element {
			al.RemoveIndex(index)
			return true
		}
	}
	return false
}

func (al *ArrayList) Contains(element interface{}) bool {
	for _, value := range al.elements {
		if value == element {
			return true
		}
	}
	return false
}

func (al *ArrayList) Size() int {
	size := 0
	for _, value := range al.elements {
		if value != nil {
			size++
		}
	}
	return size
}

func (al *ArrayList) IsEmpty() bool {
	return al.elements[0] == nil
}

func (al *ArrayList) Reverse() {
	size := al.Size()
	newElements := make([]interface{}, size)
	for i := 0; i < size; i++ {
		newElements[i] = al.elements[size-i-1]
	}
	al.elements = newElements
}

func (al ArrayList) String() string {
	stringList := ""
	for index, value := range al.elements {
		if value == nil {
			continue
		}
		if index > 0 {
			stringList += ", "
		}
		stringList += fmt.Sprint(value)
	}
	return stringList
}

func (al *ArrayList) Print() []interface{} {
	return al.elements
}
