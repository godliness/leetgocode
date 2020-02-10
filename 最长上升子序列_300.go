func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n, max := len(nums), 1
	d := make([]int, n)
	d[0] = 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			cur := 1
			if nums[i] > nums[j] { // nums[i]前面的 0...i-1的元素是否比nums[i]小，若小的话，计数+1
				cur = d[j] + 1 // d[j] 代表到第j个位置的数已经有多少个上升数字了(算上j)， 目前nums[i] > nums[j]说明又多了一个，所以是
			}
			d[i] = Max(d[i], cur) // 我们定位在d[i]的位置，然后其次查找前面0-j的上升子序列和+1的结果，选一个最大的，因为可能d[j] 0-j 有大有小
		}
		max = Max(max, d[i]) // 最后从所有的d[i]中i为0到n-1，选一个最大的。
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
