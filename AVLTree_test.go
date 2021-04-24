package AVL

import (
	"testing"
)

func comarefunc(A interface{}, B interface{}) int {
	if A.(int) < B.(int) {
		return 1
	} else if A.(int) > B.(int) {
		return -1
	} else {
		return 0
	}
}

func TestBasicAVLinsert(t *testing.T) {
	//RR type left rotate test
	avl := CreateAVLTree(comarefunc)
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(5)
	avl.Insert(4)
	avl.Insert(30)
	avl.Insert(40)
	if avl.getNodeHeight(avl.header) != 2 {
		t.Errorf("AVL Tree is not balance root tree height not equal %d", 2)
	}

	if avl.header.rightNode.value != 30 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.rightNode.value)
	}

	if avl.header.rightNode.rightNode.value != 40 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.rightNode.rightNode.value)
	}

	if avl.header.rightNode.leftNode.value != 20 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.rightNode.leftNode.value)
	}

	//RL type test
	avl = CreateAVLTree(comarefunc)
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(5)
	avl.Insert(4)
	avl.Insert(25)
	avl.Insert(21)
	if avl.getNodeHeight(avl.header) != 2 {
		t.Errorf("AVL Tree is not balance root tree height not equal %d", 2)
	}

	if avl.header.rightNode.value != 21 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.rightNode.value)
	}

	if avl.header.rightNode.rightNode.value != 25 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.rightNode.rightNode.value)
	}

	if avl.header.rightNode.leftNode.value != 20 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.rightNode.leftNode.value)
	}

	//LL type test
	avl = CreateAVLTree(comarefunc)
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(5)
	avl.Insert(4)
	avl.Insert(3)
	avl.Insert(15)
	if avl.getNodeHeight(avl.header) != 2 {
		t.Errorf("AVL Tree is not balance root tree height not equal %d", 2)
	}

	if avl.header.leftNode.value != 4 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.leftNode.value)
	}

	if avl.header.leftNode.rightNode.value != 5 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.leftNode.rightNode.value)
	}

	if avl.header.leftNode.leftNode.value != 3 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.leftNode.leftNode.value)
	}

	//LR type test
	avl = CreateAVLTree(comarefunc)
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(6)
	avl.Insert(4)
	avl.Insert(5)
	avl.Insert(15)
	if avl.getNodeHeight(avl.header) != 2 {
		t.Errorf("AVL Tree is not balance root tree height not equal %d", 2)
	}

	if avl.header.leftNode.value != 5 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.leftNode.value)
	}

	if avl.header.leftNode.rightNode.value != 6 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.leftNode.rightNode.value)
	}

	if avl.header.leftNode.leftNode.value != 4 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.leftNode.leftNode.value)
	}

	avl = CreateAVLTree(comarefunc)
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(30)

	if avl.getNodeHeight(avl.header) != 1 {
		t.Errorf("AVL Tree is not balance root tree height not equal %d", 1)
	}

	if avl.header.value != 20 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.value)
	}

	if avl.header.leftNode.value != 10 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.leftNode.value)
	}

	if avl.header.rightNode.value != 30 {
		t.Errorf("AVL Tree is not rotate correct because expect value not %d", avl.header.rightNode.value)
	}
}

type testStruct struct {
	key   int
	value int
}

func FindCompareFunc(A interface{}, B interface{}) int {
	if A.(testStruct).key < B.(testStruct).key {
		return 1
	} else if A.(testStruct).key > B.(testStruct).key {
		return -1
	} else {
		return 0
	}
}

func TestBasicAVLfind(t *testing.T) {
	avl := CreateAVLTree(FindCompareFunc)
	testArray := []testStruct{}
	for i := 0; i < 10; i++ {
		t := testStruct{key: i * 10, value: i * 10}
		testArray = append(testArray, t)
	}

	for i := 0; i < 10; i++ {
		t := testStruct{key: i, value: i}
		testArray = append(testArray, t)
	}

	for _, e := range testArray {
		avl.Insert(e)
	}

	findStruct := testStruct{key: 90}

	result := avl.Find(findStruct).(testStruct)

	if result.value != 90 {
		t.Errorf("Find is not equal expect value %d", 90)
	}

	findStruct = testStruct{key: 70}

	result = avl.Find(findStruct).(testStruct)

	if result.value != 70 {
		t.Errorf("Find is not equal expect value %d", 70)
	}

	findStruct = testStruct{key: 7}
	result = avl.Find(findStruct).(testStruct)

	if result.value != 7 {
		t.Errorf("Find is not equal expect value %d", result.value)
	}

}

func TestBasicDelete(t *testing.T) {

	//Delete right node is empty
	avl := CreateAVLTree(comarefunc)
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(5)
	avl.Insert(4)
	avl.Insert(30)
	avl.Insert(40)
	avl.Delete(5)

	if avl.Find(5) != nil {
		t.Error("Fail to delete 5")
	}

	avl = CreateAVLTree(comarefunc)
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(5)
	avl.Insert(6)

	avl.Insert(30)
	avl.Insert(40)
	avl.Delete(5)
	avl.PrintTree()

	if avl.Find(5) != nil {
		t.Error("Fail to delete 5")
	}

	avl = CreateAVLTree(comarefunc)
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(5)
	avl.Insert(4)
	avl.Insert(3)
	avl.Insert(30)
	avl.Insert(40)
	avl.Insert(15)
	avl.Insert(17)
	avl.Insert(70)
	avl.Insert(45)
	avl.Insert(22)
	avl.Insert(24)
	avl.Insert(2)
	avl.Insert(25)
	avl.Insert(1)
	avl.Delete(10)
	avl.Delete(17)
	avl.PrintTree()

	if avl.CheckBalance() != true {
		t.Error("Tree is not Balance")
	}

	if avl.Find(10) != nil {
		t.Error("Fail to delete 10")
	}

	if avl.Find(17) != nil {
		t.Error("Fail to delete 17")
	}

}
