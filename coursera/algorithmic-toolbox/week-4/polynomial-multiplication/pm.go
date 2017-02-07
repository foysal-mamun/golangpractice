package polynomial_multiplication

func nativeSolution(A, B []int) []int {
	C := make([]int, len(A)+len(B)-1)

	for i, a := range A {
		for j, b := range B {
			C[i+j] += a * b
		}
	}

	return C
}
