package main

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//示例 1：
//
//输入：s = "()"
//输出：true
//示例 2：
//
//输入：s = "()[]{}"
//输出：true
//示例 3：
//
//输入：s = "(]"
//输出：false
//示例 4：
//
//输入：s = "([)]"
//输出：false
//示例 5：
//
//输入：s = "{[]}"
//输出：true

func isValid(s string) bool {

	l:= len(s)

	if l % 2 == 1 {
		return false
	}

	pairMap := map[byte]byte {
		')' : '(',
		'}' : '{',
		']' : '[',
	}

	stack := []byte{}

	for i := 0 ; i < l ; i++ {
		if pairMap[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack) - 1] != pairMap[s[i]] {
				return false
			}

			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
