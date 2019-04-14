/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
			return nil
	}
	end := head
	node := head
	for i:= 0; i < n; i++ {
			end = end.Next
	}
	// if end is nil already that means we need to remove the first element
	if end == nil {
			return head.Next
	}
	for end.Next != nil {
			end = end.Next
			node = node.Next
	}
	node.Next = node.Next.Next
	return head
}
