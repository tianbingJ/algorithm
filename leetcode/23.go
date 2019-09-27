package leetcode



func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		l1 := lists[0]
		l2 := lists[1]
		lists = lists[2:]
		lists = append(lists, mergeLists(l1, l2))
	}
	return lists[0]
}

func mergeLists(l1, l2 * ListNode) *ListNode {
	p := l1
	q := l2
	var result, tailNode *ListNode
	for p != nil || q != nil {
		var node *ListNode
		if p == nil {
			node = q
			q = q.Next
		} else if q == nil {
			node = p
			p = p.Next
		} else {
			if p.Val < q.Val {
				node = p
				p = p.Next
			} else {
				node = q
				q = q.Next
			}
		}
		if result == nil {
			result = node
			tailNode = node
		} else {
			tailNode.Next = node
			tailNode = node
		}
	}
	return result
}