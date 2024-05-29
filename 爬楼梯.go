func climbStairs(n int) int {
     if n<=1 {
        return 1
    }
    dp := make([]int, n)
    dp[0]=1
    dp[1]=2
    for i:=2;i<n ;i++  {       
        dp[i]=dp[i-1]+dp[i-2] // dp[i]代表走到第i台阶时候所用的走法数  公式：我走到第i个台阶的总走法等于走到第i-1阶时与第i-2阶时用过的方法总和。
    }
    return dp[n-1]
}
