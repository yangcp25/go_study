package proxy

import (
	"testing"
)

func Test_user_Login(t *testing.T) {
	proxyObj := newUserProxy(&user{})
	proxyObj.Login()
}
