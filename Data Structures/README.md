# Data Structures

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

## Data Structures to implement:

1) Array List
2) Linked List
3) Doubly Link List
4) Queue
5) Stack
6) Deque (Double Ended Queue)
7) Hash Map (Hash Tables)
8) Hash Set
9)  Linked Hash Map
10) Binary Search Tree
11) Binary Heap (Min and Max heap)
12) Graph

Feel free to decide what is the best hierarchy to reuse code, if using OOP.

## Implementation details

### Array List

Implement this data structure with an array (as the name suggests)

Add an `void add` method to add elements to the list. Make sure there is enough free space in the implementation array before adding new elements. The list should grow dynamically if needed. For this, recreation of the implementing array may be needed. When the implementing array is already reaching it's capacity, create a new implementing array with double the capacity and copy the elements to this new array.
Document the method and specify the time and space complexity and the reason why this is, on the documentation comment.

Add a `T get(int index)` method to return the element at the index specified in the parameter. If an invalid index is provided (index < 0 or index >= array.length), throw an error specifying that.
Document the method and specify the time and space complexity and the reason why this is, on the documentation comment.

Add a `boolean remove(int index)` method to remove the element at the index specified in the parameter. If an invalid index is provided (index < 0 or index >= array.length), throw an error specifying that.
When removing an item from the implementing array, shifting to the left of the remaining items on the right may be needed.
If we have an array like this: `[1,2,3,4,5]`, and we are told to remove the element at index 2 (which is the element 3 in the array), then the array should look like this after the removal: `[1,2,4,5, null]`. Basically shifting left the elements to the right of the array. At this point, the last position of this array should be ignore because it doesn't represent a valid element and the size of the list should be 4, not  5.
Document the method and specify the time and space complexity and the reason why this is, on the documentation comment.

Add a `boolean remove(T element)` method to remove the element with the specified value.
Return true if the element was found and removed. False otherwise.
And follow the same rules as above.

Add a `boolean contains(T element)` method to see if the list has that element. 

Add a `int size()` method to return the number of elements in the implementation array.

Add a `boolean isEmpty` method to return if the list is empty or not.

Add a `String toString()` method to return a representation of the list values in a string.

Add a `void reverse()` method to reverse the array list.


### Linked List

Implement all the methods defined in the Array List, following the same rules when applicable.
For example, in the `remove` method, implementing it with a linked list we won't need to shift any elements to the left, simply make the previous element point to the next element and make the method to be removed point to null.
In your implementation, use a Node<E> object that will contain the node value and point to the next element.

#### Common problems to resolve on a linked list:

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



