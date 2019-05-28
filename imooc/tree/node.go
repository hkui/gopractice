package main

import "fmt"

type treeNode struct {
	value int
	left,right *treeNode
}
func createTreeNode(value int) *treeNode{
	return &treeNode{value:value}  //返回了局部变量的地址
}
func (node treeNode)print()  {
	fmt.Print(node.value)
}
func (node *treeNode)setValue(value int)  {
	if(node ==nil){
		fmt.Println("setting value=",value," to nil")
		return
	}
	node.value=value
}
//中序遍历(左中右)
func (node *treeNode)traverse()  {
	if(node ==nil){
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}


func main()  {
	var root treeNode

	root=treeNode{value:3}
	root.left=&treeNode{value:2}
	root.right=&treeNode{5,nil,nil}
	root.right.left=new(treeNode) //初始值为0
	root.right.right=createTreeNode(7)

	root.traverse() //23507
	fmt.Print("\n####################")

	pRoot:=&root
	pRoot.print()
	fmt.Println()
	pRoot.setValue(100)
	pRoot.print()
	fmt.Println("\n---------")

	var pproot *treeNode
	pproot.setValue(200)
	pproot=&root
	pproot.setValue(500)
	pproot.print()

}