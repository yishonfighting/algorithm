package main

//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
//请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

import (
	"fmt"
	"strconv"
	"strings"
)

type CodeC struct {
	l []string
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Constructor() CodeC {
	return CodeC{}
}

func (this *CodeC) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

func (this *CodeC) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			this.l = append(this.l, s[i])
		}
	}
	return this.dfs()
}

func (this *CodeC) dfs() *TreeNode {
	if this.l[0] == "#" {
		this.l = this.l[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.l[0])
	root := &TreeNode{Val: val}
	this.l = this.l[1:]
	root.Left = this.dfs()
	root.Right = this.dfs()
	return root
}

func main() {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	obj := Constructor()
	seri := obj.serialize(&root)
	deSeri := obj.deserialize(seri)
	fmt.Println(seri)
	fmt.Println(deSeri)
}
