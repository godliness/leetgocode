type NumArray struct {
    n []int
    dp []int
}

func Constructor(nums []int) NumArray {
    dp := make([]int, len(nums)+1)
    dp[0] = 0 // 加到0位置和就是0
    for i:=0; i<len(nums); i++ {
        dp[i+1] = dp[i] + nums[i] // dp[i]代表数组加到i位置的和 
    }
    return NumArray{nums, dp}
}

func (this *NumArray) SumRange(i int, j int) int {
    return this.dp[j+1] - this.dp[i]  //加到第j+1位置的和减去加到第i位置的和
}
