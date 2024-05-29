func titleToNumber(s string) int {
    rs := []rune(s)
    ret := 0
    for i := 0; i < len(rs); i++ {
        ret = ret*26 + letterToNum(rs[i])
    }
    return ret
}

func letterToNum(r rune) int {
    return int(r - 'A') + 1
}
