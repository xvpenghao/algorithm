package main

import "fmt"

func main() {
	rootA := &TreeNode3{
		Val: 4,
		Left: &TreeNode3{
			Val: 2,
			Left: &TreeNode3{
				Val: 1,
			},
			Right: &TreeNode3{
				Val: 3,
			},
		},
		Right: &TreeNode3{
			Val: 7,
			Left: &TreeNode3{
				Val: 6,
			},
			Right: &TreeNode3{
				Val: 9,
			},
		},
	}
	// 4213769
	// 4796231
	res := mirrorTree(rootA)
	print2(res)
}

// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
//

type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

// 剑指 Offer 27. 二叉树的镜像
func mirrorTree(root *TreeNode3) *TreeNode3 {
	if root == nil {
		return nil
	}
	// 左右节点的交换
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

func print2(root *TreeNode3) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	print2(root.Left)
	print2(root.Right)
}
