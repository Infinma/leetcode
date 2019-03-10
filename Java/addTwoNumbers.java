/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */


// Recursive
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        if (l1 == null) {
            return l2;
        }
        if (l2 == null) {
            return l1;
        }
        int sum = l1.val + l2.val;
        ListNode node = new ListNode(sum % 10);
        if (sum > 9) {
            if (l1.next == null) {
                l1.next = new ListNode(1);
            } else if (l2.next == null) {
                l2.next = new ListNode(1);
            } else {
                l1.next.val += 1;
            }
        }
        node.next = addTwoNumbers(l1.next, l2.next);
        return node;
    }
}

// Loops
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        if (l1 == null) return l2;
        if (l2 == null) return l1;
        int carry = 0;
        int sum = 0;
        ListNode n = new ListNode(-1);
        ListNode head = n;
        while (l1 != null || l2 != null) {
            if (l1 == null) {
                sum = l2.val + carry;
                l2 = l2.next;
            } else if (l2 == null) {
                sum = l1.val + carry;
                l1 = l1.next;
            } else {
                sum = l1.val + l2.val + carry;
                l1 = l1.next;
                l2 = l2.next;
            }
            n.next = new ListNode(sum % 10);
            carry = sum / 10;
            n = n.next;
        }
        if (carry != 0) {
            n.next = new ListNode(carry);
        }
        return head.next;
    }
}