func maxProfit(prices []int) int { //只能买一次卖一次
    if len(prices) <=0 {
        return 0
    }
    minV := prices[0]
    ans := 0
    for i:=1; i<len(prices); i++ { 
        ans = max(ans, prices[i] - minV) //后天的股价，减去之前的最小股价
        minV = min(minV, prices[i]) //求出最小股价
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func maxProfit(prices []int) int {
	profit, sell := 0, false

	for i := 0; i < len(prices); i++ {
		if sell {
			profit, sell = profit+prices[i]-prices[i-1], false
		}
		sell = i < len(prices)-1 && prices[i] < prices[i+1]
	}
	return profit
}
