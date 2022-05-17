package leetcode

import "strings"

// 自己的写法（太稚嫩了）
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	a := strings.TrimSpace(s)
	if len(a) == 0 {
		return 1
	}
	slice := make([]byte, 0)
	slice_byte := []byte(s)
	max_length := 0
	for i := 0; i < len(slice_byte); i++ {
		length := 0
		if len(slice) == 0 {
			slice = append(slice, slice_byte[i])
			length++
			if length > max_length {
				max_length = length
			}
		}
		for j := i + 1; j < len(slice_byte); j++ {
			result := 0
			for _, item := range slice {
				if slice_byte[j] == item {
					result += 1
					// 此时可以结束循环
					if length > max_length {
						max_length = length
					}
					break
				}
			}
			if result == 0 {
				slice = append(slice, slice_byte[j])
				length++
				if length > max_length {
					max_length = length
				}
			} else {
				slice = make([]byte, 0)
				break
			}
		}
	}
	return max_length
}

// 解法2：位图
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	result, left, right := 0, 0, 0
	var bitSet [256]bool
	for left < len(s) {
		if bitSet[s[right]] {
			// 说明这个字符已经存在过了，出现了重复，left需要向右移动后开始循环
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		// 如果剩下的已经小于当前的最大长度了或者已经遍历完了，跳出循环
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 解法三：滑动窗口
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
