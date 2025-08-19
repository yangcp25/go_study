package test1

//题目
// 题目描述输入两个字符串S和L，都只包含英文小写字母。S长度<=100，L长度<=500,000。
//判定S是否是L的有效子串。判定规则：S中的每个字符在L中都能找到（可以不连续），且S在Ｌ中字符的前后顺序与S中顺序要保持一致。
//（例如，S=”ace”是L=”abcde”的一个子序列且有效字符是a、c、e，而”aec”不是有效子序列，且有效字符只有a、e）
//输入输出
//输入输入两个字符串S和L，都只包含英文小写字母。S长度<=100，L长度<=500,000。
//先输入S，再输入L，每个字符串占一行。输出S串最后一个有效字符在L中的位置。
//（首位从0开始计算，无有效字符返回-1）

func SubStrLasPos(str string, sub string) int {
	// aabcdabc 8
	// abc
	findKey := -1
	path := make([]byte, 0)
	var findLastKey func(int, string)
	findLastKey = func(index int, str string) {
		if len(path) == len(sub) {
			findKey = index - 1
			return
		}
		for i := index; i < len(str); i++ {
			if str[i] == sub[len(path)] {
				path = append(path, str[index])
				findLastKey(i+1, str)
				path = path[:len(path)-1]
			}
		}
	}
	findLastKey(0, str)

	return findKey
}
