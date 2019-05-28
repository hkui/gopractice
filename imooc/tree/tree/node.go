package tree

import "fmt"

type TreeNode struct {
	Value int
	Left,Right *TreeNode
}
func CreateTreeNode(Value int) *TreeNode{
	return &TreeNode{Value:Value}  //返回了局部变量的地址
}
func (node TreeNode)Print()  {
	fmt.Print(node.Value)
}
func (node *TreeNode)SetValue(Value int)  {
	if(node ==nil){
		fmt.Println("setting Value=",Value," to nil")
		return
	}
	node.Value=Value
}
//中序遍历(左中右)
func (node *TreeNode)Traverse()  {
	if(node ==nil){
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

