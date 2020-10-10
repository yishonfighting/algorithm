package main

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//说明：不允许修改给定的链表。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	s, f := head, head
	for f != nil {
		s = s.Next
		if f.Next == nil {
			return nil
		}
		f = f.Next.Next
		if s == f {
			p := head

			for s != p {
				s = s.Next
				p = p.Next
			}
			return p

		}
	}
	return nil
}
