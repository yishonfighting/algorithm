//给你一个链表数组，每个链表都已经按升序排列。

// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

//

// 示例 1：

// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
// 示例 2：

// 输入：lists = []
// 输出：[]
// 示例 3：

// 输入：lists = [[]]
// 输出：[]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	h := new(miniHeap)

	for i := 0; i < k; i++ {
		h.Push(lists[i])
	}

	dummyHead := new(ListNode)
	pre := dummyHead

	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}

		pre.Next = tmp
		pre = pre.Next
	}

	return dummyHead.Next
}

type miniHeap []*ListNode

func (h miniHeap) Len() int {
	return len(h)
}

func (h miniHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h miniHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *miniHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *miniHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
