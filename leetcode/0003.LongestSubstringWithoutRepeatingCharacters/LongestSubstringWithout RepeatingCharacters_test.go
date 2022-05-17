package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3 struct {
	para3
	ans3
}

// para 是参数
// one 代表第一个参数
type para3 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans3 struct {
	one int
}

func Test_LongestSubstringWithoutRepeatingCharacters(t *testing.T) {

	qs := []question3{

		{
			para3{"abcabcbb"},
			ans3{3},
		},

		{
			para3{"bbbbb"},
			ans3{1},
		},

		{
			para3{"pwwkew"},
			ans3{3},
		},

		{
			para3{""},
			ans3{0},
		},

		{
			para3{" "},
			ans3{1},
		},

		{
			para3{"   "},
			ans3{1},
		},

		{
			para3{"c"},
			ans3{1},
		},

		{
			para3{"au"},
			ans3{2},
		},
	}
	fmt.Printf("------------------------Leetcode Test_LongestSubstringWithoutRepeatingCharacters------------------------\n")

	for _, q := range qs {
		except, p := q.ans3, q.para3
		result := lengthOfLongestSubstring(p.s)
		fmt.Printf("【input】:%v   【except】:%v    【output】:%v\n", p, except.one, result)
		assert.Equal(t, except.one, result)
	}
}
