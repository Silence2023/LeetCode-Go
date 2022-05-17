package leetcode

import "strings"

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
