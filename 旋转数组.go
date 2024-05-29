func rotate(nums []int, k int)  {
    n := len(nums)
    r := make([]int, n)
    for i:=0; i < n; i++ {
        r[(i+k)%n] = nums[i]
    }
    
    copy(nums, r)
}
