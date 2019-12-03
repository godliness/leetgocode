package main

func isAnagram(s string, t string) bool {

	tmp := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		tmp[s[i]]++
		tmp[t[i]]--
	}
	for _, v := range tmp {
		if v != 0 {
			return false
		}
	}

	return true

}
