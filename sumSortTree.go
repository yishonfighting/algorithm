package main

//
//给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
//
//
//
//例如：
//
//输入: 原始二叉搜索树:
//5
///   \
//2     13
//
//输出: 转换为累加树:
//18
///   \
//20     13

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(l *TreeNode)
	dfs = func(l *TreeNode) {
		if l != nil {
			dfs(l.Right)
			sum += l.Val
			l.Val = sum
			dfs(l.Left)
		}
	}

	dfs(root)
	return root
}
