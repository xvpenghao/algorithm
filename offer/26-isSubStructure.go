package main

import (
	"container/list"
	"fmt"
)

func main() {
	rootA := &TreeNode2{
		Val: 3,
		Left: &TreeNode2{
			Val: 4,
			Left: &TreeNode2{
				Val: 1,
			},
			Right: &TreeNode2{
				Val: 2,
			},
		},
		Right: &TreeNode2{
			Val: 5,
		},
	}

	rootB := &TreeNode2{
		Val: 4,
		Left: &TreeNode2{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	fmt.Println(isSubStructure2(rootA, rootB))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

// 递归的实现
// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
// 剑指 Offer 26. 树的子结构
func isSubStructure2(A *TreeNode2, B *TreeNode2) bool {
	// 前序遍历 + 判断是否有交集
	// 空树 不是任意一个树的子结构
	if B == nil || A == nil {
		return false
	}

	// 比较两者 是否 有交集
	// 这里类似 前序的输出
	if recur(A, B) {
		return true
	}

	return isSubStructure2(A.Left, B) || isSubStructure2(A.Right, B)
}

// 3 4 5 1 2
// 4 1
func isSubStructure(A *TreeNode2, B *TreeNode2) bool {
	// 空树 不是任意一个树的子结构
	if B == nil {
		return false
	}

	var list2 list.List
	// 要不然第一次不会进去
	for list2.Len() > 0 || A != nil {
		// 左边 遍历完毕了
		for A != nil {
			// 递归比较
			if recur(A, B) {
				return true
			}
			list2.PushBack(A)
			A = A.Left
		}

		// 元素的 弹出
		A = list2.Remove(list2.Back()).(*TreeNode2)
		// 设置右边的遍历节点
		A = A.Right
	}
	return false
}

func recur(A *TreeNode2, B *TreeNode2) bool {
	// 说明已经遍历完毕了 递归结束条件
	if B == nil {
		return true
	}

	if A == nil || (A.Val != B.Val) {
		return false
	}
	// A和B 都相等 才相等
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

/*
	 3
	4 5
	1 2
*/
func traverse(A *TreeNode2) {
	var list2 list.List
	// 要不然第一次不会进去
	for list2.Len() > 0 || A != nil {
		// 左边 遍历完毕了
		for A != nil {
			fmt.Println(A.Val)

			list2.PushBack(A)
			A = A.Left
		}

		// 元素的 弹出
		A = list2.Remove(list2.Back()).(*TreeNode2)
		// 设置右边的遍历节点
		A = A.Right
	}
}
