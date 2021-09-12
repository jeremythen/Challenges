package doublylinkedlist

import (
	"fmt"
)

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	value    interface{}
	next     *Node
	previous *Node
}

type DoublyLinkedListError struct {
	message string
}

func (e *DoublyLinkedListError) Error() string {
	return e.message
}

func CreateList(value interface{}) *DoublyLinkedList {
	newNode := CreateNode(value)
	return &DoublyLinkedList{
		head:   newNode,
		tail:   newNode,
		length: 1,
	}
}

func CreateNode(value interface{}) *Node {
	return &Node{
		value:    value,
		next:     nil,
		previous: nil,
	}
}

func (dll *DoublyLinkedList) Add(value interface{}) {
	if value == nil {
		return
	}
	newNode := CreateNode(value)
	newNode.previous = dll.tail
	dll.tail.next, dll.tail = newNode, newNode
	dll.length++
}

func (dll *DoublyLinkedList) GetIndex(index int) (interface{}, error) {
	if index < 0 || index >= dll.length {
		return nil, &DoublyLinkedListError{"Invalid index"}
	}
	currentNode := dll.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.value, nil
}

func (dll *DoublyLinkedList) RemoveIndex(index int) (bool, error) {
	if index < 1 || index >= dll.length {
		return false, &DoublyLinkedListError{"Invalid index"}
	}
	currentNode := dll.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.previous.next = currentNode.next
	currentNode.next.previous = currentNode.previous
	dll.length--
	return true, nil
}

func (dll *DoublyLinkedList) RemoveElement(value interface{}) (result bool) {
	if dll.head.value == value && dll.length == 1 {
		return false
	}
	currentNode := dll.head
	for currentNode != nil {
		if currentNode.value == value {
			dll.length--
			result = true
			if currentNode == dll.head {
				dll.head = currentNode.next
				dll.head.previous = nil
				return
			}
			if currentNode == dll.tail {
				dll.tail = currentNode.previous
				dll.tail.next = nil
				return
			}
			currentNode.previous.next = currentNode.next
			currentNode.next.previous = currentNode.previous
			return
		}
		currentNode = currentNode.next
	}
	return
}

func (dll *DoublyLinkedList) Contains(value interface{}) bool {
	currentNode := dll.head
	for currentNode != nil {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (dll *DoublyLinkedList) Size() int {
	return dll.length
}

func (dll *DoublyLinkedList) Reverse() {
	if dll.length == 1 {
		return
	}
	current := dll.head
	for current != nil {
		current, current.previous, current.next = current.next, current.next, current.previous
	}
	dll.tail, dll.head = dll.head, dll.tail
}

func (dll *DoublyLinkedList) ReverseRecurse() {
	var reverse func(*Node) *Node
	reverse = func(node *Node) *Node {
		if node.next == nil {
			node.next, dll.head = nil, node
			return node
		}
		newNode := reverse(node.next)
		newNode.next, node.previous = node, newNode
		return node
	}
	dll.tail = reverse(dll.head)
	dll.tail.next = nil
}

func (dll *DoublyLinkedList) String() string {
	currentNode := dll.head
	print := "List"
	for currentNode != nil {
		print += fmt.Sprintf(" <-> %v", currentNode.value)
		currentNode = currentNode.next
	}
	return print
}

func (dll *DoublyLinkedList) HasCycle() bool {
	if dll.length == 1 {
		return false
	}
	seenNodes := make(map[*Node]int)
	currentNode := dll.head
	for currentNode != nil {
		if seenNodes[currentNode] > 0 {
			return true
		}
		seenNodes[currentNode]++
		currentNode = currentNode.next
	}
	return false
}

func (dll *DoublyLinkedList) MiddleNode() interface{} {
	if dll.length == 1 {
		return dll.head.value
	}
	slowPointer := dll.head
	fastPointer := slowPointer.next
	for fastPointer != nil {
		fastPointer = fastPointer.next
		if fastPointer != nil {
			fastPointer = fastPointer.next
			slowPointer = slowPointer.next
		}
	}
	return slowPointer.value
}

func (dll *DoublyLinkedList) AddFirst(value interface{}) {
	if value == nil {
		return
	}
	newNode := CreateNode(value)
	dll.head.previous, dll.head, newNode.next = newNode, newNode, dll.head
}

func (dll *DoublyLinkedList) GetFirst() interface{} {
	return dll.head.value
}

func (dll *DoublyLinkedList) GetLast() interface{} {
	return dll.tail.value
}
