func lengthOfLastWord(s string) int {
    words := strings.Split(strings.TrimSpace(s), " ")
    
    if len(words) == 0 {
        return 0
    }
    
    return len(words[len(words) - 1])
}
