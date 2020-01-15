func generateParenthesis(n int) []string {
    results := make([]string, 0, 2*n)
    gen(0, 0, n, "", &results)
    return results
}

func gen(left, right, n int, tmp string, results *[]string) {
    if left == n && right == n {
        *results = append(*results, tmp)
        return
    }
    
    if left < n {
        gen(left + 1, right, n, tmp + "(", results)
    }
    
    if left > right && right < n {
        gen(left, right+1, n, tmp + ")", results)
    }
}
