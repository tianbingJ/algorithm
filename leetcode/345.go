package leetcode

func reverseVowels(s string) string {
	ss := []byte(s)
	for i, j := 0, len(ss) - 1; i < j; {
		if isVowel(ss[i]) && isVowel(ss[j]) {
			ss[i], ss[j] = ss[j], ss[i]
			i ++
			j --
			continue
		}
		if !isVowel(ss[i]) {
			i ++
		} else {
			j--
		}
	}
	return string(ss)
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U';
}
