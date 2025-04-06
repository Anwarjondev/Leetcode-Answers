// Solution for LeetCode 167: Two Sum II
package main

import "fmt"


func twoSum(numbers []int, target int) []int {
	var result [2]int
    for i := 0; i<len(numbers); i++{
		num := binarysearch(numbers, i+1, len(numbers)-1, target- numbers[i])
		if num != -1{
			result[0] = i+1
			result[1] = num+1
			break
		}
	}
	return result[:]
}
func binarysearch(num []int, l int,  r int, target int) int {
	for l<=r {
		mid := l+ (r-l)/2
		if num[mid] > target {
			r = mid-1
		} else if num[mid] < target{
			l = mid+1
		} else {
			return mid
		}
	}
	return -1
}


func main() {
	slc := []int{}
	number := 0
	fmt.Println(twoSum(slc, number))
}