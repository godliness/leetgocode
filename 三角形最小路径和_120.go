func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    if n == 0 {
        return 0
    } 
    dp := make([][]int, n)
    dp[n-1] = triangle[n-1]

    for m:=0; m<n-1; m++ {
        dp[m] = make([]int, len(triangle[m]))
    }

    for i:=n-2; i>=0; i-- {
        for j:=0; j<len(triangle[i]); j++ {
            dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
        }
    }

    return dp[0][0]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
