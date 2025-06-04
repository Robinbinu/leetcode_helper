# LeetCode Helper Usage Examples

This document provides practical examples of using the LeetCode Helper application with different AI providers.

## Example 1: Two Sum Problem with OpenAI

### Input

**Problem Text:**
```
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Settings:**
- Language: JavaScript
- User Level: Intermediate
- Provider: OpenAI

### Output

**Explanation:**
The Two Sum problem asks us to find two numbers in an array that add up to a target value. A naive approach would be to use nested loops to check every pair of numbers, but this would have O(n²) time complexity. Instead, we can use a hash map to achieve O(n) time complexity.

The approach is:
1. Create a hash map to store numbers we've seen and their indices
2. Iterate through the array
3. For each number, calculate its complement (target - current number)
4. Check if the complement exists in our hash map
5. If it does, we've found our pair
6. If not, add the current number and its index to the hash map

**Code:**
```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
function twoSum(nums, target) {
  const map = new Map();
  
  for (let i = 0; i < nums.length; i++) {
    const complement = target - nums[i];
    
    if (map.has(complement)) {
      return [map.get(complement), i];
    }
    
    map.set(nums[i], i);
  }
  
  return []; // No solution found
}
```

**Hints:**
- Think about using a hash map to store values you've already seen
- For each number, check if its complement (target - num) exists in the hash map
- Remember to store the index along with the value in the hash map

**Time Complexity:** O(n)
**Space Complexity:** O(n)

## Example 2: Valid Palindrome with Google Gemini

### Input

**Problem Text:**
```
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
```

**Settings:**
- Language: Python
- User Level: Beginner
- Provider: Gemini

### Output

**Explanation:**
This problem asks us to determine if a string is a palindrome after converting to lowercase and removing non-alphanumeric characters. A palindrome reads the same forwards and backwards.

The steps to solve this are:
1. Convert the string to lowercase
2. Remove all non-alphanumeric characters
3. Check if the resulting string is equal to its reverse

For a beginner approach, we'll use Python's built-in functions to make this straightforward.

**Code:**
```python
def isPalindrome(s: str) -> bool:
    # Convert to lowercase and keep only alphanumeric characters
    cleaned_string = ''.join(char.lower() for char in s if char.isalnum())
    
    # Check if the string equals its reverse
    return cleaned_string == cleaned_string[::-1]
```

**Hints:**
- Remember to convert the string to lowercase first
- Use `isalnum()` to check if a character is alphanumeric
- In Python, you can reverse a string using slicing: `s[::-1]`
- Empty strings are considered palindromes

**Time Complexity:** O(n)
**Space Complexity:** O(n)

## Example 3: Maximum Subarray with Claude

### Input

**Problem Text:**
```
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23
```

**Settings:**
- Language: Java
- User Level: Advanced
- Provider: Claude

### Output

**Explanation:**
This is the classic Maximum Subarray problem, which can be efficiently solved using Kadane's algorithm. The algorithm works by maintaining two variables: `currentSum` (the maximum sum ending at the current position) and `maxSum` (the maximum sum found so far).

For each element in the array:
1. We decide whether to start a new subarray or extend the existing one by taking the maximum of the current element and the sum of the current element plus the previous `currentSum`.
2. We update `maxSum` if the new `currentSum` is greater.

This dynamic programming approach allows us to find the maximum subarray sum in a single pass through the array.

**Code:**
```java
class Solution {
    public int maxSubArray(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        
        int currentSum = nums[0];
        int maxSum = nums[0];
        
        for (int i = 1; i < nums.length; i++) {
            // Either take the current element or add it to the previous subarray
            currentSum = Math.max(nums[i], currentSum + nums[i]);
            
            // Update maxSum if we found a new maximum
            maxSum = Math.max(maxSum, currentSum);
        }
        
        return maxSum;
    }
}
```

**Hints:**
- Consider using Kadane's algorithm for an O(n) solution
- At each position, you have two choices: start a new subarray or extend the current one
- The key insight is determining when to "reset" your current sum
- Edge cases: empty array, all negative numbers

**Time Complexity:** O(n)
**Space Complexity:** O(1)

## Example 4: Merge Two Sorted Lists with Groq

### Input

**Problem Text:**
```
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]
```

**Settings:**
- Language: Go
- User Level: Intermediate
- Provider: Groq

### Output

**Explanation:**
This problem asks us to merge two sorted linked lists into a single sorted linked list. Since the input lists are already sorted, we can use a straightforward approach:

1. Create a dummy head node to simplify edge cases
2. Maintain a current pointer to build the result list
3. Compare the values of the current nodes from both lists
4. Append the smaller value to the result list and advance that list's pointer
5. Continue until we've processed all nodes from both lists

This approach takes advantage of the fact that the lists are already sorted, allowing us to build the merged list in a single pass through both input lists.

**Code:**
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // Create a dummy head node
    dummy := &ListNode{}
    current := dummy
    
    // Traverse both lists and compare values
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }
    
    // Append remaining nodes from either list
    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }
    
    // Return the merged list (skip the dummy head)
    return dummy.Next
}
```

**Hints:**
- Use a dummy head node to simplify handling the first element
- Remember to handle the case where one list is exhausted before the other
- The problem can be solved iteratively or recursively
- Take advantage of the fact that the input lists are already sorted

**Time Complexity:** O(n + m) where n and m are the lengths of the two lists
**Space Complexity:** O(1) for the iterative solution (excluding the output list)
