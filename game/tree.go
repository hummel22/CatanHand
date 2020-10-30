package game

import "Catan/objects"

type Node struct {
	Parent *Node
	Child *Node
	Good objects.Good
	Score uint64
	ChildScore uint64
}


func (n *Node)Collapse() []objects.Good {
	goods := make([]objects.Good ,0,10)
	node := n.Child
	for node != nil {
		goods = append(goods, node.Good)
		node = node.Child
	}
	return goods
}

func (n *Node)Update(scoreSum uint64, child *Node) {
	if n == nil {
		return
	}
	if scoreSum > n.ChildScore {
		n.ChildScore = scoreSum
		n.Child = child
		n.Parent.Update(n.ChildScore + n.Score,n)
	}
}


func Graph(resources objects.Resources) []objects.Good {
	root := &Node{}
	Spend(root, resources)
	return root.Collapse()
}

func Spend(parent *Node, resources objects.Resources) {
	for good, cost := range objects.CostsKey{
		if left, ok := resources.Buy(cost); ok {
			score := objects.Values[good]
			node := &Node{
				Parent: parent,
				Good: good,
				Score: score ,
				ChildScore: 0,
			}
			parent.Update(score, node)
			Spend(node, left)
		}
	}
}


