package selectionSort

type SelectionSort struct {
	Array []int
}

func NewSelectionSort(arr []int) *SelectionSort {
	return &SelectionSort{Array: arr}
}

func (s SelectionSort) minIndex(startIndex int) int {

	for i := startIndex + 1; i < len(s.Array); i++ {
		if s.Array[startIndex] > s.Array[i] {
			startIndex = i
		}
	}

	return startIndex
}

func (s *SelectionSort) swap(val1, val2 *int) {
	*val1, *val2 = *val2, *val1

}

func (s *SelectionSort) DoSort() {

	for i := 0; i < len(s.Array)-1; i++ {
		minIndex := s.minIndex(i + 1)
		if s.Array[minIndex] < s.Array[i] {
			s.swap(&s.Array[i], &s.Array[minIndex])
		}
	}
}
