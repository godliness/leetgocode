func findLengthOfLCIS(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    
    max := 1
    cur := 1

    for i, _ := range nums {
        if i == 0 {
            continue
        }
        if nums[i] > nums[i-1] {
            cur = cur + 1
        } else {
            cur = 1
        }
        if cur > max {
            max = cur
        }
    }
    return max
    
}
