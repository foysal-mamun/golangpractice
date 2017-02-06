package maximizing_your_salary

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type intSort []int

func (i intSort) Len() int      { return len(i) }
func (i intSort) Swap(a, b int) { i[a], i[b] = i[b], i[a] }
func (i intSort) Less(a, b int) bool {
	return compareTwoArray([]int{i[a], i[b]}, []int{i[b], i[a]})
}

func mys(nums []int) string {

	sort.Sort(intSort(nums))

	return arrayToString(nums, "")
}

func compareTwoArray(a, b []int) bool {

	i, _ := strconv.Atoi(arrayToString(a, ""))
	j, _ := strconv.Atoi(arrayToString(b, ""))

	return i > j
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
