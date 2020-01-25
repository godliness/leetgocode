func climbStairs(n int) int {
    cache := make(map[int]int)
    cache[1] = 1
    cache[2] = 2
    
    return csrec(n, cache)
}

func csrec(n int, cache map[int]int) int {
    if val, ok := cache[n]; ok {
        return val
    }
    
    ret := csrec(n-1, cache) + csrec(n-2, cache)
    cache[n] = ret
    return ret
}

===========


func climbStairs(n int) int {
     if n<=1 {
        return 1
    }
    dp := make([]int, n)
    dp[0]=1
    dp[1]=2
    for i:=2;i<n ;i++  {
        dp[i]=dp[i-1]+dp[i-2]
    }
    return dp[n-1]
}
