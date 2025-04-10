// Solution for LeetCode 54: Spiral Matrix
package main

func spiralOrder(matrix [][]int) []int {
    var result []int
	var r []int
	length := 0
	for _, v := range matrix {
		length += len(v)
	}
	indexr, indexu, indexl, indexb := len(matrix[0])-1,0,0,len(matrix)-1
	for len(result) != length {
		if indexu < len(matrix[0]) {
			r, indexu = up(matrix, indexu)
			result = add(r, result)
		}
		if indexr >= 0 {
			r, indexr = right(matrix, indexr)
			result = add(r, result)
		}
		if indexb < len(matrix) {
			r, indexb = bottom(matrix, indexb)
			result = add(r, result)
		}
		if indexl < len(matrix) {
			r, indexl = left(matrix, indexl)
			result = add(r, result)
		}
		
	}
	return result
	
}
func up(m [][]int, index int) ([]int, int){
	var nums []int
	for i := 0; i<len(m[index]); i++ {
		if m[index][i] != 101 {
			nums = append(nums, m[index][i])
			m[index][i] = 101
		}
	}
	index++
	return nums, index
}
func right(m [][]int, index int) ([]int, int) {
	var nums []int
	for i:=0; i<len(m); i++ {
		if m[i][index] != 101 {
			nums = append(nums, m[i][index])
			m[i][index] = 101
		}
	}
	index --
	return nums, index
}
func bottom(m [][]int, index int) ([]int, int) {
	var nums []int
	for i:= len(m[0])-1; i>=0; i-- {
		if m[index][i] != 101 {
			nums = append(nums, m[index][i])
			m[index][i] = 101
		}
	}
	index --
	return nums, index
}
func left(m [][]int, index int) ([]int, int) {
	var nums []int
	for i := len(m)-1; i>= 0; i-- {
		if m[i][index] != 101 {
			nums = append(nums, m[i][index])
			m[i][index] = 101
		}
	}
	index ++
	return nums, index
}
func add(r , result []int) []int {
	for i:= 0; i<len(r); i++ {
		result = append(result, r[i])
	}
	return result
}