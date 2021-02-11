package main

import (
	"fmt"
	"strconv"
)

/**
*@Author:徐鹏豪
*@Time: 2021/2/11 09:06
*@Description: 剑指 Offer 28. 对称的二叉树
 */

func main() {
	rootA := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	// 3 4 2 1
	// 4 3 2 1
	fmt.Println(isSymmetric2(rootA))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
// 剑指 Offer 28. 对称的二叉树
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compare(root.Left, root.Right)
}

func compare(left *TreeNode, right *TreeNode) bool {
	// 都为nil
	if left == nil && right == nil {
		return true
	}
	// 任何一方为nil
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 对root进行 先序遍历，然后把遍历结构用字符串+链接起来
	// 然后 进行 镜像树操作 ，
	// 然后先序遍历，用字符串+连接起来，比较两者
	var str1, str2 string
	// 原始值
	str1 = precedence(root, str1)
	// 镜像操作后的值
	str2 = precedence(mirrorTree2(root), str2)
	return str1 == str2
}

func precedence(root *TreeNode, str string) string {
	if root == nil {
		// 用 nil 来作为 遍历是 为空区分，要不然可能存在遍历结果一样，但是两者的树形结构不一样
		return str + "NIL"
	}

	str += strconv.Itoa(root.Val)
	// 这里面 结果已经加好了
	str = precedence(root.Left, str)
	str = precedence(root.Right, str)
	return str
}

// 镜像操作
func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 左右值 的交换
	root.Left, root.Right = root.Right, root.Left
	mirrorTree2(root.Left)  // 左边的交换
	mirrorTree2(root.Right) // 右边的交换
	return root
}
