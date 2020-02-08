func maxProduct(nums []int) int {
    n := len(nums)
    if n <= 0 {
        return 0
    }
    
    dp := make([][]int, n)
    for i:=0; i < n; i++ {
        dp[i] = make([]int, 2)
    }

    var res int
    dp[0][0], dp[0][1], res = nums[0], nums[0], nums[0]
    for i:=1; i<len(nums); i++ {
        if nums[i] >= 0 {
            dp[i][0] = max(dp[i-1][0]*nums[i], nums[i])
            dp[i][1] = min(dp[i-1][1]*nums[i], nums[i])
        } else {
            dp[i][0] = max(dp[i-1][1]*nums[i], nums[i])
            dp[i][1] = min(dp[i-1][0]*nums[i], nums[i])
        }
        res = max(res, dp[i][0])
    }
    
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
