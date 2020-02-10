func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n, max := len(nums), 1
	d := make([]int, n)
	for i:=0; i<n; i++ {
		d[i] = 1  // 至少会为1 【5， 4， 3， 2， 1】
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] { // nums[i]前面的 0...i-1的元素是否比nums[i]小
				d[i] = Max(d[i], d[j] + 1)  // 截止到d[i]个位置我前面最长上升序列的个数，因为0-j的最长上升子序列个数有高有底
			}
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
