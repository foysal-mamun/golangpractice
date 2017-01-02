package changing_money

func NativeSolution(m int, currentindex int) int {

	coins := [3]int{10, 5, 1}
	if m == 0 || m%coins[currentindex] == 0 {
		return m
	}

	return m/coins[currentindex] + NativeSolution(m%coins[currentindex], currentindex+1)
}
