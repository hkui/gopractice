package main

import (
	"fmt"
	"imooc/channel/treeChannel/tree"
)

func main() {

	var root tree.Node

	root=tree.Node{Value:3}
	root.Left=&tree.Node{Value:2}
	root.Right=&tree.Node{5,nil,nil}
	root.Right.Left=new(tree.Node) //初始值为0
	root.Right.Right=tree.CreateNode(7)
	fmt.Println("中序遍历:")
	root.Traverse()
	fmt.Println()
	nodeCount:=0

	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Printf("nodeCount=%d\n",nodeCount)


	c:=root.TraverseWithChannel()
	maxNode:=0
	for node:=range c{
		fmt.Printf("receive %d\n",node.Value)
		if(node.Value>maxNode){
			maxNode=node.Value
		}
	}
	fmt.Printf("max node=%d\n",maxNode)
}
