func searchInsert(nums []int, target int) int {
  l, r := 0, len(nums)-1
  
  for l <= r {
    m := l + (r-l)/2
    if (nums[m] < target) {
      l = m+1
    } else if (nums[m] > target) {
      r = m-1
    } else {
      return m
    }
  }
  
  return l
}
二分查找
