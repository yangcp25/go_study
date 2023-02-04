package chain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSensitiveWordFilterChain_Filter(t *testing.T) {
	handler := SensitiveWordFilterChain{}
	handler.addFilter(&filter1{})
	handler.addFilter(&filter2{})
	assert.Equal(t, false, handler.Filter("content come in"))
}
