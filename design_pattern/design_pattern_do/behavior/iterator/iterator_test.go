package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayInt_Iterator(t *testing.T) {
	data := arrayInt{1, 3, 5, 7, 8}
	iterator := data.iterator()
	// i 用于测试
	i := 0
	for iterator.hasNext() {
		assert.Equal(t, data[i], iterator.currentItem())
		iterator.next()
		i++
	}
}
