package main

import "math"

type bTNode struct {
	data        int
	left, right *bTNode
}

func (bt *bTNode) insert(data int) {
	if bt.data == 0 {
		bt.data = data
		return
	}

	if data == bt.data {
		return
	}

	if data < bt.data {
		if bt.left != nil {
			bt.left.insert(data)
			return
		}

		bt.left = &bTNode{
			data:  data,
			left:  nil,
			right: nil,
		}
		return
	}

	if bt.right != nil {
		bt.right.insert(data)
		return
	}

	bt.right = &bTNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (bt *bTNode) delete(data int) *bTNode {
	if bt.data == 0 {
		return nil
	}

	if data < bt.data {
		if bt.left != nil {
			bt.left = bt.left.delete(data)
		}
		return bt
	}

	if data > bt.data {
		if bt.right != nil {
			bt.right = bt.right.delete(data)
		}
		return bt
	}

	if bt.left == nil {
		return bt.right
	}

	if bt.right == nil {
		return bt.left
	}

	smallestNode := bt.right
	for smallestNode.left != nil {
		smallestNode = smallestNode.left
	}
	bt.data = smallestNode.data

	bt.right = smallestNode.delete(smallestNode.data)
	return bt
}

func (bt *bTNode) preOrder() []int {
	visitedNodes := []int{}
	var traverse func(bt *bTNode)

	traverse = func(bt *bTNode) {
		if bt == nil {
			return
		}
		visitedNodes = append(visitedNodes, bt.data)
		traverse(bt.left)
		traverse(bt.right)
	}

	traverse(bt)
	return visitedNodes
}

func (bt *bTNode) postOrder() []int {
	visitedNodes := []int{}
	var traverse func(bt *bTNode)

	traverse = func(bt *bTNode) {
		if bt == nil {
			return
		}
		traverse(bt.left)
		traverse(bt.right)
		visitedNodes = append(visitedNodes, bt.data)
	}

	traverse(bt)
	return visitedNodes
}

func (bt *bTNode) inOrder() []int {
	visitedNodes := []int{}
	var traverse func(bt *bTNode)

	traverse = func(bt *bTNode) {
		if bt == nil {
			return
		}
		traverse(bt.left)
		visitedNodes = append(visitedNodes, bt.data)
		traverse(bt.right)
	}

	traverse(bt)
	return visitedNodes
}

func (bt *bTNode) exist(data int) bool {
	if bt == nil {
		return false
	}

	if data < bt.data {
		return bt.left.exist(data)
	}

	if data > bt.data {
		return bt.right.exist(data)
	}

	return true
}

func (bt *bTNode) height() int {
	if bt == nil {
		return -1
	}

	lH := bt.left.height()
	rH := bt.right.height()

	return 1 + int(math.Max(float64(lH), float64(rH)))
}

func (bt *bTNode) min() int {
	current := bt.left
	for current.left != nil {
		current = current.left
	}

	return current.data
}

func (bt *bTNode) max() int {
	current := bt.right
	for current.right != nil {
		current = current.right
	}

	return current.data
}
