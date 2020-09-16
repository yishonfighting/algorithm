package main

import "fmt"

//翻转一棵二叉树。
//
//示例：
//
//输入：
//
//4
///   \
//2     7
/// \   / \
//1   3 6   9
//输出：
//
//4
///   \
//7     2
/// \   / \
//9   6 3   1

type revTree struct {
	Val   int
	Left  *revTree
	Right *revTree
}

func main() {
	t := &revTree{
		Val: 4,
		Left: &revTree{
			Val: 2,
			Left: &revTree{
				Val:   1,
				Left:  &revTree{},
				Right: &revTree{},
			},
			Right: &revTree{
				Val:   3,
				Left:  &revTree{},
				Right: &revTree{},
			},
		},
		Right: &revTree{
			Val: 7,
			Left: &revTree{
				Val:   6,
				Left:  &revTree{},
				Right: &revTree{},
			},
			Right: &revTree{
				Val:   9,
				Left:  &revTree{},
				Right: &revTree{},
			},
		},
	}

	fmt.Print(invertTree(t))
}

/**
 * Definition for a binary tree node.
 */
func invertTree(root *revTree) *revTree {
	if root == nil {
		return root
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Right = left
	root.Left = right

	return root
}
