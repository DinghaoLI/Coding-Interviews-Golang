# 题意
题目描述

求1+2+3+...+n，

要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

样例输入

3

5

样例输出

6

15
 
# 递归+短路判断终止
计算1+2+3+...+n, 可以认为是一个递归的过程, 这点很容易理解

但是怎么不用分支判断来保证递归的终止呢

通过短路运算0&&cout使条件为假值,从而不执行输出语句，那么我们也可以通过短路来实现循环终止,从n开始递减进程递归的相加运算当递归至0时使递归短路即可。

代码如下
```java
class Solution
{
public:
    int SumRecursion(int n)
    {
        int ans = n;
        //debug <<n <<endl;
        n && (ans += SumRecursion(n - 1));
        return ans;
    }
};
```
其实也可以用下面的语句

ans && (ans += Sum_Solution(n - 1));

ans每次均被初始化为n, 因此用n或者ans进行短路运算都可以，其本质就是通过n进行短路运算判断

https://github.com/gatieme/CodingInterviews/blob/master/046-%E6%B1%821%2B2%2B3%2B...%2Bn/README.md
