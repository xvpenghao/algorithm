package main

// 剑指 Offer 18. 删除链表的节点
// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
func main() {
	head := &ListNode2{
		Val: 1,
		Next: &ListNode2{
			Val: 2,
			Next: &ListNode2{
				Val: 3,
				Next: &ListNode2{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	deleteNode2(head, 3)
}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func deleteNode(head *ListNode2, val int) *ListNode2 {
	// 删除头节点和 删除 尾部节点
	// a 1,2,3,4
	// 2,3,4
	// current 的指向 和head的指向都是同一个地址
	current := head
	for head.Val == val {
		return head.Next
	}

	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
			break
		}
		// 改变 head的指向 但是并没有改变的 current的指向
		head = head.Next
		// 这样会导致 head 和 current的节点 都会缺失
		// head.next = head.next.next

	}
	return current
}

func deleteNode2(head *ListNode2, val int) *ListNode2 {
	tmpNode := &ListNode2{
		Val:  0,
		Next: head,
	}

	tmpHead := tmpNode
	for tmpNode.Next != nil {
		if tmpNode.Next.Val == val {
			// 删除节点
			tmpNode.Next = tmpNode.Next.Next
			break
		}
		// 节点指向的移动
		tmpNode = tmpNode.Next
	}
	return tmpHead.Next
}
