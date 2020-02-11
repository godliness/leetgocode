func firstMissingPositive(nums []int) int {
    hashMap := make(map[int]struct{})
    for i:=0; i < len(nums); i++ {
        hashMap[nums[i]] = struct{}{}  
    }
    
    for i:=1; i<=len(nums); i++ { // i=1因为是最小正整数
        if _, ok := hashMap[i]; !ok { // 从1找起，按顺序找，在不在map中，若不在map中就直接返回该数。
            return i
        }
    } 
    
    return len(nums) + 1 
}
