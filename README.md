# LeetCode-Go

## 运行测试代码
```bash
root@DESKTOP-LE87MCN:/home/silence/LeetCode-Go# go test -v ./leetcode/...
=== RUN   Test_twoSum
------------------------Leetcode Problem TwoSum------------------------
【input】:{[3 2 4] 6}   【except】:[1 2]    【output】:[1 2]
【input】:{[3 2 4] 5}   【except】:[0 1]    【output】:[0 1]
【input】:{[0 8 7 3 3 4 2] 11}   【except】:[1 3]    【output】:[1 3]
【input】:{[0 1] 1}   【except】:[0 1]    【output】:[0 1]
【input】:{[0 3] 5}   【except】:[]    【output】:[]
--- PASS: Test_twoSum (0.00s)
PASS
ok      Leetcode-GO/leetcode/0001.TwoSum        (cached)
=== RUN   Test_AddTwoNumbers
------------------------Leetcode Problem AddTwoNumbers------------------------
【input】:{[] []}    【except】:[]   【output】:[]
【input】:{[1] [1]}    【except】:[2]   【output】:[2]
【input】:{[1 2 3 4] [1 2 3 4]}    【except】:[2 4 6 8]   【output】:[2 4 6 8]
【input】:{[1 2 3 4 5] [1 2 3 4 5]}    【except】:[2 4 6 8 0 1]   【output】:[2 4 6 8 0 1]
【input】:{[1] [9 9 9 9 9]}    【except】:[0 0 0 0 0 1]   【output】:[0 0 0 0 0 1]
【input】:{[9 9 9 9 9] [1]}    【except】:[0 0 0 0 0 1]   【output】:[0 0 0 0 0 1]
【input】:{[2 4 3] [5 6 4]}    【except】:[7 0 8]   【output】:[7 0 8]
【input】:{[1 8 3] [7 1]}    【except】:[8 9 3]   【output】:[8 9 3]
【input】:{[0] [7 3]}    【except】:[7 3]   【output】:[7 3]
【input】:{[9 1 6] [0]}    【except】:[9 1 6]   【output】:[9 1 6]
【input】:{[9 9 9 9 9 9 9] [9 9 9 9]}    【except】:[8 9 9 9 0 0 0 1]   【output】:[8 9 9 9 0 0 0 1]
--- PASS: Test_AddTwoNumbers (0.00s)
PASS
ok      Leetcode-GO/leetcode/0002.AddTwoNumbers 0.003s
=== RUN   Test_LongestSubstringWithoutRepeatingCharacters
------------------------Leetcode Test_LongestSubstringWithoutRepeatingCharacters------------------------
【input】:{abcabcbb}   【except】:3    【output】:3
【input】:{bbbbb}   【except】:1    【output】:1
【input】:{pwwkew}   【except】:3    【output】:2
【input】:{}   【except】:0    【output】:0
--- PASS: Test_LongestSubstringWithoutRepeatingCharacters (0.00s)
PASS
ok      Leetcode-GO/leetcode/0003.LongestSubstringWithoutRepeatingCharacters    0.002s
root@DESKTOP-LE87MCN:/home/silence/LeetCode-Go#
root@DESKTOP-LE87MCN:/home/silence/LeetCode-Go#
root@DESKTOP-LE87MCN:/home/silence/LeetCode-Go# go test -v ./leetcode/...
=== RUN   Test_twoSum
------------------------Leetcode Problem TwoSum------------------------
【input】:{[3 2 4] 6}   【except】:[1 2]    【output】:[1 2]
【input】:{[3 2 4] 5}   【except】:[0 1]    【output】:[0 1]
【input】:{[0 8 7 3 3 4 2] 11}   【except】:[1 3]    【output】:[1 3]
【input】:{[0 1] 1}   【except】:[0 1]    【output】:[0 1]
【input】:{[0 3] 5}   【except】:[]    【output】:[]
--- PASS: Test_twoSum (0.00s)
PASS
ok      Leetcode-GO/leetcode/0001.TwoSum        (cached)
=== RUN   Test_AddTwoNumbers
------------------------Leetcode Problem AddTwoNumbers------------------------
【input】:{[] []}    【except】:[]   【output】:[]
【input】:{[1] [1]}    【except】:[2]   【output】:[2]
【input】:{[1 2 3 4] [1 2 3 4]}    【except】:[2 4 6 8]   【output】:[2 4 6 8]
【input】:{[1 2 3 4 5] [1 2 3 4 5]}    【except】:[2 4 6 8 0 1]   【output】:[2 4 6 8 0 1]
【input】:{[1] [9 9 9 9 9]}    【except】:[0 0 0 0 0 1]   【output】:[0 0 0 0 0 1]
【input】:{[9 9 9 9 9] [1]}    【except】:[0 0 0 0 0 1]   【output】:[0 0 0 0 0 1]
【input】:{[2 4 3] [5 6 4]}    【except】:[7 0 8]   【output】:[7 0 8]
【input】:{[1 8 3] [7 1]}    【except】:[8 9 3]   【output】:[8 9 3]
【input】:{[0] [7 3]}    【except】:[7 3]   【output】:[7 3]
【input】:{[9 1 6] [0]}    【except】:[9 1 6]   【output】:[9 1 6]
【input】:{[9 9 9 9 9 9 9] [9 9 9 9]}    【except】:[8 9 9 9 0 0 0 1]   【output】:[8 9 9 9 0 0 0 1]
--- PASS: Test_AddTwoNumbers (0.00s)
PASS
ok      Leetcode-GO/leetcode/0002.AddTwoNumbers (cached)
=== RUN   Test_LongestSubstringWithoutRepeatingCharacters
------------------------Leetcode Test_LongestSubstringWithoutRepeatingCharacters------------------------
【input】:{abcabcbb}   【except】:3    【output】:3
【input】:{bbbbb}   【except】:1    【output】:1
【input】:{pwwkew}   【except】:3    【output】:2
【input】:{}   【except】:0    【output】:0
```
