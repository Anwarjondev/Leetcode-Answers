## LeetCode 167: Two Sum II

### Description:
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

### Link:
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

### Solution:
Solution inside the main.go file our main task is we are given a sorted array(list) and a target we should find two any numbers inside the array which is these two numbers adding result equal to the target.

### Explanation:
Array is sorted that is why i use the Binary Search main part is this when i call the binary search every time i+1 why because the resukt should be different two numbers i mean bith numbers position should be different it is possible to the both numbers value is equal but their position must be different. and also i creat the array with fixed size that is why this algorithm's space complexity equal to the O(1).