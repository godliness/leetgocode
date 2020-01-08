func majorityElement(nums []int) int {
    m := make(map[int]int)

	for _, v := range nums {
		m[v]++
		if m[v] > len(nums) / 2 {
			return v
		}
	}

	return 0
}
