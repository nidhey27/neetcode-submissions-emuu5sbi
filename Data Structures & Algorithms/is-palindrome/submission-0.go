func isPalindrome(s string) bool {

	l := 0
	r := len(s) - 1

	for l < r {
		for l < r && !isAlphaNum(s[l]) {
			l += 1
		}
		for l < r && !isAlphaNum(s[r]) {
			r -= 1
		}

		if toLowerCase(s[l]) != toLowerCase(s[r]) {
			return false
		}

		l += 1
		r -= 1
	}
	
	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLowerCase(str byte) string {
	return strings.ToLower(string(str))
}
