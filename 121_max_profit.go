package leetcode

func maxProfit(prices []int) int {
	maxProfit, lowestPrice := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] <= lowestPrice {
			lowestPrice = prices[i]
			// maxProfit = 0
		} else {
			if maxProfit < prices[i]-lowestPrice {
				maxProfit = prices[i] - lowestPrice
			}
		}
	}
	return maxProfit
}
