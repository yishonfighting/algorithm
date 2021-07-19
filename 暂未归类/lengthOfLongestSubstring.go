package main

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

func lengthOfLongestSubstring(s string) int {
	checkMap := make(map[byte]int, 0)
	ans := 0
	n := len(s)
	r := -1
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(checkMap, s[i-1])
		}
		for r+1 < n && checkMap[s[r+1]] == 0 {
			// 不断地移动右指针
			checkMap[s[r+1]]++
			r++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, r-i+1)
	}

	return ans
}
