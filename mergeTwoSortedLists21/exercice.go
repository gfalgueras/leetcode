package mergeTwoSortedLists21

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Run() {
	l1 := createListNode([]int{1, 5})
	l2 := createListNode([]int{3})

	echoList(mergeTwoSortedLists(l1, l2))
}

func mergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	act := &ListNode{}
	head := act

	for list1 != nil || list2 != nil {
		if list2 == nil {
			act.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else if list1 == nil {
			act.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		} else if list1.Val < list2.Val {
			act.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			act.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}

		act = act.Next
	}

	return head.Next
}

func echoList(node *ListNode) {
	var numbers []int

	for node != nil {
		numbers = append(numbers, node.Val)
		node = node.Next
	}

	fmt.Println(numbers)
}

func createListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0], Next: nil}
	head.Next = createListNode(vals[1:])
	return head
}
