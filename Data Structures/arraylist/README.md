# Arraylist

- [Arraylist](#arraylist)
  - [Functions](#functions)
  - [Methods](#methods)

Implemented with slices and avoiding `copy`/`append` or any slice functionality to maintain manual control of underlying array size.\
Included tests can be run with `go test ./...`

## Functions

| Method                   | Description                                               | Time | Space |
| ------------------------ | --------------------------------------------------------- | ---- | ----- |
| `CreateList() ArrayList` | Initialize new array list with length of two `nil` values | O(1) | O(1)  |

## Methods

| Method                                         | Description                                           | Time                                          | Space                                         |
| ---------------------------------------------- | ----------------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `(*ArrayList) Add(interface{})`                | Add element to list                                   | O(n) to copy array if full<br> O(1) amortized | O(n) to copy array if full<br> O(1) amortized |
| `(*ArrayList) Get(int) (interface{}, error)`   | Get element from list, error if index out of bounds   | O(1)<br>Access array index directly           | O(1)<br>No extra space needed                 |
| `(*ArrayList) RemoveIndex(int) (bool, error)`  | Remove element by index, error if index out of bounds | O(n)<br>All elements are accessed             | O(1)<br>Removed in place                      |
| `(*ArrayList) RemoveElement(interface{}) bool` | Remove element if found, false if not found           | O(n)<br>All elements are accessed             | O(1)<br>Removed in place                      |
| `(*ArrayList) Contains(interface{}) bool`      | Predicate function for value existence in list        | O(n)<br>All elements are accessed             | O(1)<br>No extra space needed                 |
| `(*ArrayList) Size() int`                      | Number of elements in list                            | O(n)<br>All elements are accessed             | O(1)<br>No extra space needed                 |
| `(*ArrayList) IsEmpty() bool`                  | Predicate function for list size                      | O(1)<br>Access first element only             | O(1)<br>No extra space needed                 |
| `(*ArrayList) Reverse()`                       | Reverse the order of elements in list                 | O(n)<br>All elements are accessed             | O(n)<br>New list is created                   |
| `(ArrayList) String() string`                  | Print list of elements as string                      | O(n)<br>All elements are accessed             | O(n)<br>String grows with list size           |
| `(*ArrayList) Print() []interface{}`           | Get the underlying array of elements                  | O(1)<br>Return array address                  | O(1)<br>No extra space needed                 |
