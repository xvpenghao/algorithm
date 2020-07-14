package main

import (
	"fmt"
)

//这个是头也有数据
type ListNode struct {
	Val  int
	Next *ListNode
}

//TODO 链表
func main() {
	L1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	addNode(L1, 2)
	addNode(L1, 3)
	addNode(L1, 2)
	addNode(L1, 1)
	fmt.Println(isPalindrome2(L1))
	//traverse(res)
}

//用快慢指针发
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//12321 假如是回文数，
	//则找到中间节点
	//反转中间的后半部分，
	//对比中间的前半分 和 后半分
	//提问，为什么使用前后指针能找到 slow能找中间节点， 因为 fast是slow的两倍

	//找到中节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	//发展中间节点的后半部分
	var tail *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = tail
		tail = slow
		slow = tmp
	}

	//比较两个前后节点值
	for tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}

	return true
}

//判断链表是否有回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	j := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[j-i] {
			return false
		}
	}

	return true
}

func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//环形判断
	//开始大家都是指向起点
	slowPointer, fastPointer := head, head
	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
		if slowPointer == fastPointer {
			//指向一个头部指针和一个相遇指针
			for head != nil && slowPointer != nil {
				//两者不相等则移动
				if head != slowPointer {
					head = head.Next
					slowPointer = slowPointer.Next
				}
			}
			return slowPointer
		}
	}

	//找出环点
	return nil
}

//在环形链表中，返回入环后的第一个节点
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[*ListNode]*ListNode)
	for head != nil {
		if node, ok := m[head]; ok {
			return node
		}
		m[head] = head
		head = head.Next
	}
	return nil

}

//快慢指针
func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	//在环形中，快慢指针最终都会相遇的
	slowPointer, fastPointer := head, head
	//fasterPointer只要判断好了，slowPointer就不会出现nil了
	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}

//判断链表中是否有环- 使用hashMap
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}

		m[head] = true
		head = head.Next
	}

	return false
}
func reorderList(head *ListNode) {

}

// -9,3 | 5,7
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	//利用升序
	//两个节点值笔记，将比较的值插入到第三方变量
	//这里还不确定next是谁，需要比较才能出来
	dummy := &ListNode{
		Val: 0,
	}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	//将剩余节点进行插入
	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return dummy.Next
}

//leetCode 21,合并两个有序链表
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//正常情况下的插入
	//2,3 1,2,3
	dummy := &ListNode{
		Val:  0,
		Next: l1,
	}
	l1 = dummy
	for l2 != nil {
		l1Temp := l1
		l2Temp := l2.Next
		for l1.Next != nil {
			//x,2,4  3,4,5
			if l2.Val <= l1.Next.Val {
				break
			}
			l1 = l1.Next
		}
		l2.Next = l1.Next
		l1.Next = l2

		l2 = l2Temp
		l1 = l1Temp
	}

	return dummy.Next
}

//leetCode 92题 链表的指定位置反转
//输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
// 参考链接 https://blog.csdn.net/qq_43152052/article/details/103648554
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	//找到m的前去节点,这样做好处是第一个位置也有前驱
	pre := dummy
	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	//1，2，3，4，5，6
	//1，2，4，3，5，6
	//获取要开始反转的节点
	cur := pre.Next
	for i := m; i < n; i++ {
		//类似头插入，这是每次将需要反转的节点元素插入到 pre后面，腾出来的节点
		temp := cur.Next
		//链接链表,如果不连接 那么 cur的Next就断了
		cur.Next = temp.Next
		//开始节点的插入
		//pre 2,3,4,5,nil 3插入到pre后面，pre.Next为最后的节点
		temp.Next = pre.Next
		pre.Next = temp
	}
	return dummy.Next

}

//链表的反转方法2
func reverseList2(head *ListNode) *ListNode {
	//输入: 1->2->3->4->5->NULL
	//输出: 5->4->3->2->1->NULL
	var prev *ListNode
	for head != nil {
		//临时保存当前节点
		temp := head.Next
		head.Next = prev
		//prev有head次数的内容了
		prev = head
		//有了下一个值了
		head = temp
	}
	return prev
}

//单链表反转-方法1，
func reverseList(head *ListNode) *ListNode {
	//输入: 1->2->3->4->5->NULL
	//输出: 5->4->3->2->1->NULL
	//遍历链表，使用头插法
	tListNode := new(ListNode)
	for head != nil {
		node := &ListNode{
			Val:  head.Val,
			Next: nil,
		}
		if tListNode.Next != nil {
			node.Next = tListNode.Next
		}
		tListNode.Next = node
		head = head.Next
	}
	return tListNode.Next
}

//LeetCode 82
//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
func deleteDuplicates3(head *ListNode) *ListNode {
	//使用一个dump的结点来存一下
	//这样这个head就有一个头的结点来存放了
	//因为头节点也可能被删除，所以新增一个节点来保护
	dumpNode := &ListNode{
		Val:  0,
		Next: head,
	}
	head = dumpNode

	//x,1,1,1,3,4
	//3,4
	//从这新的head结构上开始思考
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			//记录被删除的值，然后循环删除
			rmVal := head.Next.Val
			//不会出现空指针的
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}

	return dumpNode.Next
}

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
func deleteDuplicates(head *ListNode) *ListNode {
	L2 := head
	for L2 != nil {
		//代码优化的空间
		if L2.Next == nil {
			break
		}
		//可能还有连续的重复的，则需要删除完
		for L2.Next != nil && L2.Val == L2.Next.Val {
			L2.Next = L2.Next.Next
		}
		L2 = L2.Next
	}
	return head
}

//链表的新建,头结点也存放数据
func addNode(L *ListNode, data int) {
	if L == nil {
		return
	}
	node := &ListNode{
		Val:  data,
		Next: nil,
	}

	//循环遍历并插入
	L2 := L
	//找到要插入的位置
	for L2.Next != nil {
		L2 = L2.Next
	}
	L2.Next = node
}

//列表的遍历
func traverse(L *ListNode) {
	L2 := L
	for L2 != nil {
		fmt.Print(L2.Val, ",")
		L2 = L2.Next
	}
}
