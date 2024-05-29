func uniquePaths(m int, n int) int {
    result := make([][]int, m)
    for i:=0; i<m; i++{
        result[i] = make([]int,n)
    }
    
    for i:=0; i<m;i++{
        result[i][0] = 1 // 最边缘的只能一种路径 不是向右就是向下
    }
    
    for j:=1; j<n; j++{
        result[0][j] = 1 // 最边缘的只能一种路径
    }
    
    for i:=1;i<m; i++{
        for j:=1; j<n; j++{
            result[i][j] = result[i-1][j] + result[i][j-1] // result[i][j]的位置走到终点就是result[i-1][j] + result[i][j-1]走法的和
            //最后相当于从后往前推，把所有的格子填满，得到从起点到终点的所有步数
        }
    }
    
    return result[m-1][n-1]
}
