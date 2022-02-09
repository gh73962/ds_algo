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
	for ; list2 != nil; list2 = list2.Next {

	}
	return nil
}
