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

// MergeTwoLists 合并两个有序链表为一个
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//虚拟头节点
	dummy := &ListNode{1, nil}
	p, p1, p2 := dummy, l1, l2
	for p1 != nil && p2 != nil {
		if p1.value > p2.value {
			p.next = p2
			p2 = p2.next
		} else {
			p.next = p1
			p1 = p1.next
		}
		p = p.next
	}
	if p1 != nil {
		p.next = p1
	}
	if p2 != nil {
		p.next = p2
	}
	return dummy.next
}

// MergeKLists todo 自定义实现优先级队列
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var (
		dummy = ListNode{-1, nil}
		p     = dummy
	)
	return &p
}

// FindFromEnd 返回链表倒数第K个节点
func FindFromEnd(head *ListNode, index int) *ListNode {
	p1 := head
	//先让p1走k步
	for i := 1; i < index; i++ {
		p1 = p1.next
	}
	p2 := head
	//p1和p2走n-k步
	for p1.next != nil {
		p2, p1 = p2.next, p1.next
	}
	return p2
}

// MiddleNode 找出list中位数节点
func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		//走两步第到末尾到时候，走一步的就到中点
		slow, fast = slow.next, fast.next.next
	}
	return slow
}

/*
HasCycle 链表是否包含环
每当慢指针slow前进一步，快指针slow就前进两步。
如果fast最终遇到空指针，说明链表中没有环；
如果fast最终和slow相遇，那肯定是fast超过了slow一圈，说明链表中含有环。
*/
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

/*
DetectCycle 环链表寻找起点
当快慢指针相遇时，让其中任一个指针指向头节点
然后让它俩以相同速度前进，再次相遇时所在的节点位置就是环开始的位置
*/
func DetectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		fast, slow = fast.next.next, slow.next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.next == nil {
		// fast 遇到空指针说明没有环
		return nil
	}
	//重新指向头节点
	slow = head
	//快慢指针同步前进，相交点就是起点
	for slow != fast {
		fast, slow = fast.next, slow.next
	}
	return slow
}

// GetIntersectionNode 查看两个链表有没有相交点
func GetIntersectionNode(nodeA *ListNode, nodeB *ListNode) *ListNode {
	//p1,p2指向a,b
	p1, p2 := nodeA, nodeB
	for p1 != p2 {
		// p1 走一步，如果走到 A 链表末尾，转到 B 链表
		if p1 == nil {
			p1 = nodeB
		} else {
			p1 = p1.next
		}
		// p2 走一步，如果走到 B 链表末尾，转到 A 链表
		if p2 == nil {
			p2 = nodeA
		} else {
			p2 = p2.next
		}
	}
	return p1

}

// RemoveNthFromEnd 删除倒数第N个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	//虚拟头节点
	dummy := &ListNode{-1, nil}
	dummy.next = head
	//要删除第n个，先找到第n+1个
	x := FindFromEnd(dummy, n+1)
	//删除操作
	x.next = x.next.next
	return dummy.next
}

//========================================================

// reverse 反转链表的迭代实现
func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		//go语言语法,多变量同时赋值
		pre, cur, cur.next = cur, cur.next, pre
	}
	return pre
}

// reverse2 反转链表的递归实现
func reverse2(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	last := reverse2(head.next)
	head.next.next, head.next = head, nil
	return last
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

//========================================================

func Reverse(a *ListNode, b *ListNode) *ListNode {
	var pre *ListNode = nil
	cur, next := a, a
	for cur != b {
		next, cur.next, pre, cur = cur.next, pre, cur, next
	}
	return pre

}

func ReverseByK(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.next
	}
	newHead := Reverse(a, b)
	a.next = ReverseByK(b, k)
	return newHead
}

//=======================================================

// PalindRome 判断回文字符串
func PalindRome(str string, l int, r int) string {
	for l >= 0 && r < len([]rune(str)) && str[l] == str[r] {
		l--
		r++
	}
	return str[l+1 : r-l-1]
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
	head1 := GetNode()
	head2 := GetNode()
	head := MergeTwoLists(head2, head1)
	TraversePrint(head)
	//TraversePrint(reverse2(head))
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
