# 剑指Offer - Golang实现

## 简介

一直没找到完整版的剑指Offer的Golang实现，所以索性自己来写。

每个文件夹对应每个题目，文件夹内的README.md内有题目和必要的分析，Problem\*.go 是和算法相对应的代码。

如果大家使用的是Chrome，给推荐大家一个小插件：**[octotree](https://chrome.google.com/webstore/detail/octotree/bkhaagjahfmjljalopjnoealnfndnagc?hl=zh-CN)**，这是一个极其方便的github Repo索引小插件，真的很方便。

## 运行代码

若题目对应包内有 \*\_test.go 文件，可直接执行go test进行单元测试
```shell
go test
```

若没有test文件，请直接执行
```shell
go run problemXXX.go
```

**如果文档里有说明该题目是LeetCode上原题，请点击内附的LeetCode题目链接，并复制代码到LeetCode上执行。**

## 如有错误或者更好的算法版本，欢迎各种PR～


| 题目      | 
| ---------- | 
| [003-二维数组中的查找](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/003-%E4%BA%8C%E7%BB%B4%E6%95%B0%E7%BB%84%E4%B8%AD%E7%9A%84%E6%9F%A5%E6%89%BE) |
[004-替换空格](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/004-%e6%9b%bf%e6%8d%a2%e7%a9%ba%e6%a0%bc) |
| [005-从尾到头打印链表](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/005-%e4%bb%8e%e5%b0%be%e5%88%b0%e5%a4%b4%e6%89%93%e5%8d%b0%e9%93%be%e8%a1%a8) |
| [006-重建二叉树](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/006-%e9%87%8d%e5%bb%ba%e4%ba%8c%e5%8f%89%e6%a0%91) |
| [007-用两个栈实现队列](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/007-%e7%94%a8%e4%b8%a4%e4%b8%aa%e6%a0%88%e5%ae%9e%e7%8e%b0%e9%98%9f%e5%88%97) |
| [008-旋转数组的最小数字](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/008-%e6%97%8b%e8%bd%ac%e6%95%b0%e7%bb%84%e7%9a%84%e6%9c%80%e5%b0%8f%e6%95%b0%e5%ad%97) |
| [009-斐波那契数列](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/009-%e6%96%90%e6%b3%a2%e9%82%a3%e5%a5%91%e6%95%b0%e5%88%97) |
| [010-二进制中1的个数](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/010-%e4%ba%8c%e8%bf%9b%e5%88%b6%e4%b8%ad1%e7%9a%84%e4%b8%aa%e6%95%b0) |
| [011-数值的整数次方](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/011-%e6%95%b0%e5%80%bc%e7%9a%84%e6%95%b4%e6%95%b0%e6%ac%a1%e6%96%b9) |
| [012-打印1到最大的N位数](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/012-%e6%89%93%e5%8d%b01%e5%88%b0%e6%9c%80%e5%a4%a7%e7%9a%84N%e4%bd%8d%e6%95%b0) |
| [014-调整数组顺序使奇数位于偶数前面](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/014-%e8%b0%83%e6%95%b4%e6%95%b0%e7%bb%84%e9%a1%ba%e5%ba%8f%e4%bd%bf%e5%a5%87%e6%95%b0%e4%bd%8d%e4%ba%8e%e5%81%b6%e6%95%b0%e5%89%8d%e9%9d%a2) |
| [015-链表中倒数第k个结点](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/015-%e9%93%be%e8%a1%a8%e4%b8%ad%e5%80%92%e6%95%b0%e7%ac%ack%e4%b8%aa%e7%bb%93%e7%82%b9) |
| [016-反转链表](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/016-%e5%8f%8d%e8%bd%ac%e9%93%be%e8%a1%a8) |
| [017-合并两个排序的链表](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/017-%e5%90%88%e5%b9%b6%e4%b8%a4%e4%b8%aa%e6%8e%92%e5%ba%8f%e7%9a%84%e9%93%be%e8%a1%a8) |
| [018-树的子结构](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/018-%e6%a0%91%e7%9a%84%e5%ad%90%e7%bb%93%e6%9e%84) |
| [019-二叉树的镜像](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/019-%e4%ba%8c%e5%8f%89%e6%a0%91%e7%9a%84%e9%95%9c%e5%83%8f) |
| [020-顺时针打印矩阵](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/020-%e9%a1%ba%e6%97%b6%e9%92%88%e6%89%93%e5%8d%b0%e7%9f%a9%e9%98%b5) |
| [021-包含min函数的栈](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/021-%e5%8c%85%e5%90%abmin%e5%87%bd%e6%95%b0%e7%9a%84%e6%a0%88) |
| [022-栈的压入弹出序列](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/022-%e6%a0%88%e7%9a%84%e5%8e%8b%e5%85%a5%e5%bc%b9%e5%87%ba%e5%ba%8f%e5%88%97) |
| [023-从上往下打印二叉树](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/023-%e4%bb%8e%e4%b8%8a%e5%be%80%e4%b8%8b%e6%89%93%e5%8d%b0%e4%ba%8c%e5%8f%89%e6%a0%91) |
| [024-二叉搜索树的后序遍历序列](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/024-%e4%ba%8c%e5%8f%89%e6%90%9c%e7%b4%a2%e6%a0%91%e7%9a%84%e5%90%8e%e5%ba%8f%e9%81%8d%e5%8e%86%e5%ba%8f%e5%88%97) |
| [025-二叉树中和为某一值的路径](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/025-%e4%ba%8c%e5%8f%89%e6%a0%91%e4%b8%ad%e5%92%8c%e4%b8%ba%e6%9f%90%e4%b8%80%e5%80%bc%e7%9a%84%e8%b7%af%e5%be%84) |
| [026-复杂链表的复制](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/026-%e5%a4%8d%e6%9d%82%e9%93%be%e8%a1%a8%e7%9a%84%e5%a4%8d%e5%88%b6) |
| [027-二叉搜索树与双向链表](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/027-%e4%ba%8c%e5%8f%89%e6%90%9c%e7%b4%a2%e6%a0%91%e4%b8%8e%e5%8f%8c%e5%90%91%e9%93%be%e8%a1%a8) |
| [028-字符串的排列](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/028-%e5%ad%97%e7%ac%a6%e4%b8%b2%e7%9a%84%e6%8e%92%e5%88%97) |
| [029-数组中出现次数超过一半的数字](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/029-%e6%95%b0%e7%bb%84%e4%b8%ad%e5%87%ba%e7%8e%b0%e6%ac%a1%e6%95%b0%e8%b6%85%e8%bf%87%e4%b8%80%e5%8d%8a%e7%9a%84%e6%95%b0%e5%ad%97) |
| [030-最小的K个数](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/030-%e6%9c%80%e5%b0%8f%e7%9a%84K%e4%b8%aa%e6%95%b0) |
| [031-连续子数组的最大和](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/031-%e8%bf%9e%e7%bb%ad%e5%ad%90%e6%95%b0%e7%bb%84%e7%9a%84%e6%9c%80%e5%a4%a7%e5%92%8c) |
| [032-从1到n整数中1出现的次数](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/032-%e4%bb%8e1%e5%88%b0n%e6%95%b4%e6%95%b0%e4%b8%ad1%e5%87%ba%e7%8e%b0%e7%9a%84%e6%ac%a1%e6%95%b0) |
| [033-把数组排成最小的数](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/033-%e6%8a%8a%e6%95%b0%e7%bb%84%e6%8e%92%e6%88%90%e6%9c%80%e5%b0%8f%e7%9a%84%e6%95%b0) |
| [034-丑数](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/034-%e4%b8%91%e6%95%b0) |
| [035-第一个只出现一次的字符位置](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/035-%e7%ac%ac%e4%b8%80%e4%b8%aa%e5%8f%aa%e5%87%ba%e7%8e%b0%e4%b8%80%e6%ac%a1%e7%9a%84%e5%ad%97%e7%ac%a6%e4%bd%8d%e7%bd%ae) |
| [036-数组中的逆序对](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/036-%e6%95%b0%e7%bb%84%e4%b8%ad%e7%9a%84%e9%80%86%e5%ba%8f%e5%af%b9) |
| [037-两个链表的第一个公共结点](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/037-%e4%b8%a4%e4%b8%aa%e9%93%be%e8%a1%a8%e7%9a%84%e7%ac%ac%e4%b8%80%e4%b8%aa%e5%85%ac%e5%85%b1%e7%bb%93%e7%82%b9) |
| [038-数字在排序数组中出现的次数](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/038-%e6%95%b0%e5%ad%97%e5%9c%a8%e6%8e%92%e5%ba%8f%e6%95%b0%e7%bb%84%e4%b8%ad%e5%87%ba%e7%8e%b0%e7%9a%84%e6%ac%a1%e6%95%b0) |
| [039-二叉树的深度](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/039-%e4%ba%8c%e5%8f%89%e6%a0%91%e7%9a%84%e6%b7%b1%e5%ba%a6) |
| [039-平衡二叉树[附加]](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/039-%e5%b9%b3%e8%a1%a1%e4%ba%8c%e5%8f%89%e6%a0%91%5b%e9%99%84%e5%8a%a0%5d) |
| [040-数组中只出现一次的数字](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/040-%e6%95%b0%e7%bb%84%e4%b8%ad%e5%8f%aa%e5%87%ba%e7%8e%b0%e4%b8%80%e6%ac%a1%e7%9a%84%e6%95%b0%e5%ad%97) |
| [041-和为S的两个数字](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/041-%e5%92%8c%e4%b8%baS%e7%9a%84%e4%b8%a4%e4%b8%aa%e6%95%b0%e5%ad%97) |
| [041-和为S的连续正数序列](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/041-%e5%92%8c%e4%b8%baS%e7%9a%84%e8%bf%9e%e7%bb%ad%e6%ad%a3%e6%95%b0%e5%ba%8f%e5%88%97) |
| [042-左旋转字符串](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/042-%e5%b7%a6%e6%97%8b%e8%bd%ac%e5%ad%97%e7%ac%a6%e4%b8%b2) |
| [042-翻转单词顺序列](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/042-%e7%bf%bb%e8%bd%ac%e5%8d%95%e8%af%8d%e9%a1%ba%e5%ba%8f%e5%88%97) |
| [044-扑克牌顺子](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/044-%e6%89%91%e5%85%8b%e7%89%8c%e9%a1%ba%e5%ad%90) |
| [045-孩子们的游戏(圆圈中最后剩下的数)](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/045-%e5%ad%a9%e5%ad%90%e4%bb%ac%e7%9a%84%e6%b8%b8%e6%88%8f(%e5%9c%86%e5%9c%88%e4%b8%ad%e6%9c%80%e5%90%8e%e5%89%a9%e4%b8%8b%e7%9a%84%e6%95%b0)) |
| [046-求1+2+3+...+n](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/046-%e6%b1%821%2b2%2b3%2b...%2bn) |
| [047-不用加减乘除做加法](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/047-%e4%b8%8d%e7%94%a8%e5%8a%a0%e5%87%8f%e4%b9%98%e9%99%a4%e5%81%9a%e5%8a%a0%e6%b3%95) |
| [048-不能被继承的类](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/048-%e4%b8%8d%e8%83%bd%e8%a2%ab%e7%bb%a7%e6%89%bf%e7%9a%84%e7%b1%bb) |
| [049-把字符串转换成整数](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/049-%e6%8a%8a%e5%ad%97%e7%ac%a6%e4%b8%b2%e8%bd%ac%e6%8d%a2%e6%88%90%e6%95%b4%e6%95%b0) |
| [051-数组中重复的数字](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/051-%e6%95%b0%e7%bb%84%e4%b8%ad%e9%87%8d%e5%a4%8d%e7%9a%84%e6%95%b0%e5%ad%97) |
| [052-构建乘积数组](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/052-%e6%9e%84%e5%bb%ba%e4%b9%98%e7%a7%af%e6%95%b0%e7%bb%84) |
| [053-正则表达式匹配](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/053-%e6%ad%a3%e5%88%99%e8%a1%a8%e8%be%be%e5%bc%8f%e5%8c%b9%e9%85%8d) |
| [054-表示数值的字符串](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/054-%e8%a1%a8%e7%a4%ba%e6%95%b0%e5%80%bc%e7%9a%84%e5%ad%97%e7%ac%a6%e4%b8%b2) |
| [055-字符流中第一个不重复的字符](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/055-%e5%ad%97%e7%ac%a6%e6%b5%81%e4%b8%ad%e7%ac%ac%e4%b8%80%e4%b8%aa%e4%b8%8d%e9%87%8d%e5%a4%8d%e7%9a%84%e5%ad%97%e7%ac%a6) |
| [056-链表中环的入口结点](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/056-%e9%93%be%e8%a1%a8%e4%b8%ad%e7%8e%af%e7%9a%84%e5%85%a5%e5%8f%a3%e7%bb%93%e7%82%b9) |
| [057-删除链表中重复的结点](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/057-%e5%88%a0%e9%99%a4%e9%93%be%e8%a1%a8%e4%b8%ad%e9%87%8d%e5%a4%8d%e7%9a%84%e7%bb%93%e7%82%b9) |
| [058-二叉树的下一个结点](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/058-%e4%ba%8c%e5%8f%89%e6%a0%91%e7%9a%84%e4%b8%8b%e4%b8%80%e4%b8%aa%e7%bb%93%e7%82%b9) |
| [059-对称的二叉树](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/059-%e5%af%b9%e7%a7%b0%e7%9a%84%e4%ba%8c%e5%8f%89%e6%a0%91) |
| [060-把二叉树打印成多行](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/060-%e6%8a%8a%e4%ba%8c%e5%8f%89%e6%a0%91%e6%89%93%e5%8d%b0%e6%88%90%e5%a4%9a%e8%a1%8c) |
| [061-按之字形顺序打印二叉树](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/061-%e6%8c%89%e4%b9%8b%e5%ad%97%e5%bd%a2%e9%a1%ba%e5%ba%8f%e6%89%93%e5%8d%b0%e4%ba%8c%e5%8f%89%e6%a0%91) |
| [062-序列化二叉树](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/062-%e5%ba%8f%e5%88%97%e5%8c%96%e4%ba%8c%e5%8f%89%e6%a0%91) |
| [063-二叉搜索树的第K个结点](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/063-%e4%ba%8c%e5%8f%89%e6%90%9c%e7%b4%a2%e6%a0%91%e7%9a%84%e7%ac%acK%e4%b8%aa%e7%bb%93%e7%82%b9) |
| [064-数据流之中的中位数](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/064-%e6%95%b0%e6%8d%ae%e6%b5%81%e4%b9%8b%e4%b8%ad%e7%9a%84%e4%b8%ad%e4%bd%8d%e6%95%b0) |
| [065-滑动窗口的最大值](https://github.com/DinghaoLI/Coding-Interviews-Golang/tree/master/065-%e6%bb%91%e5%8a%a8%e7%aa%97%e5%8f%a3%e7%9a%84%e6%9c%80%e5%a4%a7%e5%80%bc) |

