package leetcode

// 给定一个头结点为 head 的非空单链表，返回链表的中间结点。
// 如果有两个中间结点，则返回第二个中间结点。

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var length int
	for temp := head; temp != nil; temp = temp.Next {
		length++
	}
	switch length {
	case 1:
		return head
	case 2:
		return head.Next
	default:
		temp := head
		mid := length / 2
		for mid > 0 {
			temp = temp.Next
			mid--
		}
		return temp
	}
}

// 快慢指针也可
