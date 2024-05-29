func getSum(a int, b int) int {
    if b == 0{
        return a
    } 
    return getSum(a ^ b, (a & b)<<1)
}

a + b 的问题拆分为 (a 和 b 的无进位结果) + (a 和 b 的进位结果)
无进位加法使用异或运算计算得出
进位结果使用与运算和移位运算计算得出
循环此过程，直到进位为 0
