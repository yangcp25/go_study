package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// str = "13.3.3.24.4.4.5"
	str := "1,3,3,3,2,4,4,4,5"
	strS := strings.Split(str, ",")
	checkMap := make(map[string]int)

	sortS := make(SortItem, 0)
	for index, s := range strS {
		if sk, ok := checkMap[s]; ok {
			sortS[sk].Count++
		} else {
			checkMap[s] = len(sortS)
			sortS = append(sortS, Item{
				Str:   s,
				Count: 1,
				Index: index,
			})
		}
	}
	sort.Sort(sortS)
	strRes := make([]string, 0)
	for _, item := range sortS {
		strRes = append(strRes, item.Str)
	}
	fmt.Println(strings.Join(strRes, ","))

}

type Item struct {
	Str   string
	Count int
	Index int
}

type SortItem []Item

func (s SortItem) Len() int {
	return len(s)
}
func (s SortItem) Less(i, j int) bool {
	if s[i].Count == s[j].Count {
		return s[i].Index < s[j].Index
	}
	return s[i].Count > s[j].Count
}
func (s SortItem) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
