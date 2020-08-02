package main

import (
	"container/list"
	"fmt"
)

//2020-0802
func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	node1.Next = node2
	node2.Next = node3
	fmt.Println(reversePrint3(node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint3(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var stack list.List
	for head != nil {
		stack.PushFront(head.Val)
		head = head.Next
	}

	res := make([]int, stack.Len())
	var i int
	for ele := stack.Front(); ele != nil; ele = ele.Next() {
		res[i] = ele.Value.(int)
		i++
	}
	return res
}

func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var eleLen int
	//链表的2次遍历
	//遍历第一次 确定链表的长度
	head2 := head
	for head2 != nil {
		eleLen++
		head2 = head2.Next
	}
	//遍历第二次 进行值的分配
	eles := make([]int, eleLen)
	for head != nil && eleLen-1 >= 0 {
		eles[eleLen-1] = head.Val
		head = head.Next
		eleLen--
	}
	//链表遍历
	return eles
}

//剑指 Offer 06. 从尾到头打印链表
//https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var node *ListNode
	var eleLen int
	//链表的反转
	for head != nil {
		eleLen++
		temp := head.Next
		//新节点，作为头节点
		head.Next = node
		node = head
		head = temp
	}

	//初始化好空间，避免内存空间多次分配
	eles := make([]int, eleLen)
	var i int
	for node != nil && i < eleLen {
		eles[i] = node.Val
		node = node.Next
		i++
	}
	//链表遍历
	return eles
}
