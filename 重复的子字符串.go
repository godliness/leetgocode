func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	ss := (s + s)[1 : len(s)*2-1]
	return strings.Contains(ss, s)
}
