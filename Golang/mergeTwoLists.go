/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	head := node
	for l1 != nil || l2 != nil {
			node.Next = &ListNode{}
			node = node.Next
			if l1 == nil {
					node.Val = l2.Val
					l2 = l2.Next
			} else if l2 == nil {
					node.Val = l1.Val
					l1 = l1.Next
			} else if l1.Val < l2.Val {
					node.Val = l1.Val
					l1 = l1.Next
			} else {
					node.Val = l2.Val
					l2 = l2.Next
			}
	}
	return head.Next
}
