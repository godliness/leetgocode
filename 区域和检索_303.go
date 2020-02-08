type NumArray struct {
    n []int
    dp []int
}


func Constructor(nums []int) NumArray {
    dp := make([]int, len(nums)+1)
    dp[0] = 0
    for i:=0; i<len(nums); i++ {
        dp[i+1] = dp[i] + nums[i]
    }
    return NumArray{nums, dp}
}


func (this *NumArray) SumRange(i int, j int) int {
    return this.dp[j+1] - this.dp[i]
}
