func maxSubArray(nums []int) int {   
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    ans := dp[0]
    for i := 1; i < len(nums); i++ {
        dp[i] = max(dp[i-1]+nums[i], nums[i]) // dp[i] 从不确定的位置到第i的位置子数组的最大和
        ans = max(ans, dp[i]) // 选出dp[i]中最大的
    }
    return ans    
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
