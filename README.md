# Simple AVL Tree implement by golang
AVL Tree is the self-balanced tree  

Simple ALV Tree have the insert, delete, find feature
alse have the print feature to print the current tree
and check the tree is balance or not.

## Install
```
go get github.com/Noahnut/AVLTree
```
## Usage
### Insert
```go
avl := CreateAVLTree(comarefunc)
avl.Insert(10)
avl.Insert(20)
avl.Insert(5)
avl.Insert(4)
avl.Insert(30)
avl.Insert(40)
avl.PrintTree()
//
//    10       
//   5   30   
// 4   20 40 
```
### Delete
```go
avl := CreateAVLTree(comarefunc)
avl.Insert(10)
avl.Insert(20)
avl.Insert(5)
avl.Insert(4)
avl.Insert(30)
avl.Insert(40)
avl.Delete(5)
avl.PrintTree()
//
//    10       
//  4   30   
//     20 40
```
### Find
```go
avl := CreateAVLTree(comarefunc)
avl.Insert(10)
avl.Insert(20)
avl.Insert(5)
avl.Insert(4)
//return 4
avl.Find(4)
//return nil
avl.Find(30)
```

### reference
https://josephjsf2.github.io/data/structure/and/algorithm/2019/06/22/avl-tree.html
