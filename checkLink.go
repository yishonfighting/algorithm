package main

//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

	l := []int{}
	for head != nil {
		l = append(l,head.Val)
		head = head.Next
	}

	for i:=0; i < len(l); i++ {

	}
	start , end := 0 , len(l)-1
	for start < end {
		if l[start] == l[end] {
			start++
			end--
			continue
		} else {
			return false
		}
	}

	return true
}