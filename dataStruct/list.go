/**
 * @author: lqs
 * @date: 2022/1/6
 */
package main

import (
	"fmt"
)

//全局变量必须使用var声明,且不能为空，定义了可以不使用
var successor = new(ListNode)

// ListNode 链表定义
type ListNode struct {
	value int
	next  *ListNode
}

//反转链表的实现
func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		//go语言语法,多变量同时赋值
		pre, cur, cur.next = cur, cur.next, pre
	}
	return pre
}

// ReverseN 反转前n个元素
func ReverseN(head *ListNode, start int) *ListNode {
	if start == 1 {
		successor = head.next
		return head
	}
	last := ReverseN(head.next, start-1)
	head.next.next = head
	head.next = successor
	return last
}

// ReverseBetween 反转链表一部分
func ReverseBetween(head *ListNode, start int, end int) *ListNode {
	if end == 1 {
		return ReverseN(head, start)
	}
	head.next = ReverseBetween(head.next, start-1, end-1)
	return head
}

// TraversePrint 遍历打印
func TraversePrint(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.value)
	TraversePrint(head.next)
}

// ReversePrint 递归遍历单链表，倒序打印链表元素
func ReversePrint(head *ListNode) {
	if head == nil {
		return
	}
	ReversePrint(head.next)
	fmt.Println(head.value)
}

func main() {
	head := GetNode()
	TraversePrint(reverse(head))
}

func GetNode() *ListNode {
	head := new(ListNode)
	head.value = 1
	ln2 := new(ListNode)
	ln2.value = 2
	ln3 := new(ListNode)
	ln3.value = 3
	ln4 := new(ListNode)
	ln4.value = 4
	ln5 := new(ListNode)
	ln5.value = 5
	head.next = ln2
	ln2.next = ln3
	ln3.next = ln4
	ln4.next = ln5
	return head
}
