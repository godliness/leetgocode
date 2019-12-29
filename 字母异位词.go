
func isAnagram(s string, t string) bool {
	m1 := map[byte]int{}
	m2 := map[byte]int{}
	if len(s) != len(t) {
		return false
	}
	
	for i := 0; i < len(s); i++ {
		m1[s[i]] = m1[s[i]] + 1
		m2[t[i]] = m2[t[i]] + 1
	}
	
	return reflect.DeepEqual(m1, m2)
}
