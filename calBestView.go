package main

import "fmt"

//给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。
//
//一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。
//
//返回一对观光景点能取得的最高分。

//计算一次就可以知道 A[i]+i    A[j]-j 两者的最大值
func maxScoreSightseeingPair(A []int) int {
	rmax, imax := 0, A[0]
	for i := 0; i < len(A); i++ {
		imax = max(imax, rmax+A[i]-i)
		rmax = max(rmax, A[i]+i)
	}
	return imax
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{
		8, 1, 5, 2, 6,
	}
	fmt.Println(maxScoreSightseeingPair(arr))
}
