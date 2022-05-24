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
	// // 判断字符串是否是回文
	// for i := 0; i < len(s)/2; i++ {
	// 	if !(s[i] == s[len(s)-i-1]) {
	// 		return false
	// 	}
	// }
	// 判断长度是偶数还是奇数
	length := len(s)
	result := length % 2 // 余数
	mid := length / 2    // 轴心
	for i := 1; i <= mid; i++ {
		if result == 0 {
			// 偶数
			if s[mid-i] != s[mid+i-1] {
				return false
			}
		} else {
			// 奇数
			if s[mid-i] != s[mid+i] {
				return false
			}
		}
	}
	return true
}

// 解法二
// 中心扩散法，如果字符串长度是偶数，那么回文的轴心就是虚拟的，如果长度是奇数，那么就是中间的为轴心，向两边扩散
func longestPalindrome2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}

// 解法三
// 滑动窗口
func longestPalindrome3(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移动到相同字母的最右边（如果有相同字母）
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 找到回文的边界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置到下一次寻找回文的中心
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
