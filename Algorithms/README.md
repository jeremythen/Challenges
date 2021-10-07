# algorithms
Repository to hold algorithms descriptions and solutions in multiple programing languages, as practice.

- [algorithms](#algorithms)
  - [List of algorithms to implement:](#list-of-algorithms-to-implement)
    - [Max Chars](#max-chars)
    - [Fibonacci with 2 variables](#fibonacci-with-2-variables)
    - [Matrix Spiral](#matrix-spiral)
    - [Revert a number](#revert-a-number)
    - [Array Chunk](#array-chunk)
    - [Capitalize a string](#capitalize-a-string)
    - [Steps](#steps)
    - [Pyramid](#pyramid)
    - [Weave](#weave)
    - [Queue From Stacks](#queue-from-stacks)
    - [Permutation](#permutation)
    - [Row Sum Of Triangle Of Odd Numbers](#row-sum-of-triangle-of-odd-numbers)
    - [Sums](#sums)
      - [Can Sum](#can-sum)
      - [How Sum](#how-sum)
      - [Best Sum](#best-sum)
    - [Construct](#construct)
      - [Can Construct](#can-construct)
      - [How Construct](#how-construct)
      - [All Construct](#all-construct)
    - [First Non Repeating Character](#first-non-repeating-character)
    - [Palindromic Substring](#palindromic-substring)
    - [Digit At](#digit-at)
    - [Remove Duplicates](#remove-duplicates)
    - [Remove target occurrences](#remove-target-occurrences)
    - [Implement indexOf](#implement-indexof)
    - [Length of last word](#length-of-last-word)
    - [Sum in Array](#sum-in-array)
    - [Max 1 Bits Sequence](#max-1-bits-sequence)
    - [Move Zeros](#move-zeros)
    - [What Letter Was Added](#what-letter-was-added)
    - [Concatenate integers](#concatenate-integers)
    - [Sorting](#sorting)

## List of algorithms to implement:

### Max Chars
Receive a strings and return the most repeated character.
Given "abcdefacasdfaaaa" should return "a", because the "a" is the most repeated character in that string.

### Fibonacci with 2 variables
Return the value of the fibonacci number at the fibonacci position of n, where n is a positive number.
Given 6, return 8. Given 7, return 13, etc.
Only use a loop and 2 integers to hold the fibonacci values to solve the problem.

### Matrix Spiral

Given a number N, create and return a matrix N * N and fill the values in a matrix fashion. Check example image below.

Examples:

```js

matrix(2);

// [[1, 2]
// [4, 3]]

matrix(3);

// [[1, 2, 3],
// [8, 9, 4],
// [7, 6, 5]]

```

![Matrix Spirals](spiral.jpg)

### Revert a number
Given an integer, return an integer with the digits reversed. For example:
123 -> 321
500 -> 5
503 -> 305

Solve it any way, but create a final version where only arithmetics is used.

### Array Chunk

Given an array and chunk size, divide the array into many subarrays where each subarray is of the provided size.

Examples:

```js
chunk([1,2,3,4], 2) // -> [[1,2], [3,4]]
chunk([1,2,3,4,5], 2) // -> [[1,2], [3,4], [5]]

```

### Capitalize a string

Give a string, capitalize the first letter of all the words in it.

Examples:
```js
capitalize("this is an example") // -> "This Is An Example"
```

### Steps

Write a function that accepts a positive number n. The function should console.log a step shape with N levels using the # character. Make sure the step has spaces on the right hand side.

Examples:
```js
step(2);

// '# '
// '##'

step(3);

// '#  '
// '## '
// '###'

```

### Pyramid

Write a function that accepts a positive number N.

The function should print a pyramid shape with N levels using the # character.
Make sure the pyramid has spaces on both the left and right hand sides.

Given:

```js
pyramid(2);

' # '
'###'
```
```js
pyramid(3);

'  #  '
' ### '
'#####'

```

```js
pyramid(4);

'   #   '
'  ###  '
' ##### '
'#######'

```

### Weave

Weave receives two queues as arguments and combines the contents of each into a new, third queue.
The third queue should contain the alternating content of the two queues.
The function should handle queues of different lengths, without inserting 'undefined' or 'null' into the new one. If it doesn't have more values, it should try to add any.

Example:

 
```js

// Given:
queue1 = [1,2,3,4]
queue2 = [5,6,7,8]

// return
queue3 = [1,5,2,6,3,7,4,8]

// Given
queue1 = [1,2,3,4,5,6]
queue2 = [7,8,9]

// return
queue3 = [1,7,2,8,3,9,4,5,6]

```

### Queue From Stacks

Implement a Queue data structure using two stacks.
Create another abstraction that holds the tow stacks and provide queue methods to access this abstraction. Only use the `pop` and `push` methods of the stacks to simulate the queue.

### Permutation

Given a string S, return a list with all the permutation of that string.

Given S = "abc", return [abc, acb, bac, bca, cab, cba]

### Row Sum Of Triangle Of Odd Numbers

Given a triangle of consecutive odd numbers.

If given 1, return 1.
If given 2, return the sum of the numbers in the second row, of 3 + 5 => 8.
If given 3, return the sum of the numbers in the third row, of 7 + 9 + 11 => 27.

![triangle of odd numbers](triangle_of_odd_numbers.jpg)

### Sums

#### Can Sum
Given a number `N`, and an array of numbers `numbers`, verify if you can recreate the target number `N` summing numbers from the `numbers` array.
Return true if it's possible to create the target number `N` summing numbers from the `numbers` array. Return false otherwise.


Given

n = 7
numbers = [2, 3, 4, 7]

return true because you can create the n 7 from the numbers array.

Given

n = 5
numbers [2,4,6]

Return false because 5 cannot be recreated from any combination of numbers from numbers.

Use any combination of numbers and use any numbers as many times as needed.

Write a function solution(n, numbers)

#### How Sum
Like Can Sum, but return an array showing which numbers from the `numbers` when summed will produce the number `N`.
#### Best Sum
Like Can Sum and How Sum, but should return an array with the minimum numbers from the `numbers` array that summed will produce the number `N`.

Check the dedicated folder for this problem, for more details.
### Construct


#### Can Construct

Create a function `function solution(str, words)` that checks whether a string can be recreated from a bank or words (array of strings), how it can be created and how many times it can be created from that word bank. Use as many words from the word bank as you want.

Given

s = "hello"

And an array:

words = ["he", "l", "o", "llo"]

Return true, because you can use values from the words array to recreate the 'hello' s.

Given

s = "tomorrow"
words = ["to", "m", "rro", "w", "r", "t"]

Return false, because the s "tomorrow" cannot be recreated from the strings given in words array.

#### How Construct
The same as Can Construct, bue returns an array with the strings used from the words array, to create the string s.
#### All Construct
Same as How Construct, but returns an array of arrays, showing each combination of the ways the string s could be formed from the strings in the words array.

Check the dedicated folder for this problem, for more details.

### First Non Repeating Character

Find the first character that doesn't repeat in a string.

Examples:

Given "abcde", return "a" because "a" is the first character that doesn't repeat.
Given "abcda", return "b" because "b" is the first character that doesn't repeat.
Given "abcdeabcdf", return "e", because although "f" doesn't repeat in the string, "e" is the first that doesn't repeat.

Assume a non empty string is given.

### Palindromic Substring

Find the largest sub palindrome in a string.
If there is no palindrome or the palindrome is less than 3 chars, return "none".

Given hellosannasmith, return sannas.
Given abcdefgg, return none

### Digit At

Implement function that receives an int number `num` and an int number `position`, where the `position` is represent the ith digit of `num` to return, counting from the back of the number.

Given `num` = 7895, `position` = 2, return 9, because 9 is the 2nd digit of `num` from the back.

### Remove Duplicates

Given a sorted array of integers in ascending order, modify the given array removing the duplicates. I.e, each element should appear only once. And print the resulting array and k, where k is the number of unique numbers in the array.

Examples:

Given numbers = [1,1,2]
Return: numbers = [1,2,] and k = 2, because there are 2 unique numbers.

Given numbers =  [0,0,1,1,1,2,2,3,3,4]
Return: numbers =  [0,1,2,3,4,,,,,] and k = 5, because there are 5 unique numbers.

### Remove target occurrences

Given an array only containing integers and a value called target which is also an integer, Modify the input array to delete all the elements that are equal to target, shifting all the other elements to the left and print the count of the remaining elements on that array that were not removed.

Given: numbers = [3,2,2,3], val = 3
Print: [2,2,,], 2, because the target is 3, and all it's occurrences were removed from the original array, shifting the other elements to the left. And 2 because that was the amount of elements remaining.

Given: numbers = [0,1,2,2,3,0,4,2], val = 2
Print: numbers = [0,1,4,0,3,,,_], 5.

### Implement indexOf

Given a string and a substring, return the index first occurrence of the substring in the string. Return -1 if no occurrence was found.

Given: str = "Hello world", substr = "ll"
Return: 2, because the first occurrence of "ll" is in the index 2 of the string.

### Length of last word

Given a string containing words separated by one or more spaces, return the length of the last word.

Given: "   hello      world  "
Return: 5, because the last word is "world" and it's length is 5.

### Sum in Array

Given a number represented in an array where each index in the array holds a digit for that number, increment the last digit in the array by 1. The whole array values should change if needed.

Given: digits = [1,2,3]
Return: digits = [1,2,4]

Given: digits = [9,9,9]
Return: digits = [1,0,0,0]

Given: digits = [1,2,3,4,5,6]
Return: digits = [1,2,3,4,5,7]

### Max 1 Bits Sequence

Given a binary number, write a function findMax1BitsSequence(n) that receives a number (or binary) and returns the longest subsequence of 1s.


```javascript
console.log(findMax1BitsSequence(0b1111111110010111011) === 9); // 9 because the longest subsequence of 1s is 9 1s one next to the other.
console.log(findMax1BitsSequence(0b10011110010111011) === 4);
console.log(findMax1BitsSequence(0b1111001) === 4);
console.log(findMax1BitsSequence(0b101) === 1);
```

### Move Zeros

Given an integer array numbers, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

 
Example 1:

Input: numbers = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: numbers = [0]
Output: [0]

Input: numbers = [1,0,4,0,0,10,20,100]
Output: [1,4,10,20,100,0,0,0]

### What Letter Was Added

You are given two strings s and t.
String t is generated by random shuffling string s and then add one more letter at a random position.
write a function that returns the letter that was added to t.


Example 1:

Input: s = "abcd", t = "abcde"
Output: "e"
Explanation: 'e' is the letter that was added.
Example 2:

Input: s = "", t = "y"
Output: "y"
Example 3:

Input: s = "a", t = "aa"
Output: "a"
Example 4:

Input: s = "ae", t = "aea"
Output: "a"

### Concatenate integers

Given 2 integers, write a function that concatenates them into one.

Given: n1 = 234, n2 = 567
Return: 234567

Given: n1 = 6234, n2 = 723
Return: 6234723

We are expecting an efficient solution using arithmetics.

### Sorting

1) Bubble Sort
2) Selection Sort
3) Insertion Sort
4) Merge Sort
5) Quick Sort
6) Heap Sort
7) Radix Sort
8) Bucket Sort
9) Counting Sort
10) Pigeonhole Sort