package leetcode

import "math/rand"

type Solution2 struct {
	head * ListNode
}


/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor3(head *ListNode) Solution2 {
	return Solution2{head:head}
}


/** Returns a random node's value. */
func (this *Solution2) GetRandom() int {
	cur := this.head
	count := 1
	result := 0
	for cur != nil {
		randN := rand.Intn(count)
		if randN == 0 {
			result = cur.Val
		}
		cur = cur.Next
		count ++
	}
	return result
}
