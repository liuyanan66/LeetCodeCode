package main

import (
	"fmt"
	//"runtime/debug"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 普通的算法
// 如果是第二个节点则插入尾节点前面
// 如果是单数节点则插入到尾节点后面
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	new_head, tail, pre_tail, cur := head, head, head, head.Next
	new_head.Next = nil
	cur_num := 1
	for cur != nil {
		cur_num++
		tmp_node := cur
		cur = cur.Next
		if cur_num%2 == 0 {
			if cur_num == 2 {
				tmp_node.Next = new_head
				pre_tail = tmp_node
				new_head = tmp_node
			} else {
				pre_tail.Next = tmp_node
				tmp_node.Next = tail
				pre_tail = tmp_node
			}
		} else {
			tail.Next = tmp_node
			pre_tail = tail
			tail = tmp_node
			tail.Next = nil
		}
	}
	return new_head
}

//
func swapPairsTwo(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next.Next
	new_head := head.Next
	new_head.Next = head
	new_head.Next.Next = nil
	tail := new_head.Next
	for cur != nil && cur.Next != nil {
		tmp := cur
		cur = cur.Next.Next
		tail.Next = tmp.Next
		tmp.Next.Next = tmp
		tail = tmp
		tail.Next = nil
	}
	tail.Next = cur
	return new_head
}

// 加上一个空的头节点Next指针指向Head指针
// 然后对满足节点的指针pre节点的Next和Next.Next都不为空的指针进行操作
// 这样避免了单指针节点为奇数的时候单独处理
func swapPairsTwoSimple(head *ListNode) *ListNode {
	pre := new(ListNode)
	pre.Next = head
	new_head := pre
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := pre.Next.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
		pre = a
	}
	return new_head.Next
}

//递归的算法
func swapPairsRecur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l1 := head.Next
	head.Next = swapPairsRecur(l1.Next)
	l1.Next = head

	return l1
}

// 遍历链表
func traverseList(head *ListNode) {
	if head == nil {
		return
	}
	cur := head
	for cur != nil {
		fmt.Printf("%d  ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func main() {
	var head, tail *ListNode
	for i := 1; i < 10; i++ {
		if head == nil {
			head = new(ListNode)
			tail = head
			tail.Val = i
		} else {
			tail.Next = new(ListNode)
			tail = tail.Next
			tail.Val = i
		}
	}
	new_head := head.Next
	new_head.Next = head
	new_head.Next.Next = nil
	traverseList(new_head)
	//debug.PrintStack()
}
