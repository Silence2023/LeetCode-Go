package leetcode

func longestPalindrome(s string) string {
	palindromic_string := ""
	// 判断完整的字符串是不是回文
	if isPalindromic(s) {
		return s
	}

	for i := 0; i < len(s)-1; i++ {
		for j := len(s) - 1; i <= j; j-- {
			if len(palindromic_string) >= len(s[i:j+1]) {
				break
			}
			if !isPalindromic(s[i : j+1]) {
				continue
			}

			// 找到了回文
			if len(palindromic_string) < len(s[i:j+1]) {
				palindromic_string = s[i : j+1]
			}
		}
	}
	return palindromic_string
}

func isPalindromic(s string) bool {
	// 判断字符串是否是回文
	for i := 0; i < len(s)/2; i++ {
		if !(s[i] == s[len(s)-i-1]) {
			return false
		}
	}
	return true
}
