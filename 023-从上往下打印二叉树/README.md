# 题意

从上往下打印出二叉树的每个节点，同层节点从左至右打印。

# 分析

要不然模仿[leetCode 102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

要不然维护一个队列，把同级左右字数添加到对列，然后每次都从头取元素来操作
