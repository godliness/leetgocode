func maxProfit(prices []int) int {
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
