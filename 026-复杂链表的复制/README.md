# 题意

题目描述

输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针指向任意一个节点）。

要求你编写函数复制这个复杂链表

## 分析

利用hashmap
```golang
m := make(map[*RandNodeList]*RandNodeList)
```
只需要遍历一次List，新旧两边的Node按照Next的顺序在hashmap中对应，遍历一次，map中有就链接，没有就创建后在链接。 