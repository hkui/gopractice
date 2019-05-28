package main

import (
	"fmt"
	"imooc/tree/tree"
)


func main()  {
	var root tree.TreeNode

	root=tree.TreeNode{Value:3}
	root.Left=&tree.TreeNode{Value:2}
	root.Right=&tree.TreeNode{5,nil,nil}
	root.Right.Left=new(tree.TreeNode) //初始值为0
	root.Right.Right=tree.CreateTreeNode(7)

	root.Traverse() //23507
	fmt.Print("\n####################")

	pRoot:=&root
	pRoot.Print()
	fmt.Println()
	pRoot.SetValue(100)
	pRoot.Print()
	fmt.Println("\n---------")

	var pproot *tree.TreeNode
	pproot.SetValue(200)
	pproot=&root
	pproot.SetValue(500)
	pproot.Print()

}