func gcdOfStrings(str1 string, str2 string) string {
    if len(str1) == 0 || str1 == str2 {
        return str1
    }
    if len(str2) == 0 {
        return str2
    }
    if len(str1) > len(str2) {
        x := strings.Index(str1, str2)
        if x > -1 {
            return gcdOfStrings(str2, str1[x+len(str2):])
        }
    } else {
        y := strings.Index(str2, str1)
        if y > -1 {
            return gcdOfStrings(str1, str2[y+len(str1):])
        }
    }
    
    return ""
}
