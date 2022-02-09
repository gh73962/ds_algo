package leetcode

// 删除链表的倒数第 n 个结点，并且返回链表的头结点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var length int
	for temp := head; temp != nil; temp = temp.Next {
		length++
	}
	var (
		count      int
		targetPrev = length - n
	)
	if length == n {
		return head.Next
	}
	for temp := head; temp != nil; temp = temp.Next {
		count++
		if targetPrev != count {
			continue
		}
		temp.Next = temp.Next.Next
	}

	return head
}

// TODO 快慢指针
