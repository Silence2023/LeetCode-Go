# [AddTwoNumbers](https://leetcode.com/problems/add-two-numbers/)

## 题目
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:
```bash
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

Example 2:

```bash
Input: l1 = [0], l2 = [0]
Output: [0]
```

Example 3:
```bash
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

## 题目大意
2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点。


## 解题思路
生成一个初始节点的结构体，Val为0，用于最后的结果返回
遍历两个节点，当其中一个节点为空时，遍历另外一个节点时，可以提前判断有没有进位，没有的话可以直接将Next指向非空节点即可。
