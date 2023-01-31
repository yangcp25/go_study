package adapter

import (
	"testing"
)

func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口
	var a IRun = &AliyunAdapter{
		aliyun: aliyun{},
	}

	a.CrateServer("aliyun", "2")
}

func TestAwsClientAdapter_CreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口
	var a IRun = &tenxunyunAdapter{
		tenxunyun: tenxunyun{},
	}

	a.CrateServer("tunxunyun", "3")
}
