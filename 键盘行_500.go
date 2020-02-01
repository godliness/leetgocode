func lower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return c + 32
    }
    return c
}

func findWords(words []string) []string {
    levels := map[byte]int {
        'q' : 1,
        'w' : 1,
        'e' : 1,
        'r' : 1,
        't' : 1,
        'y' : 1, 
        'u' : 1,
        'i' : 1,
        'o' : 1,
        'p' : 1,
        
        'a' : 2,
        's' : 2,
        'd' : 2,
        'f' : 2,
        'g' : 2,
        'h' : 2,
        'j' : 2,
        'k' : 2,
        'l' : 2,
        
        'z' : 3,
        'x' : 3,
        'c' : 3,
        'v' : 3,
        'b' : 3,
        'n' : 3,
        'm' : 3,
    }
    var ans []string
    for _, word := range words {
        level := levels[lower(word[0])]
        for i := 1; i < len(word); i++ {
            if levels[lower(word[i])] != level {
                level = 0
                break
            }
        }
        if level > 0 {
            ans = append(ans, word)
        }
    }
    return ans
}
