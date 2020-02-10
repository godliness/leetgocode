func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    if n == 0 {
        return 0
    } 
    dp := make([][]int, n) // 二维动态方程
    dp[n-1] = triangle[n-1] // 从后往前推

    for m:=0; m<n-1; m++ {
        dp[m] = make([]int, len(triangle[m]))  // 初始化二维状态方程
    }

    for i:=n-2; i>=0; i-- { //从后往前推， 倒是第二层
        for j:=0; j<len(triangle[i]); j++ {
            dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j] // 三角形底边相邻两个点选最小的 再加上顶点 一直往上推 每一层都是最小的
                                                                      // 最后选出的也是最小的
        }
    }

    return dp[0][0] // 推到最上方后一定会得到最小和，因为
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
