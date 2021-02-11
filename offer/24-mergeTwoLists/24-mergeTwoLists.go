package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := mergeTwoLists2(head, head2)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 剑指 Offer 25. 合并两个排序的链表
// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil && l1 != nil {
		return l1
	}

	if l1 == nil && l2 != nil {
		return l2
	}
	// 两个链表都是升序的
	// 输入：1->2->4, 1->3->4
	// 输出：1->1->2->3->4->4
	// 能不能同时遍历,两者 相互比较，然后添加到一个新的节点然后，最后返回
	cur := new(ListNode)
	// head2 执行 head
	head2 := cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			// 不用创建一个新的节点
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		// cur 到下一步
		cur = cur.Next
	}

	// 可能 l1 或者 l2 这两个其中一个没有值
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return head2.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil && l1 != nil {
		return l1
	}

	if l1 == nil && l2 != nil {
		return l2
	}
	// 两个链表都是升序的
	// 输入：1->2->4, 1->3->4
	// 输出：1->1->2->3->4->4
	// 能不能同时遍历,两者 相互比较，然后添加到一个新的节点然后，最后返回
	head := new(ListNode)
	// head2 执行 head
	head2 := head
	for l1 != nil || l2 != nil {
		tmpNode := new(ListNode)
		// 比较大小添加新值
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				// 添加 l1 的值
				tmpNode.Val = l1.Val
				l1 = l1.Next
			} else {
				// 添加 l2的值
				tmpNode.Val = l2.Val
				l2 = l2.Next
			}
		} else if l2 == nil && l1 != nil {
			tmpNode.Val = l1.Val
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			tmpNode.Val = l2.Val
			l2 = l2.Next
		}

		head.Next = tmpNode
		head = tmpNode
	}

	return head2.Next
}
