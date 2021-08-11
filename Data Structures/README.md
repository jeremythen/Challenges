# Data Structures

- [Data Structures](#data-structures)
  - [Goal](#goal)
  - [Data Structures to implement](#data-structures-to-implement)
  - [Implementation details](#implementation-details)
    - [Array List](#array-list)
    - [Linked List](#linked-list)
      - [Problems to resolve on a linked list:](#problems-to-resolve-on-a-linked-list)
    - [Doubly Link List](#doubly-link-list)
    - [Queue](#queue)
      - [Problems to resolve on a queue:](#problems-to-resolve-on-a-queue)
    - [Stack](#stack)
      - [Methods to implement](#methods-to-implement)
      - [Problems to resolve on a stack:](#problems-to-resolve-on-a-stack)
    - [Double Ended Queue (Deque)](#double-ended-queue-deque)
      - [Methods to implement](#methods-to-implement-1)
    - [HashMap](#hashmap)
      - [Methods to implement](#methods-to-implement-2)
    - [HashSet](#hashset)
      - [Methods to implement](#methods-to-implement-3)
    - [Linked Hash Map](#linked-hash-map)
      - [Methods to implement](#methods-to-implement-4)
    - [Binary Search Tree](#binary-search-tree)
      - [Methods to implement](#methods-to-implement-5)

A challenge designed to help beginners and intermediate programmers to implement and understand Data Structures.

To begin, fork this project, clone it, create a new folder inside "Data Structures" folder and start solving the challenge there. Push a pull request for me to be able to review your solution.

Any questions? Email me: jeremythen16@gmail.com

## Goal

With this challenge, we try to cover a few important points:

1) Understand data structures
2) Implement the most common data structures
3) Solve the most common and interesting problems in the most common data structures
4) Analyze the time complexity of the common operations on the data structures
5) Know when to use them and when to choose one over another.
6) Advantages and disadvantages

## Data Structures to implement

1) Array List
2) Linked List
3) Doubly Link List
4) Queue
   1) ArrayQueue
   2) LinkedQueue
5) Stack
   1) ArrayStack
   2) LinkedStack
6) Deque (Double Ended Queue)
   1) ArrayDeque
   2) LinkedDeque
7) Hash Map (Hash Tables)
8) Hash Set
9) Linked Hash Map
10) Binary Search Tree
11) Binary Heap (Min and Max heap)
12) Priority Queue
13) Graph

Feel free to decide what is the best hierarchy to reuse code, if using OOP.

## Implementation details

### Array List

Implement this data structure with an array (as the name suggests).

Document all the methods and specify the time and space complexity and the reason why this is, on the documentation comment.

Add an `void add` method to add elements to the list. Make sure there is enough free space in the implementation array before adding new elements. The list should grow dynamically if needed. For this, recreation of the implementing array may be needed. When the implementing array is already reaching it's capacity, create a new implementing array with double the capacity and copy the elements to this new array.

Implement a `T get(int index)` method to return the element at the index specified in the parameter. If an invalid index is provided (index < 0 or index >= array.length), throw an error specifying that.

Implement a `boolean remove(int index)` method to remove the element at the index specified in the parameter. If an invalid index is provided (index < 0 or index >= array.length), throw an error specifying that.
When removing an item from the implementing array, shifting to the left of the remaining items on the right may be needed.
If we have an array like this: `[1,2,3,4,5]`, and we are told to remove the element at index 2 (which is the element 3 in the array), then the array should look like this after the removal: `[1,2,4,5, null]`. Basically shifting left the elements to the right of the array. At this point, the last position of this array should be ignore because it doesn't represent a valid element and the size of the list should be 4, not  5.

Implement a `boolean remove(T element)` method to remove the element with the specified value.
Return true if the element was found and removed. False otherwise.
And follow the same rules as above.

Implement a `boolean contains(T element)` method to see if the list has that element. 

Implement a `int size()` method to return the number of elements in the implementation array.

Implement a `boolean isEmpty` method to return if the list is empty or not.

