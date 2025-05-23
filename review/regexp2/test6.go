package regexp2

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
)

func rege1() {
	//rege := regexp.MustCompile(`((\w){2,})\1`)
	//
	//test1 := rege.FindAllStringSubmatch("xxabxxab", -1)
	//fmt.Println(test1)

	rege := regexp.MustCompile(`^(?:`+
		`(?=(.*[A-Z]))(?=(.*[a-z]))(?=(.*\d))|`+ // 大写 + 小写 + 数字
		`(?=(.*[A-Z]))(?=(.*[a-z]))(?=(.*[!@#\$%\^&\*\-_]))|`+ // 大写 + 小写 + 特殊
		`(?=(.*[A-Z]))(?=(.*\d))(?=(.*[!@#\$%\^&\*\-_]))|`+ // 大写 + 数字 + 特殊
		`(?=(.*[a-z]))(?=(.*\d))(?=(.*[!@#\$%\^&\*\-_]))`+ // 小写 + 数字 + 特殊
		`).{8,}$`, regexp.None,
	)

	test1, _ := rege.FindStringMatch("Az$")
	fmt.Println(test1)

	re := regexp.MustCompile(`^(?!.*(.{3,}).*\1).+$`, 0)
	ok, ok2 := re.MatchString("1212")
	fmt.Println(ok, ok2)
}
