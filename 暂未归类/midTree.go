package main

import "fmt"

//给定一个二叉树，返回它的中序 遍历。
//
//示例:
//
//输入: [1,null,2,3]
//1
//\
//2
///
//3
//
//输出: [1,3,2]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？

type treeNode struct {
	Val       int
	LeftNode  *treeNode
	RightNode *treeNode
}

func main() {
	tree := &treeNode{
		Val:      1,
		LeftNode: &treeNode{},
		RightNode: &treeNode{
			Val:       2,
			LeftNode:  &treeNode{Val: 3, LeftNode: &treeNode{}, RightNode: &treeNode{}},
			RightNode: &treeNode{},
		},
	}

	fmt.Print(inorderTraversal(tree))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *treeNode) []int {
	var dfs func(n *treeNode)
	var l []int

	dfs = func(n *treeNode) {
		if n == nil {
			return
		}
		dfs(n.LeftNode)
		l = append(l, n.Val)
		dfs(n.RightNode)
	}
	dfs(root)
	return l
}
