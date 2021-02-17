package main

import "fmt"

/**
*@Author:徐鹏豪
*@Time: 2021/2/17 11:54
*@Description:
 */
func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	res := treeToDoublyList(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
// 剑指 Offer 36. 二叉搜索树与双向链表
// 将 二叉搜索数 转换为 有序的升序链表 利用二叉树搜索树的性质， 中序遍历后 是一个有序升序的结果
func treeToDoublyList(root *TreeNode) *TreeNode {
	var dfs func(cur *TreeNode)
	// 需要将 根和尾部链接起来
	var head, pre *TreeNode
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		// 左
		dfs(cur.Left)
		// 根
		if pre != nil {
			pre.Right = cur
		} else {
			head = cur
		}
		cur.Left = pre
		pre = cur
		// 右
		dfs(cur.Right)
	}
	if root == nil {
		return nil
	}
	dfs(root)

	// 头尾 相互指向
	head.Left, pre.Right = pre, head
	return head
}
