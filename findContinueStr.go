package main

//输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

//示例 1：
//
//输入：target = 9
//输出：[[2,3,4],[4,5]]
//示例 2：
//
//输入：target = 15
//输出：[[1,2,3,4,5],[4,5,6],[7,8]]

func findContinuousSequence(target int) [][]int {
	lv, rv := 1, 2
	var res [][]int

	for lv < rv {
		sum := (lv + rv) * (rv - lv + 1) / 2

		if sum == target {
			s := make([]int, 0)
			for i := lv; i <= rv; i++ {
				s = append(s, i)
			}
			res = append(res, s)
			lv++
		} else {
			if sum > target {
				lv++
			} else {
				rv++
			}
		}
	}

	return res
}
