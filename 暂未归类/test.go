package main

import "fmt"

func removeDuplicates(nums []int) int {
	left, right, l := 1, 1, len(nums)

	if l < 1 {
		return 0
	}

	for right < l {
		if nums[right-1] != nums[right] {
			nums[left] = nums[right]
			left++
		}
		right++
		fmt.Println(nums)
	}

	nums = nums[:left]
	fmt.Println(nums)
	return left
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	removeDuplicates(nums)
}
