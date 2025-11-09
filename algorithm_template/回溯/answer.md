```go
func backtrack(path []int, nums []int, res *[][]int, used []bool) {
	// 1. 结束条件
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	// 2. 遍历所有选择
	for i := 0; i < len(nums); i++ {
		// 剪枝（避免重复选择）
		if used[i] {
			continue
		}

		// 做选择
		used[i] = true
		path = append(path, nums[i])

		// 进入下一层
		backtrack(path, nums, res, used)

		// 撤销选择
		used[i] = false
		path = path[:len(path)-1]
	}
}

// 全排列

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			tmp := append([]int{}, path...)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs([]int{})
	return res
}

func permute2(nums []int) [][]int {
	var res [][]int
	var dfs func(int)
	dfs = func(first int) {
		// 所有位置都固定完了
		if first == len(nums) {
			tmp := append([]int{}, nums...) // 拷贝结果
			res = append(res, tmp)
			return
		}

		// 从当前位置开始，依次和后面的交换
		for i := first; i < len(nums); i++ {
			nums[first], nums[i] = nums[i], nums[first] // 交换
			dfs(first + 1)                              // 递归固定下一个位置
			nums[first], nums[i] = nums[i], nums[first] // 撤销交换（回溯）
		}
	}
	dfs(0)
	return res
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums) // 关键：先排序
	used := make([]bool, len(nums))
	path := make([]int, 0, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			tmp := append([]int(nil), path...)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 同层去重：如果当前位置与前一个相同，且前一个在本层未被使用，
			// 说明当前选择会产生与前一个相同的分支 => 跳过
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			// 选择
			used[i] = true
			path = append(path, nums[i])
			dfs()
			// 撤销选择
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs()
	return res
}

func permuteUnique2(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	var dfs func(int)
	dfs = func(first int) {
		if first == len(nums) {
			tmp := append([]int{}, nums...)
			res = append(res, tmp)
			return
		}
		used := make(map[int]bool)
		for i := first; i < len(nums); i++ {
			if used[nums[i]] { // 同一层剪枝
				continue
			}
			used[nums[i]] = true
			nums[first], nums[i] = nums[i], nums[first]
			dfs(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	dfs(0)
	return res
}



```
```go
// 组合总数
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var dfs func(int, int)
	dfs = func(start, sum int) {
		if sum == target {
			res = append(res, append([]int(nil), path...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(i, sum+candidates[i]) // 可以重复选 i
			path = path[:len(path)-1] // 回溯
		}
	}
	dfs(0, 0)
	return res
}
```