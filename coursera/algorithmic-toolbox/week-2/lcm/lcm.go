package lcm

import (
	"github.com/foysal-mamun/golangpractice/coursera/algorithmic-toolbox/week-2/gcd"
)

func NaiveSolution(a, b int64) int64 {
	var i int64 = 1
	for ; i <= a*b; i++ {
		if i%a == 0 && i%b == 0 {
			return i
		}
	}
	return a * b
}

func EfficientSolution(a, b int64) int64 {
	return (a * b) / gcd.EffectiveAlgo(a, b)
}
