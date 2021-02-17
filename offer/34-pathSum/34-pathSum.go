package main

import "fmt"

/**
*@Author:徐鹏豪
*@Time: 2021/2/16 15:08
*@Description:
 */
func main() {
	/*
	         5
	        / \
	       4   8
	      /   / \
	     11  13  4
	    /  \    / \
	   7    2  5   1
	*/
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	sumEle := pathSum2(root, 22)
	for _, val := range sumEle {
		fmt.Println(val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指 Offer 34. 二叉树中和为某一值的路径
// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
func pathSum2(root *TreeNode, sum int) [][]int {

	var sumEle [][]int
	var stack []int
	var recurFn func(root *TreeNode, target int)
	// 学到了一个新技能，匿名函数如何自己调用自己
	recurFn = func(root *TreeNode, target int) {
		if root == nil {
			return
		}

		stack = append(stack, root.Val)
		target -= root.Val

		// 表示是叶子节点，并且是和 sum一致 已经相等
		if target == 0 && root.Left == nil && root.Right == nil {
			tmpEle := make([]int, len(stack))
			copy(tmpEle, stack)
			sumEle = append(sumEle, tmpEle)
		}

		recurFn(root.Left, target)
		recurFn(root.Right, target)
		// 回溯到上一个状态 根节点
		stack = stack[:len(stack)-1]
	}

	recurFn(root, sum)
	return sumEle
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	return preorder(root, sum, 0, nil)
}

func preorder(root *TreeNode, sum, eleSum int, ele []int) [][]int {
	// 根节点 到 叶子节点 走过的路
	// 走过路的和 == sum 并将这条路进行存储
	// 叶子节点  节点的 左右都为nil则为叶子节点
	// 是叶子节点了，开始存储路径
	if root == nil {
		return nil
	}

	// 从上到下 添加遍历的元素
	ele = append(ele, root.Val)
	// 从上到下 累加 遍历的元素和
	eleSum += root.Val

	// 说明是叶子节点了，开始返回
	if root.Right == nil && root.Left == nil {
		if sum == eleSum {
			tmp := make([]int, len(ele))
			copy(tmp, ele)
			return [][]int{tmp}
		}
		return nil
	}

	var sumEle [][]int
	sumEle = append(sumEle, preorder(root.Left, sum, eleSum, ele)...)
	sumEle = append(sumEle, preorder(root.Right, sum, eleSum, ele)...)
	return sumEle
}
