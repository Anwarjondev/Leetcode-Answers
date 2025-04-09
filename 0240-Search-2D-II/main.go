// Solution for LeetCode 240: Search a 2D Matrix II
package main

func searchMatrix(matrix [][]int, target int) bool {
	for _, v := range matrix {
		if binarysearch(v, target) {
			return true
		}
	}
	return false
}
func binarysearch(nums []int, t int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l+(r-l)/2
		if nums[mid] == t {
			return true
		} else if nums[mid] < t{
			l = mid+1
		} else {
			r = mid-1
		}
	}
	return false
}