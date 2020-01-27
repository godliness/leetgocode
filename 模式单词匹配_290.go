func wordPattern(pattern string, str string) bool {
    s := strings.Fields(str)
    if len(pattern) != len(s) { return false }
    for i := 0; i < len(s); i++ {
        for j := i+1; j < len(s); j++ {
            if pattern[i] == pattern[j] && s[i] != s[j] || pattern[i] != pattern[j] && s[i] == s[j] {
                return false
            }
        }
    }
    return true
}
