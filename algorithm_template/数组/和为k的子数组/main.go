package main

func main() {

}

func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	pre := 0
	count := 0
	for _, x := range nums {
		pre += x
		if v, ok := m[pre-k]; ok {
			count += v
		}
		m[pre]++
	}
	return count
}
