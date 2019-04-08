package main

import (
	"fmt"
)

// 题目：一个整型数组里除了一个数字之外，其他的数字都出现了两次
// 证明：我们把每个比特分开来看，有题目可知，每个元素出现两次，只有一个元素出现一次，那么，我们只需要保留每一位比特的出现单数次的符号就可以了
// 比如: [4, 1, 2, 1, 2] 用二进制表示
// 0 1 0 0
// 0 0 0 1
// 0 0 1 0
// 0 0 0 1
// 0 0 1 0
// 所以 第一位0出现单数次，我们取0
// 第二位， 1 出现单数次， 我们取1
// 第三位，0 出现单数次，我们取0
// 第四位，0 出现单数次，我们取0
// 所以 结果位 0100， 以上的操作用xor按顺序运行过来就可以。
// 测试地址
// https://leetcode.com/problems/single-number/
func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		// n^n == 0
		// a^b^a^b^a == a
		res ^= n
	}
	return res
}


// 一个整型数组里除了一个数字之外，其他的数字都出现了三次
// 如果是出现两次的话，用一个 bit 就可以
// 比如 ones，初始为0
// 当 5 第 1 次出现， ones = 5
// 当 5 第 2 次出现， ones 清空为 0，表示 ones 可以去处理其他数字了
// 所以，最后 如果 ones != 0的话， ones 记录的就是只出现了一次的那个数字

// 公式是 ones = ones xor i

// 那么，如果是三次的话，香农定理，需要用 2 bits 进行记录

// 当 5 第 1 次出现的时候，ones = 5, twos = 0,  ones 记录这个数字
// 当 5 第 2 次出现的时候，ones = 0, twos = 5， twos 记录了这个数字
// 当 5 第 3 次出现的时候，ones = 0, twos = 0， 都清空了，可以去处理其他数字了
// 所以，如果有某个数字出现了 1 次，就存在 ones 中，出现了 2 次，就存在 twos 中

// 公式方面， 上面 2 次的时候，ones 清空的公式是 ones = ones xor i
// 而第 3 次时， ones 要等于零， 而这时 twos 是 True ， 所以再 & 一个 twos 的非就可以， ones = (ones xor i) & ~twos
// 所以，总的公式是
//    ones = (ones xor i) & ~twos
//    twos = (twos xor i) & ~ones
// 测试地址
// https://leetcode.com/problems/single-number-ii/
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

// 一个整型数组里除了两个数字之外，其他的数字都出现了两次
// 此题考察的是异或运算的特点：即两个相同的数异或结果为0。
// 此题用了两次异或运算特点：
// 第一次使用异或运算，得到了两个只出现一次的数相异或的结果。
// 因为两个只出现一次的数肯定不同，即他们的异或结果一定不为0，一定有一个位上有1。
// 另外一个此位上没有1，我们可以根据此位上是否有1，将整个数组重新划分成两部分，
// 一部分此位上一定有1，另一部分此位上一定没有1，
// 然后分别对每部分求异或，因为划分后的两部分有这样的特点：
// 其他数都出现两次，只有一个数只出现一次。
// 因此，我们又可以运用异或运算，分别得到两部分只出现一次的数。

func singleNumber(nums []int) []int {
    xor := 0
    for _,v := range nums {
        xor ^= v
    }
    
    // 取xor最低位为1的数
    lowest := xor & -xor
    
    a,b := 0,0
    for _, v := range nums {
        if lowest&v == 0 {
            a ^= v
        } else {
            b ^= v
        }
    }
    return []int{a,b}
}
// 测试地址
// https://leetcode.com/problems/single-number-iii/
