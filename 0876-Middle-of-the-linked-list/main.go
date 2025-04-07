// Solution for LeetCode 876: Middle of the linked list
package main

type ListNode struct {
    Val int
	Next *ListNode
}
func middleNode(head *ListNode) *ListNode {
    leng := length(head)
    mid := leng / 2
    curr := head
    for i:=0; i<mid; i++ {
        curr = curr.Next
    }
    return curr

}
func length(head *ListNode) int {
    leng := 0
    curr := head
    for curr != nil {
        leng++
        curr = curr.Next
    }
    return leng
}