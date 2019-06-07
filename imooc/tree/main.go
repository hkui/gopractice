package main

import (
	"fmt"
	"imooc/tree/tree"
)
//扩展已有类型
type myTreeNode struct {
	node *tree.Node
}

//后序遍历，左右中
func (myNode *myTreeNode)postOrder()  {
	if myNode==nil || myNode.node==nil{
		return
	}
	left:=myTreeNode{myNode.node.Left}
	right:=myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}


func main()  {
	var root tree.Node

	root=tree.Node{Value:3}
	root.Left=&tree.Node{Value:2}
	root.Right=&tree.Node{5,nil,nil}
	root.Right.Left=new(tree.Node) //初始值为0
	root.Right.Right=tree.CreateNode(7)

	fmt.Println("中序遍历:")
	root.Traverse() //23507


	myRoot:=myTreeNode{&root}
	fmt.Println("\n后续遍历:")
	myRoot.postOrder()


	fmt.Println("\n------------------traverse1----------")

	root.Traverse1()
	nodeCount:=0

	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Printf("nodeCount=%d\n",nodeCount)




	/*fmt.Print("\n####################")

	pRoot:=&root
	pRoot.Print()
	fmt.Println()
	pRoot.SetValue(100)
	pRoot.Print()
	fmt.Println("\n---------")

	var pproot *tree.Node
	pproot.SetValue(200)
	pproot=&root
	pproot.SetValue(500)
	pproot.Print()*/



}