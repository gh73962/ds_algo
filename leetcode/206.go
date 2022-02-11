package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newList *ListNode
	for temp := head; temp != nil; temp = temp.Next {
		tempData := ListNode{
			Val:  temp.Val,
			Next: newList,
		}
		newList = &tempData
	}
	return newList
}
