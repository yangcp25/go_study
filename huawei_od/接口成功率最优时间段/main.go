package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	k, _ := strconv.Atoi(input.Text())

	input.Scan()
	nums := input.Text()
	numsSlice := strings.Split(nums, " ")

	/**
	1
	0 1 2 3 4

	6
	0 0 100 2 2 99 0 2


	maxL = 0
	path := make([][]int, 0)
	sum += nums[j]
	sum/j-i+1 <= k  j-i+1 > maxL maxL= j-i+1
	else sum > k && i <j i++ sum -= nums[i] path = path[:0]
	*/

	maxL, sum, i, j := 0, 0, 0, 0
	path := make([][]int, len(numsSlice)+1)

	for j < len(numsSlice) {
		num, _ := strconv.Atoi(numsSlice[j])
		sum += num
		if float64(sum)/float64(j-i+1) <= float64(k) {
			if j-i+1 >= maxL {
				maxL = j - i + 1
				path[maxL] = append(path[maxL], i)
				path[maxL] = append(path[maxL], j)
			}
		} else {
			for float64(sum)/float64(j-i+1) > float64(k) && i <= j {
				num, _ := strconv.Atoi(numsSlice[i])
				sum -= num
				i++
			}
		}
		j++
	}

	if len(path[maxL]) > 0 {
		str := make([]string, 0)
		sort.Slice(path[maxL], func(i, j int) bool {
			return path[maxL][i] < path[maxL][j]
		})
		for i := 0; i < len(path[maxL]); i = i + 2 {
			str2 := make([]string, 0)
			str2 = append(str2, strconv.Itoa(path[maxL][i]), strconv.Itoa(path[maxL][i+1]))
			str = append(str, strings.Join(str2, "-"))
		}
		fmt.Println(strings.Join(str, " "))
	} else {
		fmt.Println("NULL")
	}

}
