package maximizing_your_salary

import (
	"fmt"
	"sort"
	"strings"
)

type intSort []int

func (i intSort) Len() int      { return len(i) }
func (i intSort) Swap(a, b int) { i[a], i[b] = i[b], i[a] }
func (i intSort) Less(a, b int) bool {
	if firstDigit(i[a]) == firstDigit(i[b]) {
		return i[a] > i[b]
	}
	return firstDigit(i[a]) > firstDigit(i[b])
}

func mys(nums []int) string {

	sort.Sort(intSort(nums))

	return arrayToString(nums, "")
}

func firstDigit(num int) int {
	l := len(string(num))
	if l == 0 {
		return 0
	}
	return num / l
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
