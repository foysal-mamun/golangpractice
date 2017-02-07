package binary_search

func bs(a []int, key int) int {
	return search(a, 0, len(a), key)
}

func search(a []int, l, h, k int) int {

	if h < l {
		return l - 1
	}

	m := l + (h-l)/2

	if k == a[m] {
		return m
	} else if k < a[m] {
		return search(a, l, h-1, k)
	} else {
		return search(a, l+1, h, k)
	}
}
