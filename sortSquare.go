package main

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//示例 2：
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]

func sortedSquares(nums []int) []int {
	lv, rv := 0, len(nums)-1
	i := rv + 1
	ul := make([]int, i)

	for lv <= rv {
		if nums[lv] < 0 {
			nums[lv] = 0 - nums[lv]
		}

		if nums[lv] > nums[rv] {
			ul[i] = nums[lv] * nums[lv]
			lv++
		} else {
			ul[i] = nums[rv] * nums[rv]
			rv--
		}
		i--
	}

	return ul
}
