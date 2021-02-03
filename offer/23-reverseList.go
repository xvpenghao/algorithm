package main

import "fmt"

func main() {
	head := &ListNode4{
		Val: 1,
		Next: &ListNode4{
			Val: 2,
			Next: &ListNode4{
				Val: 3,
				Next: &ListNode4{
					Val: 4,
					Next: &ListNode4{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	res := reverseList(head)
	fmt.Println(res)
}

type ListNode4 struct {
	Val  int
	Next *ListNode4
}

// 剑指 Offer 24. 反转链表
// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
func reverseList(head *ListNode4) *ListNode4 {
	/*
		输入: 1->2->3->4->5->NULL
		输出: 5->4->3->2->1->NULL
	*/
	// 头为nil 或者只有一个节点则直接返回 head
	if head == nil || head.Next == nil {
		return head
	}

	var head2 *ListNode4
	for head != nil {
		// 提前保存下一个
		tmp := head.Next
		// 下一个和之前的已经 断开 连接了
		// head.Next 下一个头，head2上一个头
		head.Next = head2
		head2 = head
		head = tmp
	}

	return head2
}
