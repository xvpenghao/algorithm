package main

import "fmt"

// 剑指 Offer 19. 正则表达式匹配
// https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/
// 修订日期
// 2021年 1月17日 星期日 16时20分01秒 CST
func main() {
	// aaa ab*a*c*a
	fmt.Println(isMatch2("aab", "ab*a*c*a"))
}

/*
解题思路
B：是要匹配的字符
A：是字符
1、如果 B 的最后一个字符是正常字符，那就是看 A[n-1] 是否等于B[m−1]，相等则看 A0..n−2与 B0..m−2，不等则是不能匹配，这就是子问题。
A:abc
B:abc

2、如果 B的最后一个字符是.，它能匹配任意字符，直接看 A0..n−2与 B0..m−2
​A:abc
B:ab.


3、如果 B 的最后一个字符是*它代表 B[m−2]=c 可以重复0次或多次，它们是一个整体 c∗
A:abb
B:ab*

情况一：A[n−1] 是 0 个 c，B 最后两个字符废了，能否匹配取决于 A0..n−1 和 B0..m−3是否匹配
A:abd
B:abc*

情况二：A[n−1] 是多个 c 中的最后一个（这种情况必须 A[n−1]=c 或者 c='.'），所以 A 匹配完往前挪一个，B 继续匹配，
因为可以匹配多个，继续看 A0..n−2和 B0..m−1是否匹配
????

转义方程

*/
// todo 看的我自闭了
func isMatch2(A string, B string) bool {
	n, m := len(A), len(B)
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}

	// i=0 代表是 空字符
	// j=0 代表是 空正则
	// n*m的比较
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// 分成空正则和非空正则
			if j == 0 {
				// j为 0 则代表是 空正则
				// 如果 i是0(代表空串) ，则算直接匹配成功
				// 如果 i是1(代表不是空串)  非空串和 空正则肯定是不能匹配成功的
				f[i][j] = i == 0
			} else {
				// 这里可能会出现 空字符 非空正则 || 非空字符 非空正则这两种case
				// 非空正则分为 两种情况 * 和 非*
				// j==1时候 j-1 刚好取的就是 B的第一个
				if B[j-1] != '*' {
					// 两个 A=abc B=abc a相等 || A=abc B=.abc  因为B的 首字母是.
					if i > 0 && (A[i-1] == B[j-1] || B[j-1] == '.') {
						f[i][j] = f[i-1][j-1]
					}
				} else {
					// 碰到了 *了，分为看 和不看两种情况
					// 不看
					// aab ab*a*c*a
					// aba ab* 0个c，B的最后两个字符可作废了
					if j >= 2 {
						// 取前一个状态 然后赋值给后一个状态 f[1][3]， A的前1个字符是否和B的前3个字符相等，==不看== f[1][1] 取A的前1字符是否和B的前1个字符相等
						f[i][j] = f[i][j-2]
					}
					// 看 abcd abcd*
					if i >= 1 && j >= 2 && (A[i-1] == B[j-2] || B[j-2] == '.') {
						// f[i][j] 可能取值是 上一个的 比如 f[1][5]的时候 ，思考 f[i-1][j]的作用
						// 为什么要 f[i-1][j]的这步骤操作
						f[i][j] = f[i][j] || f[i-1][j]
					}
				}
			}
		}
	}
	// 后一个状态 依赖前一个状态
	// 最后的 f[3][8] A的前3个是否和B的前8个匹配
	return f[n][m]
}

// no bad
func isMatch(s string, p string) bool {
	m, n := len(s)+1, len(p)+1
	dp := make([][]bool, m)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}
	// 代表的是 代表两个空字符串能够匹配，所以 dp[i][j]对应的添加字符是 s[i-1]和 p[j-1]
	dp[0][0] = true
	// 初始首位 *肯定可定不会出现在首位，比如 *abc是个错误的举例
	for j := 2; j < n; j += 2 {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	// 状态转移
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[j-1] == '*' { // 当匹配的字符是* 的时候
				if dp[i][j-2] { // 1
					dp[i][j] = true
				} else if dp[i][j-1] { // 2
					dp[i][j] = true
				} else if dp[i-1][j] && s[i-1] == p[j-2] { // 3
					dp[i][j] = true
				} else if dp[i-1][j] && p[j-2] == '.' { // 4
					dp[i][j] = true
				}
			} else {
				if dp[i-1][j-1] && s[i-1] == p[j-1] { // 1
					dp[i][j] = true
				} else if dp[i-1][j-1] && p[j-1] == '.' { // 2
					dp[i][j] = true
				}
			}
		}
	}
	return dp[m-1][n-1]
}
