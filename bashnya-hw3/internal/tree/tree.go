package tree

import (
	"fmt"
	"strings"
)

type Node struct {
	value     int
	leftNode  *Node
	rightNode *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Find(value int) *Node {
	if t.root == nil {
		return nil
	}

	node := t.root
	for node != nil {
		switch {
		case value < node.value:
			node = node.leftNode
		case value > node.value:
			node = node.rightNode
		default:
			return node
		}
	}
	return nil
}

func (t *Tree) Depth() int {
	if t.root == nil {
		return 0
	}
	type stackItem struct {
		node  *Node
		depth int
	}
	max_depth := 0
	stack := make([]*stackItem, 0, 16)
	stack = append(stack, &stackItem{node: t.root, depth: 1})

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currentNode := item.node
		currentDepth := item.depth
		if currentNode.leftNode != nil {
			stack = append(stack, &stackItem{node: currentNode.leftNode, depth: currentDepth + 1})
		}
		if currentNode.rightNode != nil {
			stack = append(stack, &stackItem{node: currentNode.rightNode, depth: currentDepth + 1})

		}
		if max_depth < currentDepth {
			max_depth = currentDepth
		}
	}
	return max_depth
}

func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
		return
	}

	node := t.root

	for {
		if value < node.value {
			if node.leftNode == nil {
				node.leftNode = &Node{value: value}
				return
			}
			node = node.leftNode

		} else if value > node.value {

			if node.rightNode == nil {
				node.rightNode = &Node{value: value}
				return
			}
			node = node.rightNode

		} else {
			return
		}
	}
}

func (t *Tree) Remove(value int) {
	if t.root == nil {
		return
	}

	if t.root.value == value {
		t.root = nil
		return
	}
	var parent *Node
	var removedNode *Node
	node := t.root
searchTree:
	for node != nil {
		switch {
		case value < node.value:
			parent = node
			node = node.leftNode
		case value > node.value:
			parent = node
			node = node.rightNode
		default:
			break searchTree

		}
	}
	if node == nil {
		return
	}
	removedNode = node
	switch {
	case removedNode.leftNode != nil && removedNode.rightNode != nil:
		parent = removedNode
		node = removedNode.rightNode
		for node.leftNode != nil {
			parent = node
			node = node.leftNode
		}
		if parent.leftNode == node {
			parent.leftNode = node.rightNode
		} else {
			parent.rightNode = node.rightNode
		}
		removedNode.value = node.value
	case removedNode.leftNode == nil && removedNode.rightNode == nil:
		if parent.leftNode.value == value {
			parent.leftNode = nil
		} else {
			if parent.rightNode.value == value {
				parent.rightNode = nil
			}
		}
	default:
		var child *Node
		if removedNode.leftNode != nil {
			child = removedNode.leftNode
		} else {
			child = removedNode.rightNode
		}

		if parent.leftNode == removedNode {
			parent.leftNode = child
		} else {
			if parent.rightNode == removedNode {
				parent.rightNode = child
			}
		}

	}

}

func (t *Tree) Print() {
	if t.root == nil {
		fmt.Println("Дерево пустое")
		return
	}
	fmt.Println("Структура дерева:")
	t.printNodeHorizontal(t.root, 0)
}

func (t *Tree) printNodeHorizontal(node *Node, level int) {
	if node == nil {
		return
	}

	t.printNodeHorizontal(node.rightNode, level+1)

	indent := strings.Repeat("    ", level)
	fmt.Printf("%s%3d\n", indent, node.value)

	t.printNodeHorizontal(node.leftNode, level+1)
}
