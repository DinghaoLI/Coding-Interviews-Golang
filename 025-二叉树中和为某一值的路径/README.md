# 题意

题目描述

输入一颗二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径

# 分析

DFS

## 参考Leetcode原题

[113. Path Sum II](https://leetcode.com/problems/path-sum-ii/)

```golang
// 易错点：
// 语法问题！！！
// 公用solution的时候，记得递归前后恢复原样
// 添加到res时注意深拷贝数组
func pathSum(root *TreeNode, sum int) [][]int {//写个递归函数，从左到右把每条路线都试一下，有符合的就把路径加入paths
    var res [][]int
    solution := []int{}
    var dfs func(root *TreeNode, sum int)
    dfs = func(root *TreeNode, sum int){
        if root == nil { return }
        rest := sum - root.Val
        if rest == 0 && root.Left == nil && root.Right == nil {
            solution = append(solution, root.Val)
            tmp := make([]int, len(solution))
            copy(tmp, solution)
            res = append(res, tmp)
            solution = solution[:len(solution)-1]
            return
        }
        solution = append(solution, root.Val)
        dfs(root.Left, rest)
        dfs(root.Right, rest)
        solution = solution[:len(solution)-1]
    }
    dfs(root, sum)
    return res
}
```