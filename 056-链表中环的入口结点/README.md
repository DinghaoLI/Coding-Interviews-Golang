# 题目描述

一个链表中包含环，请找出该链表的环的入口结点。

[LeetCode 142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii)

# 双指针法

受到第15题的启发剑指Offer--015-链表中倒数第k个结点, 我们考虑这样一个事实

假设链表长度为N, 那么第N链接到了第k个节点形成了环，即我们需要查找到倒数第N-K+1个节点, 那么环中就有N-K+1个节点，这时候我们定义两个指针$P_1$和$P_2$指向链表的头部, 指针$P_1$先在链表中向前移动n-k+1步,到达第n-k+2个节点, 然后两个指针同步向前移动, 当$P_2$走了K-1步到达环的入口的时候, 指针$P_1$正好走了N+1步, 到达了环的入口, 即两个指针会相遇在环的入口处

那么我们剩下的问题就是如何得到环中节点的数目?

我们可以使用一快一慢两个指针（比如慢指针一次走一步，　慢指针一次走两步），如果走的过程中发现快指针追上了慢指针, 说明遇见了环，而且相遇的位置一定在环内, 考虑一下环内, 从任何一个节点出现再回到这个节点的距离就是环的长度, 于是我们可以进一步移动慢指针，快指针原地不动, 当慢指针再次回到相遇位置时, 正好在环内走了一圈, 从而我们通过计数就可以获取到环的长度

第一步，找环中相汇点。分别用p1，p2指向链表头部，p1每次走一步，p2每次走二步，直到p1==p2找到在环中的相汇点。

第二步，找环的长度。从环中的相汇点开始, p2不动, p1前移，　当再次相遇时，p1刚好绕环一周, 其移动即为环的长度K

第三步, 求换的起点, 转换为求环的倒数第N-K个节点，则两指针left和right均指向起始, right先走K步, 然后两个指针开始同步移动, 当两个指针再次相遇时, right刚好绕环一周回到起点, left则刚好走到了起点位置

```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://juejin.im/post/59e5544851882551dd311710
// n为head到entrance的节点数
// m为环的长度
// k为slow过了entrance后走的节点数
// m-k为从fast、slow的相遇点到entrance的节点数(向前走)
// slow = n + k
// fast = qm + n + k  (q为正整数，为圈数)
// fast = 2 * slow
// 可得 n = qm -k
// 亦是 n = (q-1)m + (m-k)
// 证毕
func detectCycle(head *ListNode) *ListNode {
  slow, fast := head, head
  if head == nil || head.Next == nil {
    return nil
  }
  
  for fast.Next != nil && fast.Next.Next != nil {
    fast = fast.Next.Next
    slow = slow.Next
    if slow == fast {
      tmp := head
      // n = (q-1)m + (m-k)
      // 所以从slow和head同时开始走，当slow==head那说明，这就是entrance
      for tmp != slow {
        slow = slow.Next
        tmp = tmp.Next
      }
      return tmp
    }
  }
  
  return nil
}
```