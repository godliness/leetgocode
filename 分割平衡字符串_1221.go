func balancedStringSplit(s string) int {
    var cnt, L, R int
    
    for i:=0; i<len(s); i++{
        if L==R{
            cnt++
            L=0
            R=0
        }
        if s[i]=='L' {
            L++
        }else{
            R++
        }
    }
    return cnt
}
