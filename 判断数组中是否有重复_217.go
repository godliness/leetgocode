func containsDuplicate(nums []int) bool {
    var mp = make(map[int]int)
    
    for idx, value := range nums {
        if _, ok := mp[value]; ok {
            return true
        }
        mp[value] = idx
    }

    return false
}
