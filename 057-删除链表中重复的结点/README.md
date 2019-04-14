# 题意

在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。

例如，链表1->2->3->3->4->4->5

处理后为 1->2->5

# 递归大法好

我们的思路是:

每次返回没有重复的head，如果有重复删除头节点，删完了，就用相同逻辑处理Next节点

[LeetCode 82. Remove Duplicates from Sorted List II](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/)

