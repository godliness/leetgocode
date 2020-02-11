func search(nums []int, target int) int {
    n := len(nums)
    l := 0
    r := n - 1
    
    for l <= r {
        if nums[l] == target {
            return l
        }
        
        if nums[r] == target {
            return r
        }
        
        mid := (l + r) / 2
        
        if nums[mid] == target {
            return mid
        }
        
        if nums[l] > nums[r] { 找到旋转的那个索引，找到后，左右部分就都是有序的了，然后就可以使用二分查找了
            l += 1
            r -= 1
        } else {
            if nums[mid] > target {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
    }
    
    return -1
}
