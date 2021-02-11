package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	res := getKthFromEnd2(head, 2)
	fmt.Println(res)
}

// 剑指 Offer 22. 链表中倒数第k个节点
// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return nil
	}
	// 1 ，2 ，3 ，4 ，5 k=2,
	// 结果 4，5
	// 我需要走3步骤才能获取到 4,5 ,怎么才能走3步呢，走k=2步就剩下走三步了
	front, tail := head, head
	for i := 1; i <= k; i++ {
		// 可能你输入的 k 比len(head)的长度大
		if front == nil {
			return nil
		}
		front = front.Next
	}

	// front 和 tail 同时走
	for front != nil {
		front = front.Next
		tail = tail.Next
	}
	return tail
}

// 给定一个链表: 1->2->3->4->5, 和 k = 2.
// 返回链表 4->5.
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return nil
	}

	tmpHead := head
	var nodeLen int
	for tmpHead != nil {
		nodeLen++
		tmpHead = tmpHead.Next
	}
	startPosition := nodeLen - k

	// 这是要的全部
	if startPosition == 0 {
		return head
	}

	index := 1
	for head != nil {
		// 找到要截取的位置了
		head = head.Next
		if index == startPosition {
			return head
		}
		index++

	}
	return nil
}
