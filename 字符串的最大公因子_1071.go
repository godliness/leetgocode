func gcdOfStrings(str1 string, str2 string) string {
    prefix := str1
    
    for prefix != "" {
        if strings.HasPrefix(str1, prefix) && strings.HasPrefix(str2, prefix) &&
        len(str1) % len(prefix) == 0 && len(str2) % len(prefix) == 0 {
            return prefix
        }
        
        prefix = prefix[:len(prefix)-1]
    }
    
    return ""
}
