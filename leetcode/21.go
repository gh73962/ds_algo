package leetcode

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil && list2 != nil {
		return list2
	}
	if list1 != nil && list2 == nil {
		return list1
	}
	var (
		newList, tail *ListNode
		head1         = list1
		head2         = list2
	)

	for head1 != nil || head2 != nil {

		var temp ListNode
		switch {
		case head1 == nil:
			temp.Val = head2.Val
			head2 = head2.Next
		case head2 == nil:
			temp.Val = head1.Val
			head1 = head1.Next
		case head1.Val <= head2.Val:
			temp.Val = head1.Val
			head1 = head1.Next
		case head1.Val > head2.Val:
			temp.Val = head2.Val
			head2 = head2.Next
		}

		if newList == nil {
			newList = &temp
			tail = &temp
			continue
		}

		tail.Next = &temp
		tail = tail.Next
	}
	return newList
}
