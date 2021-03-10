package main

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]

func moveZeroes(nums []int) {
	left, right, l := 0, 0, len(nums)

	for right < l {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
