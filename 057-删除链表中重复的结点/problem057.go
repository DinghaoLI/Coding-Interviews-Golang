// 对于每一部分处理方式相同，所以我我们考虑用递归
// 递归
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
 // 长度 <=1 的 list ，可以直接返回
 if head == nil || head.Next == nil {
     return head
 }

 // 要么 head 重复了，那就删除 head
 if head.Val == head.Next.Val {
     for head.Next != nil && head.Val == head.Next.Val {
         head = head.Next
     }
        // 有重复所以直接不带head
     return deleteDuplicates(head.Next)
 }

 // 要么 head 不重复，递归处理 head 后面的节点
 head.Next = deleteDuplicates(head.Next)
 return head
}

// 测试地址
// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/