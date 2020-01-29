func firstMissingPositive(nums []int) int {
    hashMap := make(map[int]struct{})
    for i:=0; i < len(nums); i++ {
        hashMap[nums[i]] = struct{}{}
    }
    
    for i:=1; i<=len(nums); i++ {
        if _, ok := hashMap[i]; !ok {
            return i
        }
    } 
    
    return len(nums) + 1
}
