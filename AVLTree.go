package AVL

import (
	"fmt"
	"math"
)

type Compare func(interface{}, interface{}) int

type AVLNode struct {
	value     interface{}
	rightNode *AVLNode
	leftNode  *AVLNode
	height    int
}

type AVLTree struct {
	header     *AVLNode
	compareFun Compare
}

func (avl *AVLTree) maxHeight(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (avl *AVLTree) getBalanceFactor(Node *AVLNode) int {
	if Node == nil {
		return 0
	}

	value := avl.nodeHeight(Node.leftNode) - avl.nodeHeight(Node.rightNode)
	return value
}

func (avl *AVLTree) setNodeHeight(Node *AVLNode) {
	Node.height = avl.getNodeHeight(Node)
}

func (avl *AVLTree) nodeHeight(Node *AVLNode) int {
	if Node == nil {
		return -1
	}
	return Node.height
}

func (avl *AVLTree) getNodeHeight(Node *AVLNode) int {
	return avl.maxHeight(avl.nodeHeight(Node.leftNode), avl.nodeHeight(Node.rightNode)) + 1
}

func (avl *AVLTree) findMaxNode(Node *AVLNode) *AVLNode {
	if Node.leftNode == nil || Node.rightNode == nil {
		return Node
	}
	if avl.compareFun(Node.leftNode.value, Node.rightNode.value) == -1 {
		return avl.findMaxNode(Node.rightNode)
	} else {
		return avl.findMaxNode(Node.leftNode)
	}
}

func (avl *AVLTree) leftRotate(Node *AVLNode) *AVLNode {
	rightChildNode := Node.rightNode
	tempNode := rightChildNode.leftNode

	rightChildNode.leftNode = Node
	Node.rightNode = tempNode
	avl.setNodeHeight(Node)
	avl.setNodeHeight(rightChildNode)
	return rightChildNode
}

func (avl *AVLTree) rightRotate(Node *AVLNode) *AVLNode {
	leftChildNode := Node.leftNode
	tempNode := leftChildNode.rightNode

	leftChildNode.rightNode = Node
	Node.leftNode = tempNode
	avl.setNodeHeight(Node)
	avl.setNodeHeight(leftChildNode)
	return leftChildNode
}

func (avl *AVLTree) balanceAVLTree(Node *AVLNode) *AVLNode {
	blf := avl.getBalanceFactor(Node)
	if blf > 1 {
		if Node.leftNode != nil && Node.leftNode.leftNode != nil {
			// LL type
			return avl.rightRotate(Node)
		} else {
			// LR type
			Node.leftNode = avl.leftRotate(Node.leftNode)
			return avl.rightRotate(Node)
		}
	} else if blf < -1 {
		if Node.rightNode != nil && Node.rightNode.rightNode != nil {
			// RR type
			return avl.leftRotate(Node)
		} else {
			// RL type
			Node.rightNode = avl.rightRotate(Node.rightNode)
			return avl.leftRotate(Node)
		}
	}
	return Node
}

func (avl *AVLTree) insertToAVLTree(data interface{}, currentNode **AVLNode) *AVLNode {
	if *currentNode == nil {
		*currentNode = &AVLNode{value: data, height: 0}
		return *currentNode
	}
	if avl.compareFun(data, (**currentNode).value) == 1 {
		avl.insertToAVLTree(data, &(**currentNode).leftNode)

	} else {
		avl.insertToAVLTree(data, &(**currentNode).rightNode)
	}

	(**currentNode).height = avl.getNodeHeight(*currentNode)
	(*currentNode) = avl.balanceAVLTree((*currentNode))
	return *currentNode
}

func (avl *AVLTree) findFormAVLTree(data interface{}, currentNode *AVLNode) interface{} {
	if currentNode == nil {
		return nil
	}
	if avl.compareFun(data, currentNode.value) == 1 {
		return avl.findFormAVLTree(data, currentNode.leftNode)
	} else if avl.compareFun(data, currentNode.value) == -1 {
		return avl.findFormAVLTree(data, currentNode.rightNode)
	} else {
		return currentNode.value
	}
}

func (avl *AVLTree) deleteFromAVLTree(data interface{}, currentNode *AVLNode) *AVLNode {
	if currentNode == nil {
		return nil
	}
	if avl.compareFun(data, currentNode.value) == 1 {
		currentNode.leftNode = avl.deleteFromAVLTree(data, currentNode.leftNode)
	} else if avl.compareFun(data, currentNode.value) == -1 {
		currentNode.rightNode = avl.deleteFromAVLTree(data, currentNode.rightNode)
	} else {
		// Delete
		if currentNode.rightNode == nil && currentNode.leftNode == nil {
			return nil
		} else if currentNode.leftNode == nil {
			return currentNode.rightNode
		} else if currentNode.rightNode == nil {
			return currentNode.leftNode
		} else {
			maxLeftNode := avl.findMaxNode(currentNode.leftNode)
			currentNode.value = maxLeftNode.value
			currentNode.leftNode = avl.deleteFromAVLTree(currentNode.value, currentNode.leftNode)
		}
	}
	avl.getNodeHeight(currentNode)
	return avl.balanceAVLTree(currentNode)
}

func (avl *AVLTree) printAVLTree(header *AVLNode) {
	treeHeight := header.height
	treeArray := make([]*AVLNode, 0)
	treeArray = append(treeArray, header)
	for i := 0; i <= treeHeight+1; i++ {
		firstSpace := math.Pow(2, float64(treeHeight-i)) - 1
		betweenSpace := math.Pow(2, float64(treeHeight-i)+1) - 1
		for j := int32(firstSpace); j >= 0; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < int(math.Pow(2, float64(i))); k++ {
			if len(treeArray) == 0 {
				break
			}
			node := treeArray[0]
			treeArray = treeArray[1:]
			if node.value == nil {
				fmt.Print(" ")
			} else {
				fmt.Print(node.value)
			}

			for l := betweenSpace; l > 0; l-- {
				fmt.Printf(" ")
			}
			leftNode := node.leftNode
			rightNode := node.rightNode
			if leftNode == nil {
				leftNode = &AVLNode{}
			}

			if rightNode == nil {
				rightNode = &AVLNode{}
			}

			treeArray = append(treeArray, leftNode)
			treeArray = append(treeArray, rightNode)
		}
		fmt.Print("\n")
	}
}

func (avl *AVLTree) checkAVLTreeBalance(currentNode *AVLNode) bool {
	balanceFactor := avl.getBalanceFactor(currentNode)
	if balanceFactor > 1 || balanceFactor < -1 {
		return false
	}

	if currentNode.leftNode != nil {
		avl.checkAVLTreeBalance(currentNode.leftNode)
	}

	if currentNode.rightNode != nil {
		avl.checkAVLTreeBalance(currentNode.rightNode)
	}

	return true
}

func (avl *AVLTree) PrintTree() {
	avl.printAVLTree(avl.header)
}

func (avl *AVLTree) Insert(data interface{}) {
	avl.header = avl.insertToAVLTree(data, &avl.header)
}

func (avl *AVLTree) Delete(data interface{}) {
	avl.header = avl.deleteFromAVLTree(data, avl.header)
}

func (avl *AVLTree) Find(data interface{}) interface{} {
	return avl.findFormAVLTree(data, avl.header)
}

func (avl *AVLTree) CheckBalance() bool {
	return avl.checkAVLTreeBalance(avl.header)
}

func CreateAVLTree(compareFun Compare) *AVLTree {
	avl := AVLTree{header: nil, compareFun: compareFun}
	return &avl
}
