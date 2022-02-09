package list

import "log"

// LeetCode对应编号：206，141，21，19，876
// list is a ring 链表中环的检测
// merge two sorted list 两个有序的链表合并
// delete node which last n 删除链表倒数第 n 个结点
// get mid node of list 求链表的中间结点

type SinglyLinkedListNode struct {
	Next  *SinglyLinkedListNode
	Value int
}

type SinglyLinkedList struct {
	Head   *SinglyLinkedListNode
	length int
}

func NewSingle() *SinglyLinkedList {
	return &SinglyLinkedList{
		Head: &SinglyLinkedListNode{
			Next:  nil,
			Value: 0,
		},
		length: 0,
	}
}

func (s *SinglyLinkedList) Len() int {
	return s.length
}

func (s *SinglyLinkedList) InsertHead(value int) {
	s.length++
	temp := SinglyLinkedListNode{
		Value: value,
		Next:  s.Head.Next,
	}
	s.Head.Next = &temp
}

func (s *SinglyLinkedList) Middle() *SinglyLinkedListNode {
	return nil
}

// Reverse头插法反转
func (s *SinglyLinkedList) Reverse() *SinglyLinkedList {
	newList := NewSingle()
	for temp := s.Head.Next; temp != nil; temp = temp.Next {
		newList.InsertHead(temp.Value)
	}
	return newList
}

//ReverseV2逆置反转
func (s *SinglyLinkedList) ReverseV2() {
	var (
		begin = s.Head.Next
		end   = s.Head.Next.Next
	)
	for end != nil {
		begin.Next = end.Next
		end.Next = s.Head.Next
		s.Head.Next = end
		end = begin.Next
	}
}

func (s *SinglyLinkedList) Print() {
	if s.Head == nil {
		return
	}
	log.Printf("len=%d", s.Len())

	for temp := s.Head.Next; temp != nil; temp = temp.Next {
		log.Printf("value=%d", temp.Value)
	}
}
