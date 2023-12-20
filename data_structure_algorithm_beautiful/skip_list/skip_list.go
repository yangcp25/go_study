package main

func main() {

	// 双向链表
	//
	/**
	先理解查找过程
	Level 3: 1
	Level 2: 1   3   6
	Level 1: 1 2 3 4 6

	比如 查找2 ; 从高层往下找;
	如果查找的值比当前值小 说明没有可查找的值
	2比1大 往当前层的下个节点查找，3层的后面没有了 ，往下层找
	2层 查找值比下个节点3还小 往下层找
	最后一层找到

	比如查找 4
	没有找到 3层往下到2层; 2层里 4比3大继续往前，比6小，往下层找
	从第一层的继续往前找

	比如查找 5
	第一层的3开始往前找到6比查找值5大，说明没有待查找值
	*/

	/**
	插入流程
	*/

}

// MAX_LEVEL 最高层数
const MAX_LEVEL = 16

type T any

var _ skipListHandle[T] = &skipList[T]{}

type skipListHandle[T any] interface {
	insert(data T) bool
	delete(data T) bool
	foreach()
}

type skipListNode[T any] struct {
	data T
	prev *skipListNode[T]

	forwards []*skipListNode[T]
}

type skipList[T any] struct {
	head, tail *skipListNode[T]
	// 跳表高度
	level uint32
	// 跳表长度
	length uint32
}

func createSkipList[T any]() *skipList[T] {
	return &skipList[T]{
		level:  0,
		length: 1,
	}
}
func (list skipList[T]) insert(data T) bool {
	return false
}

func (list skipList[T]) delete(data T) bool {
	//TODO implement me
	panic("implement me")
}

func (list skipList[T]) foreach() {
	//TODO implement me
	panic("implement me")
}
