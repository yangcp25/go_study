package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	notify := &Subject{}
	notify.Register(Observe1{})
	notify.Register(Observe2{})
	notify.Notify("haha")
}
