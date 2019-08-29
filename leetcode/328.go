package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	oddHead, oddTail := head, head
	evenHead, evenTail := head.Next, head.Next

	for p := evenTail.Next; p != nil; {
		oddTail.Next = p
		oddTail = oddTail.Next
		p = p.Next
		if p == nil {
			break
		}
		evenTail.Next = p
		evenTail = evenTail.Next
		p = p.Next
	}
	evenTail.Next = nil
	oddTail.Next = evenHead
	return oddHead
}
