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
            dp[i][0] = max(dp[i-1][0]*nums[i], nums[i]) // 乘积最大子连续序列的积  dp[i]固定右边边界
            dp[i][1] = min(dp[i-1][1]*nums[i], nums[i]) // 乘积最小子连续序列的积  dp[i]固定右边边界
        } else {
            dp[i][0] = max(dp[i-1][1]*nums[i], nums[i]) // 当为负数的时候，用乘积最小子连续序列的积乘以当前负数，为了能够求出最大值，因为负负为正
            dp[i][1] = min(dp[i-1][0]*nums[i], nums[i]) // 当为负数的时候，用乘积最大子连续序列的积乘以当前负数，为了能够求出最小值，因为正负为负
        }
        res = max(res, dp[i][0]) //最后选出最大值
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
