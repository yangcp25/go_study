package iterator

type IIterator interface {
	hasNext() bool
	next()
	currentItem() interface{}
}

type arrayInt []int

type arrayIntIterator struct {
	arrayInt arrayInt
	index    int
}

func (a *arrayIntIterator) hasNext() bool {
	return a.index < (len(a.arrayInt) - 1)
}

func (a *arrayIntIterator) next() {
	a.index++
}

func (a *arrayIntIterator) currentItem() interface{} {
	return a.arrayInt[a.index]
}

func (a arrayInt) iterator() *arrayIntIterator {
	return &arrayIntIterator{
		arrayInt: a,
		index:    0,
	}
}
