package main

import "fmt"

/**
*@Author:徐鹏豪
*@Time: 2021/2/16 10:14
*@Description:
 */
func main() {
	// nums := []int{1, 3, 2, 6, 5}
	nums := []int{1, 6, 3, 2, 5}
	fmt.Println(verifyPostorder2(nums))
}

// 剑指 Offer 33. 【二叉搜索树】的后序遍历序列
// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
/*
     5
    / \
   2   6
  / \
 1   3
 [1,3,2,6,5] true
[1,6,3,2,5] false
*/
func verifyPostorder2(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, startIndex, endIndex int) bool {
	if startIndex >= endIndex {
		return true
	}
	// 左边应该都比 根小
	p := startIndex
	root := postorder[endIndex]
	for postorder[p] < root {
		p++
	}
	m := p // 右边的开始位置
	for postorder[p] > root {
		p++
	}
	// 正常情况下 p == j 应该是true的
	// p == endIndex (可能存在 右边部分 小于root的情况) 左边比较，右边比较
	return p == endIndex && recur(postorder, startIndex, m-1) && recur(postorder, m, endIndex-1)
}

func verifyPostorder(postorder []int) bool {
	// 切片中最后一个是根， 如果根只有一个 或者 为nil 则直接返回true
	if len(postorder) <= 1 {
		return true
	}
	// 获取根节点 即末尾最后一个
	root := postorder[len(postorder)-1]
	var leftNums, rightNums []int
	for i := 0; i < len(postorder)-1; i++ {
		// 找出左边End
		// root 没有左子树  7 ，9，10 ，5
		// root 没有右树  1，2，3，5
		// 非正常情况 9,10,2,11,8

		// 1,3,2,6,5 true
		// 1,6,3,2,5 false
		// 一定是现有左，才有右的
		if postorder[i] < root && len(rightNums) == 0 {
			leftNums = append(leftNums, postorder[i])
		} else {
			// 如果 右边的节点 有 小于根据节点 都直接 返回 false
			if postorder[i] < root {
				return false
			}
			rightNums = append(rightNums, postorder[i])
		}
	}
	// 找出左边，-》 比根小的是左边
	// 找出右边-》比根大的是右边
	return verifyPostorder(leftNums) && verifyPostorder(rightNums)
}
