package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node52 := ListNode{5, nil}
	node51 := ListNode{5, &node52}
	node43 := ListNode{4, &node51}
	node42 := ListNode{4, &node43}
	node41 := ListNode{4, &node42}
	node32 := ListNode{3, &node41}
	node31 := ListNode{3, &node32}
	node2 := ListNode{2, &node31}
	node11 := ListNode{1, &node2}
	head := ListNode{1, &node11}

	headNode := deleteDuplicates(&head)
	for headNode != nil {
		fmt.Println(headNode.Val)
		headNode = headNode.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	prev := &ListNode{math.MinInt32, head}
	for curr != nil {
		next := curr.Next
		if next != nil {
			if next.Val == curr.Val {
				for next != nil && next.Val == curr.Val {
					prev.Next = next
					curr = next
					next = curr.Next
				}
				prev.Next = next
				curr = next
				if prev.Val == math.MinInt32 {
					head = curr
				}
			} else {
				prev = curr
				curr = next
			}
		} else {
			curr = next
		}
	}
	return head
}
