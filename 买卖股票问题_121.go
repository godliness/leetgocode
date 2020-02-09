func maxProfit(prices []int) int {
    if len(prices) <=0 {
        return 0
    }
    minV := prices[0]
    ans := 0
    for i:=1; i<len(prices); i++ {
        ans = max(ans, prices[i] - minV)
        minV = min(minV, prices[i])
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
