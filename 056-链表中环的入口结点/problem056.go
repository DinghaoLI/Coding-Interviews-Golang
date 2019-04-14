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
    if head == nil || head.Next == nil {
        return nil
    }
    slow := head.Next
    fast := head.Next.Next
    
    for fast != nil && fast.Next != nil {
        if slow == fast {
            break
        }
        slow, fast = slow.Next, fast.Next.Next
    }
    
    tmp := head
    // n = (q-1)m + (m-k)
    // 所以从slow和head同时开始走，当slow==head那说明，这就是entrance
    for tmp!=nil && slow!=nil {
        if slow == tmp {
            return slow
        }
        slow, tmp = slow.Next, tmp.Next
    }
    return nil
}

// 测试地址
// https://leetcode.com/problems/linked-list-cycle-ii/


