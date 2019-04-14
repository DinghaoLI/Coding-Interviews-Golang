
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var isMirror func(*TreeNode,*TreeNode) bool
    isMirror = func(t1, t2 *TreeNode) bool {
        // 判断两个树是否都存在
        if t1 == nil && t2 == nil {
            return true
        }
        if t1 == nil || t2 == nil {
            return false
        }
        // t1和t2得相等，同时t1.Left和t2.Right， t1.Right和t2.Left 都需要对称
        return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
    }
    return isMirror(root.Left, root.Right)
}

// 测试地址
// https://leetcode.com/problems/symmetric-tree/