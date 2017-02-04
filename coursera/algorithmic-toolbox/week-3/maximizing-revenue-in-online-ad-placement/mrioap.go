package maximizing_revenue_in_online_ad_lacement

import "sort"

func NativeSolution(profitPerClick, averageNumberOfClicks []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(profitPerClick)))
	sort.Sort(sort.Reverse(sort.IntSlice(averageNumberOfClicks)))

	totalRevenue := 0
	for i, v := range profitPerClick {
		totalRevenue += v * averageNumberOfClicks[i]
	}

	return totalRevenue
}
