/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSubPath(head *ListNode, root *TreeNode) bool {
	 if head == nil {
		 return true
	 }
	 if root == nil {
		 return false
	 }

	 condition1 := dfs(head,root)
	 condition2 := isSubPath(head,root.Left)
	 condition3 := isSubPath(head,root.Right)
	 
	 return condition1 || condition2 || condition3
}
func dfs(head *ListNode, root *TreeNode) bool {
	if head == nil {
        return true
	}
	if root == nil {
		return false
	}

	if root.Val == head.Val {
        a := dfs(head.Next,root.Left)
		b := dfs(head.Next,root.Right)
		return a || b
	}
	return false
}
