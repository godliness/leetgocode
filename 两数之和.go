func twoSum(nums []int, target int) []int {
    valueStore := make(map[int]int)
    var result []int
    for i:= 0; i < len(nums); i++ {
        if value, ok := valueStore[target - nums[i]]; ok {
            result = append(result, value, i)
            return result
        }
        valueStore[nums[i]] = i
    }
    return result
}
