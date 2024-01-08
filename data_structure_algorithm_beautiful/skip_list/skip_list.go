package main

import (
	"errors"
	"math/rand"
)

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

type T comparable

type skipListHandle[T comparable] interface {
	insert(data T, score uint32) (err error)
	delete(data T, score uint32) int
	findNode(data T, score uint32) (error, *skipListNode[T])
}

type skipListNode[T comparable] struct {
	data T
	// 上一个节点 用于遍历
	prev *skipListNode[T]
	// 排序分数
	score uint32
	// 下个节点 同时也是索引
	forwards []*skipListNode[T]
}

type skipList[T comparable] struct {
	head, tail *skipListNode[T]
	// 跳表高度
	level int
	// 跳表长度
	length uint32
}

func createSkipList[T comparable]() *skipList[T] {
	return &skipList[T]{
		level:  1,
		length: 0,
	}
}

func createNode[T comparable](data T, score uint32) *skipListNode[T] {
	return &skipListNode[T]{
		data:     data,
		prev:     nil,
		score:    score,
		forwards: make([]*skipListNode[T], 0, MAX_LEVEL),
	}
}
func (list skipList[T]) insert(data T, score uint32) error {
	currenNode := list.head
	maxIndex := MAX_LEVEL - 1
	// 找到插入的位置
	// 记录插入的路径 记录第一个比待查找的值小的位置
	path := [MAX_LEVEL]*skipListNode[T]{}
	for i := list.level - 1; i >= 0; i++ {
		for currenNode.forwards[i] != nil {
			// 如果插入的位置比当前数据小 直接跳出循环并且高度下降
			if currenNode.forwards[i].score > score {
				path[i] = currenNode
				break
			}
			// 插入位置比当前的大，在当前层继续往前找
			currenNode = currenNode.forwards[i]
		}
		// 如果currenNode.forwards[i] == nil 说明是最后一个值了 所以直接插入
		if currenNode.forwards[i] == nil {
			path[i] = currenNode
		}
	}

	// 随机算法求得最大层数
	level := 1

	for i := 1; i < maxIndex; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	newNode := createNode(data, score)

	// 原有节点连接
	for i := 0; i < maxIndex; i++ {
		next := path[i].forwards[i]
		// path[i]拿到第一个插入值小的位置 forwards[i] 是指在当前层它指向的下个节点
		newNode.forwards[i] = next
		path[i].forwards[i] = newNode
	}

	// 更新level
	if level > list.level {
		list.level = level
	}

	list.length++

	return errors.New("插入失败")
}

func (list skipList[T]) delete(data T, score uint32) int {
	currenNode := list.head
	// 找到插入的位置
	// 记录插入的路径 记录第一个比待查找的值小的位置
	path := [MAX_LEVEL]*skipListNode[T]{}
	for i := list.level - 1; i >= 0; i++ {
		path[i] = list.head
		for currenNode.forwards[i] != nil {
			// 記錄刪除的位置
			if currenNode.forwards[i].score == score && currenNode.forwards[i].data == data {
				path[i] = currenNode
				break
			}
			// 插入位置比当前的大，在当前层继续往前找
			currenNode = currenNode.forwards[i]
		}
	}
	currenNode = path[0].forwards[0]
	for i := list.level - 1; i >= 0; i-- {
		if path[i] == list.head && currenNode.forwards[i] == nil {
			list.level = i
		}

		if nil == path[i].forwards[i] {
			path[i].forwards[i] = nil
		} else {
			path[i].forwards[i] = path[i].forwards[i].forwards[i]
		}
	}

	list.length--

	return 0
}

func (list skipList[T]) findNode(v T, score uint32) (err error, node *skipListNode[T]) {
	if nil == v || list.length == 0 {
		return errors.New("请传入查找的值"), node
	}

	cur := list.head
	for i := list.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].data == v {
				return nil, cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}
	return errors.New("请传入查找的值"), nil
}
