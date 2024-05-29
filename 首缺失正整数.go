func firstMissingPositive(nums []int) int {
    hashMap := make(map[int]struct{})
    for i:=0; i < len(nums); i++ {
        hashMap[nums[i]] = struct{}{}  
    }
    
    for i:=1; i<=len(nums); i++ { // i=1因为是最小正整数
        if _, ok := hashMap[i]; !ok { // 从1找起，按顺序找，在不在map中，若不在map中就直接返回该数。因为数字是能从1开始连续往后
                                      // 1,2,3,4,5这样占位 若不在该map中那就是它了
            return i
        }
    } 
    
    return len(nums) + 1 // 若1,2,3,4,5都给占用了，那只能返回6，因为他是目前缺失唯一最小的整数
}
