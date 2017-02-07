package linear_search

func ls(n []int, key int) int {
	return search(n, 0, len(n), key)
}

func search(n []int, l, h, k int) int {

	for i := l; i < h; i++ {
		if n[i] == k {
			return i
		}
	}

	return -1
}
