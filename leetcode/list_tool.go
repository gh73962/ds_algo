package leetcode

import "log"

func NewHead(length int) *ListNode {
	var head = ListNode{
		Val:  length,
		Next: nil,
	}

	for i := length - 1; i > 0; i-- {
		temp := ListNode{
			Val:  i,
			Next: head.Next,
		}
		// log.Printf("temp = %+v", temp)
		head.Next = &temp
	}

	return &head
}

func ListPrint(s *ListNode) {
	for temp := s; temp != nil; temp = temp.Next {
		log.Printf("ListPrint value=%d", temp.Val)
	}
}
