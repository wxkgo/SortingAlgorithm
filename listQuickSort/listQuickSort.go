package sort

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func quickSort(head *ListNode,end *ListNode)  {
	if head != end {
		mid := sr(head,end)
		quickSort(head,mid)
		quickSort(mid.Next,end)
	}
}

func sr(head *ListNode,end *ListNode) *ListNode {
	p := head.Next
	small := head
	for p != end {
		if p.Val < head.Val {
			small = small.Next
			small.Val,p.Val = p.Val,small.Val
		}
		p = p.Next
	}
	head.Val,small.Val = small.Val,head.Val
	return small
}

func pr(head *ListNode)  {
	for head != nil {
		fmt.Print(head.Val," ")
		head = head.Next
	}
}
//func Test()  {
//	node9 := &ListNode{Val:9}
//	node8 := &ListNode{Val:5,Next:node9}
//	node7 := &ListNode{Val:2,Next:node8}
//	node6 := &ListNode{Val:7,Next:node7}
//	node5 := &ListNode{Val:10,Next:node6}
//	node4 := &ListNode{Val:8,Next:node5}
//	node3 := &ListNode{Val:6,Next:node4}
//	node2 := &ListNode{Val:4,Next:node3}
//	node1 := &ListNode{Val:1,Next:node2}
//	head := &ListNode{Val:3,Next:node1}
//	quickSort(head,nil)
//	pr(head)
//}
