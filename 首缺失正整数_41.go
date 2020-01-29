func firstMissingPositive(nums []int) int {
    hashmap := make(map[int]int)
    for i:=0; i<len(nums);i++ {
        hashmap[nums[i]] = i
    }
    for i:=1; i<=len(nums);i++ {
        if _ , ok := hashmap[i]; !ok {
            return i
        }
    }
    return len(nums)+1
}
