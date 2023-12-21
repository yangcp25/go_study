package main

func main() {

	// 双向链表
	//
	/**
	先理解查找过程
	Level 3: 1		 6
	Level 2: 1   3   6
	Level 1: 1 2 3 4 6

	比如 查找2 ; 从高层往下找;
	如果查找的值比当前值小 说明没有可查找的值
	2比1大 往当前层的下个节点查找，3层的后面没有了或者比后面的6小 ，往下层找
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
		找到插入的位置
		确定他当前的层数
		在他的层数连接当前节点

	    如何确定层数？
			来一个概率的算法就行
			这样在数量大的时候能基本能达到2分查找的效果（概率是1/2）

	    更新索引数组？
		我们在查找的时候的路径就可以拿来做插入的数据
	    比如查找4
		找的路径是 3层的 1，2层的3 ；
		如果4是第三层的
		更新3层 1->4>6
		更新2层 1->3->4->6
	*/

	/**
	删除流程 基本同上

	*/

	/**

	 */

}

// MAX_LEVEL 最高层数
const MAX_LEVEL = 16

type T any

var _ skipListHandle[T] = &skipList[T]{}

type skipListHandle[T any] interface {
	insert(data T, score uint32) bool
	delete(data T) bool
	headForeach()
}

type skipListNode[T any] struct {
	data T
	// 上一个节点 用于遍历
	prev *skipListNode[T]
	// 分数
	score uint32
	// 下个节点 同时也是索引
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
func (list skipList[T]) insert(data T, score uint32) bool {
	return false
}

func (list skipList[T]) delete(data T) bool {
	//TODO implement me
	panic("implement me")
}

func (list skipList[T]) headForeach() {
	headNode := list.head

	for headNode.forwards != nil {

	}
}
