func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0 // 凑够0元钱当然需要0个硬币
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // 凑够i元钱你不能比要凑够的钱数还要大，因为最小硬币面值就是1元
		for _, coin := range coins { //这多的硬币种类
			if coin <= i { // 硬币面值要小于最终值
				dp[i] = min(dp[i], dp[i-coin]+1) //dp[i] i元钱所需要的最小数量硬币，所有硬种都考虑一遍，取其中的最小值，迭代判断去最小值
			}
		}
	}

	if dp[amount] > amount { // 说明没有找到，若找到值就会被覆盖
		return -1
	}

	return dp[amount]
}
