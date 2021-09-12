package linkedlist

import (
	"fmt"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	value interface{}
	next  *Node
}

type LinkedListError struct {
	message string
}

func (e *LinkedListError) Error() string {
	return e.message
}

func CreateList(value interface{}) *LinkedList {
	head := CreateNode(value)
	return &LinkedList{
		head:   head,
		tail:   head,
		length: 1,
	}
}

func CreateNode(value interface{}) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

func (ll *LinkedList) Add(value interface{}) {
	if value == nil {
		return
	}
	ll.tail.next = CreateNode(value)
	ll.tail = ll.tail.next
	ll.length++
}

func (ll *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= ll.length {
		return nil, &LinkedListError{"Invalid index"}
	}
	currentNode := ll.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.value, nil
}

func (ll *LinkedList) RemoveIndex(index int) (bool, error) {
	if index < 1 || index >= ll.length {
		return false, &LinkedListError{"Invalid index"}
	}
	currentNode := ll.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
	ll.length--
	return true, nil
}

func (ll *LinkedList) RemoveElement(value interface{}) bool {
	if ll.head.value == value && ll.length == 1 {
		return false
	}
	currentNode := ll.head
	prevNode := &Node{}
	for currentNode != nil {
		if currentNode.value == value {
			if currentNode == ll.head {
				ll.head = currentNode.next
			}
			prevNode.next = currentNode.next
			ll.length--
			return true
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}
	return false
}

func (ll *LinkedList) Contains(value interface{}) bool {
	currentNode := ll.head
	for currentNode != nil {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (ll *LinkedList) Size() int {
	return ll.length
}

func (ll *LinkedList) Reverse() {
	if ll.length == 1 {
		return
	}
	prev := &Node{}
	current := ll.head
	for current != nil {
		tempPrev := prev
		prev = current
		current = current.next
		prev.next = tempPrev
	}
	ll.tail = ll.head
	ll.tail.next = nil
	ll.head = prev
}

func (ll *LinkedList) ReverseRecurse() {
	var reverse func(*Node) *Node
	reverse = func(node *Node) *Node {
		if node.next == nil {
			ll.head = node
			return node
		}
		newNode := reverse(node.next)
		newNode.next = node
		return node
	}
	ll.tail = reverse(ll.head)
	ll.tail.next = nil
}

func (ll *LinkedList) String() string {
	currentNode := ll.head
	print := "List"
	for currentNode != nil {
		print += fmt.Sprintf(" -> %v", currentNode.value)
		currentNode = currentNode.next
	}
	return print
}

func (ll *LinkedList) HasCycle() bool {
	if ll.length == 1 {
		return false
	}
	seenNodes := make(map[*Node]int)
	currentNode := ll.head
	for currentNode != nil {
		if seenNodes[currentNode] > 0 {
			return true
		}
		seenNodes[currentNode]++
		currentNode = currentNode.next
	}
	return false
}

func (ll *LinkedList) MiddleNode() interface{} {
	if ll.length == 1 {
		return ll.head.value
	}
	slowPointer := ll.head
	fastPointer := slowPointer.next
	for fastPointer != nil && fastPointer.next != nil {
		slowPointer = slowPointer.next
		fastPointer = fastPointer.next.next
	}
	return slowPointer.value
}
