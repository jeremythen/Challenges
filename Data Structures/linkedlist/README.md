# LinkedList

- [LinkedList](#linkedlist)
  - [Functions](#functions)
  - [Methods](#methods)

Implemented with custom `Node` type that holds the value of the node and the next node in the list.\
Includes full test coverage.

## Functions

| Method                    | Description                              | Time | Space |
| ------------------------- | ---------------------------------------- | ---- | ----- |
| `CreateList() LinkedList` | Initialize new linked list with one node | O(1) | O(1)  |

## Methods

| Method                                          | Description                                           | Time                                           | Space                                  |
| ----------------------------------------------- | ----------------------------------------------------- | ---------------------------------------------- | -------------------------------------- |
| `(*LinkedList) Add(interface{})`                | Add element to list                                   | O(1)<br>Add to the end                         | O(1)<br>No extra space needed          |
| `(*LinkedList) Get(int) (interface{}, error)`   | Get element from list, error if index out of bounds   | O(n)<br>Traverse the list until index          | O(1)<br>No extra space needed          |
| `(*LinkedList) RemoveIndex(int) (bool, error)`  | Remove element by index, error if index out of bounds | O(n)<br>Traverse the list until index          | O(1)<br>No extra space needed          |
| `(*LinkedList) RemoveElement(interface{}) bool` | Remove element if found, false if not found           | O(n)<br>All elements are accessed              | O(1)<br>No extra space needed          |
| `(*LinkedList) Contains(interface{}) bool`      | Predicate function for value existence in list        | O(n)<br>All elements are accessed              | O(1)<br>No extra space needed          |
| `(*LinkedList) Size() int`                      | Number of elements in list                            | O(1)<br>List keep tracks of length             | O(1)<br>No extra space needed          |
| `(*LinkedList) Reverse()`                       | Reverse the order of elements in list                 | O(n)<br>All elements are accessed              | O(1)<br>Elements are swapped           |
| `(*LinkedList) ReverseRecursive()`              | Reverse the order of elements in list recursively     | O(n)<br>Linear recursion                       | O(n)<br>Stack created for each element |
| `(*LinkedList) String() string`                 | Print list of elements as string                      | O(n)<br>All elements are accessed              | O(n)<br>String grows with list size    |
| `(*LinkedList) HasCycle() bool`                 | Predicate for cyclic list                             | O(n)<br>All elements are accessed              | O(1)<br>Seen nodes are saved           |
| `(*LinkedList) MiddleNode() interface{}`        | Find the element in the middle                        | O(n)<br>Grows linearly with number of elements | O(1)<br>No extra space needed          |
