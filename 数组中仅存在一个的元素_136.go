func singleNumber(nums []int) int {
    tmp := make(map[int]int)
    sum1 := 0
    sum2 := 0
    for _, v := range nums {
        sum1 += v
        tmp[v] = 1
    }
    for i := range tmp {
        sum2 += i
    }
    return 2*sum2 - sum1
}

func singleNumber(nums []int) int {

	var n int

	for _, v := range nums {
		n = n ^ v
	}

	return n
    
}
