package main

import "fmt"

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//示例:
//
//输入: n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]

func combine(n int, k int) [][]int {
	l := make([][]int, 0)
	s := make([]int, k)

	var dfs func(start, idx int)

	dfs = func(start, idx int) {
		for i := start; i <= n-(k-1-idx); i++ {
			s[idx] = i
			if idx == k-1 {
				l = append(l, append([]int{}, s...))
			} else {
				dfs(i+1, idx+1)
			}
		}
	}
	dfs(1, 0)
	return l
}

func main() {
	fmt.Print(combine(4, 2))
}
