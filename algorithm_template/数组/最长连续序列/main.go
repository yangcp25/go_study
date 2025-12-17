package main

func main() {

}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}

	longest := 0

	for n := range set {
		// 只从“起点”开始
		if !set[n-1] {
			length := 1
			cur := n

			for set[cur+1] {
				cur++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
