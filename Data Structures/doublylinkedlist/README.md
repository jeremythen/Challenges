# DoublyDoublyLinkedList

- [DoublyDoublyLinkedList](#doublydoublylinkedlist)
  - [Functions](#functions)
  - [Methods](#methods)

Implemented with custom `Node` type that holds the value of the node and the previous/next nodes in the list.\
Includes full test coverage.

## Functions

| Method                          | Description                                     | Time | Space |
| ------------------------------- | ----------------------------------------------- | ---- | ----- |
| `CreateList() DoublyLinkedList` | Initialize new doubly linked list with one node | O(1) | O(1)  |

## Methods

| Method                                                   | Description                                           | Time                                           | Space                                  |
| -------------------------------------------------------- | ----------------------------------------------------- | ---------------------------------------------- | -------------------------------------- |
| `(*DoublyLinkedList) Add(interface{})`                   | Add element to list                                   | O(1)<br>Add to the end                         | O(1)<br>No extra space needed          |
| `(*DoublyLinkedList) GetIndex(int) (interface{}, error)` | Get element from list, error if index out of bounds   | O(n)<br>Traverse the list until index          | O(1)<br>No extra space needed          |
| `(*DoublyLinkedList) RemoveIndex(int) (bool, error)`     | Remove element by index, error if index out of bounds | O(n)<br>Traverse the list until index          | O(1)<br>No extra space needed          |
| `(*DoublyLinkedList) RemoveElement(interface{}) bool`    | Remove element if found, false if not found           | O(n)<br>All elements are accessed              | O(1)<br>No extra space needed          |
| `(*DoublyLinkedList) Contains(interface{}) bool`         | Predicate function for value existence in list        | O(n)<br>All elements are accessed              | O(1)<br>No extra space needed          |
| `(*DoublyLinkedList) Size() int`                         | Number of elements in list                            | O(1)<br>List keep tracks of length             | O(1)<br>No extra space needed          |
| `(*DoublyLinkedList) Reverse()`                          | Reverse the order of elements in list                 | O(n)<br>All elements are accessed              | O(1)<br>Elements are swapped           |
| `(*DoublyLinkedList) ReverseRecursive()`                 | Reverse the order of elements in list recursively     | O(n)<br>Linear recursion                       | O(n)<br>Stack created for each element |
| `(*DoublyLinkedList) String() string`                    | Print list of elements as string                      | O(n)<br>All elements are accessed              | O(n)<br>String grows with list size    |
| `(*DoublyLinkedList) HasCycle() bool`                    | Predicate for cyclic list                             | O(n)<br>All elements are accessed              | O(n)<br>Seen nodes are saved           |
| `(*DoublyLinkedList) MiddleNode() interface{}`           | Find the element in the middle                        | O(n)<br>Grows linearly with number of elements | O(1)<br>No extra space needed          |
| `(*DoublyLinkedList) AddFirst(interface{})`              | Add element to beginning of list                      | O(1)<br>Add element and update head            | O(1)<br>No extra space needed          |
| `(*DoublyLinkedList) GetFirst() interface{}`             | Get element at beginning of list                      | O(1)<br>Get head                               | O(1)<br>No extra space needed          |
| `(*DoublyLinkedList) GetLast() interface{}`              | Get element at end of list                            | O(1)<br>Get tail                               | O(1)<br>No extra space needed          |
