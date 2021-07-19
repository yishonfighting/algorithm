package main

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	ans, l := 0, len(nums)
	for i := 0; i < l; i++ {
		if i == 0 {
			for j := 0; j < k; j++ {
				if nums[j] > ans {
					ans = nums[j]
				}
			}
			res = append(res, ans)
		} else {
			if i+k > l {
				return res
			}

			if nums[i+k-1] > ans && nums[]{
				ans = nums[i+k-1]
			}

			res = append(res, ans)
		}

	}
	return res
}
