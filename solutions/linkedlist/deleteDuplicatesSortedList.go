package main

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func deleteDuplicates1(head *ListNode1) *ListNode1 {
	node := head
	for node != nil {
		temp := node.Next
		if temp != nil {
			if temp.Val == node.Val {
				node.Next = temp.Next
			} else {
				node = temp
			}
		} else {
			node = temp
		}
	}
	return head
}
