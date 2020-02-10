func searchInsert(nums []int, target int) int {
  l, r := 0, len(nums)-1
  
  for l <= r { // 最后左右边界交错以后证明已经找到了位置
    m := l + (r-l)/2 // 这里之所以加l, 是要确定索引而不是数值
    if (nums[m] < target) {
      l = m+1
    } else if (nums[m] > target) {
      r = m-1
    } else {
      return m // 若正好等于 那么就代表 找到了该值了
    }
  }
  
  return l // 左边界才是这个target应该插入的位置
}
二分查找