Implement a `String toString()` method to return a representation of the list values in a string.

Implement a `void reverse()` method to reverse the array list.


### Linked List

Implement all the methods defined in the Array List, following the same rules when applicable.
For example, in the `remove` method, implementing it with a linked list we won't need to shift any elements to the left, simply make the previous element point to the next element and make the method to be removed point to null.
In your implementation, use a Node<E> object that will contain the node value and point to the next element.

#### Problems to resolve on a linked list:

1) Reverse the linked list
   1) Iteratively, using a second data structure to hold the data temporarily.
   2) Recursively, without using any intermediate data structure, just recursion.
2) Check if a linked list has cycles.
   1) Using intermediate data structures
   2) Using no other data structure
3) Finding the the element in the middle of the linked list
   1) Without having the size beforehand
   2) Having the size beforehand
   3) Using fast and slow fast pointers

...

### Doubly Link List

Document the methods and specify the time and space complexity and the reason why this is, on the documentation comment.

Implement all the methods defined in the Linked List and these additional methods:

Implement a `void addFirst(T element)` method to add an element to the head of the list.

Implement a `T getFirst()` method to return, not remove, the head of the list.
 
Implement a `T getLast()` method to , not remove, the tail of the list.


### Queue

Document the methods and specify the time and space complexity and the reason why this is, on the documentation comment.

1) Implement an ArrayQueue (queue implemented with an array as the backing data structure).
   
2) Implement a LinkedQueue (queue implemented with a Doubly Link List as the backing data structure)

Implement a `void push(T element)` method to add an element to the tail of the queue.

Implement a `T poll()` method to remove an element from the head of the queue.

Implement a `T peek()` method to retrieve, but not remove, the head of the queue.

Implement a `int size()` method to see the size of the queue.

Implement a `boolean isEmpty()` method to check if the queue is empty.

#### Problems to resolve on a queue:

1) Reverse the queue
   1) Reverse it using a backing data structure
   2) Reverse it using recursion

### Stack

Document the methods and specify the time and space complexity and the reason why this is, on the documentation comment.

A Stack is a LIFO (Last In First Out) data structure.

1) Implement an ArrayStack (stack implemented with an array as the backing data structure).
   
2) Implement a LinkedStack (stack implemented with a Doubly Link List as the backing data structure)
#### Methods to implement

Implement a `void push(T element)` method to add an element to the top of the stack.

Implement a `T pop()` method to remove an element from the top of the stack.

Implement a `T peek()` method to retrieve, but not remove, the element on the top of the stack.

Implement a `int size()` method to see the size of the stack.

Implement a `boolean isEmpty()` method to check if the stack is empty.

#### Problems to resolve on a stack:

These problems needs to be solved on the two implementation of the stack, ArrayStack and LinkedStack.

1) Reverse the stack
   1) Reverse it using a backing data structure
   2) Reverse it using recursion

### Double Ended Queue (Deque)

Document the methods and specify the time and space complexity and the reason why this is, on the documentation comment.

1) Implement an ArrayDeque (deque implemented with an array as the backing data structure).
   
2) Implement a LinkedDeque (deque implemented with a Doubly Link List as the backing data structure)
#### Methods to implement

Implement all the methods defined in the Queue and these additional methods:

Implement a `void addFirst(T element)` method to add an element to the head of the queue.
 
Implement a `T peekLast()` method to return, not remove, the tail of the queue.

Implement a `T removeLast()` method to return and remove the tail of the queue.

### HashMap

Document the methods and specify the time and space complexity and the reason why this is, on the documentation comment.

This data structure is implemented with an array as the backing data structure.
The keys (which may be of any type) need to be converted into a valid positive integer to use as an index for the array. This has to be done with a `int hashingFunction(K key)`, a function that takes an input and converts it into a valid array index.

This map will need to grow dynamically to any required number of elements.

This data structure needs to be able to manage collisions with a **Separate Chaining** approach.

#### Methods to implement

Implement a `V put(K key, V value)` method to add a record to the hash map.

Implement a `V get(K key)` method to get the value of the element under that key.

