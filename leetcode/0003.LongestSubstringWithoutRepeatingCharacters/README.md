# [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 题目
Given a string s, find the length of the longest substring without repeating characters.


Example 1:
```bash
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

Example 2:
```bash
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example 3:

```bash
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.


## 题目大意
给定一个字符串s，求最长不重复字符子字符串的长度

## 解题思路
我的思路：
* 先对针对空字符串或者纯空格字符串的特殊情况单独处理
* 生成slice，遍历字符串，当slice为空时，追加元素，增加子循环遍历，当元素不在slice时，追加并且判断最大长度，否则跳出子循环。
