package main

//给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。
//
//如果不存在满足条件的子数组，则返回 0 。

//示例 1：
//
//输入：nums = [8,2,4,7], limit = 4
//输出：2
//解释：所有子数组如下：
//[8] 最大绝对差 |8-8| = 0 <= 4.
//[8,2] 最大绝对差 |8-2| = 6 > 4.
//[8,2,4] 最大绝对差 |8-2| = 6 > 4.
//[8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
//[2] 最大绝对差 |2-2| = 0 <= 4.
//[2,4] 最大绝对差 |2-4| = 2 <= 4.
//[2,4,7] 最大绝对差 |2-7| = 5 > 4.
//[4] 最大绝对差 |4-4| = 0 <= 4.
//[4,7] 最大绝对差 |4-7| = 3 <= 4.
//[7] 最大绝对差 |7-7| = 0 <= 4.
//因此，满足题意的最长子数组的长度为 2 。


//示例 2：
//
//输入：nums = [10,1,2,4,7,2], limit = 5
//输出：4
//解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。

//示例 3：
//
//输入：nums = [4,2,2,2,4,4,2,2], limit = 0
//输出：3


func longestSubarray(nums []int, limit int) (ans int) {
	var minQ , maxQ []int

	l := 0
	for r, v := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ,v)

		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ,v)

		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0] - minQ[0] > limit {
			if nums[l] == maxQ[0] {
				maxQ = maxQ[1:]
			}

			if nums[l] == minQ[0] {
				minQ = minQ[1:]
			}

			l++
		}
		ans = max(ans, r - l +1)

	}


	return
}

func max(a ,b int) int {
	if a > b {
		return a
	}
	return b
}