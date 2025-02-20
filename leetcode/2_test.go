package leetcode

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l3 := addTwoNumbers(l1, l2)
	t.Log(l3)
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	i1 := getIntFromNode(l1)
	i2 := getIntFromNode(l2)
	sum := i1 + i2

	return getNodeFromInt(sum)
}

func getNodeFromInt(i int64) *ListNode {
	l := &ListNode{}
	if i == 0 {
		l.Val = 0
	} else {
		l.Val = int(i % 10)
		if i/10 != 0 {
			l.Next = getNodeFromInt(i / 10)
		}
	}
	return l
}

func getIntFromNode(l *ListNode) int64 {
	var i int64
	if l.Next == nil {
		i = int64(l.Val)
	} else {
		i = getIntFromNode(l.Next)*10 + int64(l.Val)
	}

	return i
}
