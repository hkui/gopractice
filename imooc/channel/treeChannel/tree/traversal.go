package tree

import "fmt"

func (node *Node)Traverse1()  {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node)TraverseFunc(f func(*Node))  {
	if node ==nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node)TraverseWithChannel()chan *Node {
	out:=make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			fmt.Printf("send %d\n",node.Value)
			out<-node
		})
		close(out)
	}()
	return out
}