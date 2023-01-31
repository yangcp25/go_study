package bridge

import (
	"testing"
)

func TestNewEmailArray(t *testing.T) {
	sender := NewEmailArray([]string{"820", "765"})

	notifyHandler := newNotify(sender)

	notifyHandler.notifyTo("ni m p")
}
