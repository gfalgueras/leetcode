package addtwonumbers

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	head := &ListNode{}
	res := head

	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		suma := sum + carry

		head.Next = &ListNode{Val: suma % 10, Next: nil}
		head = head.Next

		carry = suma / 10
	}

	if carry > 0 {
		head.Next = &ListNode{Val: carry}
	}

	return res.Next
}

func echoList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func Run() {
	l1 := createListNode([]int{5})
	l2 := createListNode([]int{5})

	sumList := addTwoNumbers(l1, l2)

	echoList(sumList)
}

func createListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0], Next: nil}
	head.Next = createListNode(vals[1:])
	return head
}
