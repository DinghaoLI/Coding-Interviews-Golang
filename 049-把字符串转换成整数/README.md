#题意

将一个字符串转换成一个整数，要求不能使用字符串转换整数的库函数。


[LeetCode 8. String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/)

```
Example 1:

Input: "42"
Output: 42
```
```
Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign. Then take as many numerical digits as possible, which gets 42.
```
```
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
```
```
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical digit or a +/- sign. Therefore no valid conversion could be performed.
```
