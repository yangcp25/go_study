package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]string, 0)

	for i := 0; i < n; i++ {
		var temp string
		fmt.Scan(&temp)
		nums = append(nums, temp)
	}
	usedAll := make(map[string]struct{}, 0)

	var backtrack func(nums, path []string, used map[int]struct{}, ans *[]string)
	backtrack = func(nums, path []string, used map[int]struct{}, ans *[]string) {
		if len(path) == n {
			str := strings.Join(path, "")
			if _, ok := usedAll[str]; !ok {
				*ans = append(*ans, str)
				usedAll[str] = struct{}{}
			}
		}

		for i := 0; i < n; i++ {
			if _, ok := used[i]; !ok {
				used[i] = struct{}{}
				path = append(path, nums[i])
				backtrack(nums, path, used, ans)
				// 回溯
				path = path[:len(path)-1]
				delete(used, i)
			}
		}
	}
	path := make([]string, 0)
	used := make(map[int]struct{}, 0)
	ans := make([]string, 0)
	backtrack(nums, path, used, &ans)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
