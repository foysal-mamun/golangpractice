package gcd

func NativeAlgo(a int64, b int64) int64 {
	var result int64 = 0
	limit := a
	if a < b {
		limit = b
	}
	var i int64 = 1
	for ; i <= limit; i++ {
		if a%i == 0 && b%i == 0 {
			result = i
		}
	}

	return result
}

func EffectiveAlgo(a int64, b int64) int64 {
	c := a % b
	if c == 0 {
		return b
	}
	return EffectiveAlgo(b, c)
}
