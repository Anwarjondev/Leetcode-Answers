// Solution for LeetCode 3396: Minimum Number of Operations to Make Elements in Array Distint
package main

func minimumOperations(nums []int) int {
	count := 0
	for !isDistinct(nums) {
		count++
		if len(nums) < 3 {
			break
		}
		nums = nums[3:]

	}
	return count
}
func isDistinct(nums []int) bool {
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if mp[nums[i]] == true {
			return false
		}
		mp[nums[i]] = true
	}
	return true
}


