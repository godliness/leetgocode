import "math/rand"
type Solution struct {
    a []int
}


func Constructor(nums []int) Solution {
    return Solution{nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.a
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    tmp := make([]int, len(this.a))
    copy(tmp, this.a)
    for i := 0; i< len(tmp); i++ {
        r := rand.Intn(len(tmp))
        tmp[i], tmp[r] = tmp[r], tmp[i]
    }
    return tmp
}