Implement a `V remove(K key)` method to remove the element under that key, returning it too.

Implement a `boolean containsKey(K key)` method to check if the map contains that key.

Implement a `boolean containsValue(V v)` method to check if the map contains that value.

Implement a `int size()` method to check the size of the map.

Implement a `boolean isEmpty` method to check if the map is empty.

Implement a `V[] getValues()` method to return an array of all the values in the hash map.

Implement a `V[] getEntries()` method to return an array of all the entries (key, value) in the hash map.

### HashSet

Document the methods and specify the time and space complexity and the reason why this is, on the documentation comment.

A set is a data structure that doesn't allow any duplicates.

This data structure is implemented with an array as the backing data structure.
The elements need to be converted into a valid positive integer to use as an index for the array. This has to be done with a `int hashingFunction(K key)`, a function that takes an input and converts it into a valid array index.

This set will need to grow dynamically to any required number of elements.

This data structure needs to be able to manage collisions with a **Separate Chaining** approach.

#### Methods to implement

Implement a `void add(T element)` method to add an element to the set.

Implement a `boolean remove(T element)` method to remove an element from the set.

Implement a `boolean contains(T element)` to check if the set already contains the specified element.

Implement a `int size()` method to check the size of the set.

Implement a `boolean isEmpty` method to check if the set is empty.


### Linked Hash Map

A Linked Hash Map is a hash map implemented with an array of linked nodes, in a way that each know inserted points to another node, just a regular singly linked list. This is useful to keep the order in each the elements were inserted in the hash map, while keeping the same time complexity and efficiency on the put/get/remove operations of the hash map.

#### Methods to implement

Implement all the methods defined in the Hash Map.

When implementing `V[] getValues()` and `V[] getEntries()` methods, the values and entries should be returned in the order in which they were added to the map.

### Binary Search Tree

A Binary Search Tree is a data structure composed of nodes that can have up to 2 childing nodes. The first and main node is called the root node. Each node's child greater than the parent node needs to be to the right of the parent, and each node's child node less than the parent node needs to be to the left of the node.

In this way, the search is really efficient, because when searching for a specific key or value, you will either go to the left or to the right and after each step, you will skip have of the remaining tree.

To implement this data structure, use a sort of linked node, where the node is an objet that contains 3 properties:

```
Node
 Node leftChild
 Node rightChild
 V value
```

where the leftChild will point to the node (or subtree) to the left (nodes whose values are less than their parent value).
Where the rightChild will point to the node (or subtree) to the right (nodes whose values are greater than their parent value).
And value, which will contain the value to keep and compare.

#### Methods to implement

Implement a `void add(T element)` method to add an element to the tree.  If the tree is empty, this will be the root node.

Implement a `boolean remove(T element)` method to remove the specify value from the tree. Note that you will need to find a successor or predecessor to replace that node to keep the structure of the tree internally.

Implement a `T getSuccessor(T element)` method to return the successor of this element. The successor is the element that would naturally succeed the specified element. If we have the elements 4,8,12,17,22,16, the successor of 12 if 17, because if we have all the numbers in order, 17 is the number that exactly goes after 12, based on their natural order.
The successor of a node in a binary search tree is the left-most child of the first right child of the node.

Implement a `T getPredecessor(T element)` method to return the predecessor of this element. The predecessor is the element that would naturally precede the specified element. If we have the elements 4,8,12,17,22,16, the predecessor of 12 if 8, because if we have all the numbers in order, 8 is the number that exactly goes before 12, based on their natural order.
The predecessor of a node in a binary search tree is the right-most child of the first left child of the node.

Implement a `String toStringInOder()` method to return a string representation of the binary search tree being traversed in an In Order way.

Implement a `String toStringPostOrder()` method to return a string representation of the binary search tree being traversed in an Post Order way.

Implement a `String toStringPreOrder()` method to return a string representation of the binary search tree being traversed in an Pre Order way.

Implement a `T[] toArrayInOrder()` method to return an array with all the elements from the tree, traversed in a In Order way.

