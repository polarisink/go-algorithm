package main

import "fmt"

func main() {
	a := ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}
	b := ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}
	result := addTwoNumbers(&a, &b)
	LogListNode(result)
}

func LogListNode(l *ListNode) {
	if l == nil {
		return
	}
	fmt.Println(l.Val)
	LogListNode(l.Next)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2, dummy := l1, l2, ListNode{Val: -1}
	p, carry := dummy, 1
	for p1 != nil || p2 != nil || carry > 0 {
		val := carry
		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}
		// 处理进位情况
		carry = val / 10
		val = val % 10
		// 构建新节点
		p.Next = &ListNode{
			Val: val,
		}
		p = *p.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
