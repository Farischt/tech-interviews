package sliding_window

import "math"

/* ------------------------------- Max Profit ------------------------------- */
// Time Complexity O(n)
// Space Complexity O(1)
// Difficulty: Easy
func MaxProfit(prices []int) int {
	highestProfit := 0
	left := 0
	right := 1

	for right < len(prices) {
		for right < len(prices) && left <= right && prices[left] > prices[right] {
			left++
		}

		currentProfit := prices[right] - prices[left]
		if currentProfit > highestProfit {
			highestProfit = currentProfit
		}
		right++
	}

	return highestProfit
}

func MaxProfitBis(prices []int) int {
	highestProfit := 0
	left := 0
	right := 1

	for right < len(prices) {
		if prices[left] < prices[right] {
			currentProfit := prices[right] - prices[left]
			highestProfit = int(math.Max(float64(highestProfit), float64(currentProfit)))
		} else {
			left = right
		}
		right++
	}

	return highestProfit
}
