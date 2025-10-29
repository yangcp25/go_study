``` go
// 两数之和
func towSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, number := range numbers {
		if j, ok := m[target-number]; ok {
			return []int{j, i}
		}
		m[number] = i
	}
	return nil
}

// 三数字和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

    // 盛雨水
	start, end := 0, len(height)-1
	area := 0
	for start < end {
		a := (end - start) * Min(height[start], height[end])
		if a > area {
			area = a
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return area
	
	
// 快速排序
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// 分区操作
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

```