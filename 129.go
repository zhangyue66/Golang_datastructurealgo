/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    var res *int
    res = new(int)
    *res = 0
    dfs(root,0,res)
    //fmt.Println(&res)
    return *res
    
    
}

func dfs(node *TreeNode, prev int,res *int)  {
    if node == nil{
        return 
    }
    if node.Left == nil && node.Right == nil {
        *res += (prev*10 + node.Val)
        //fmt.Println(&res)
        return
    }
    prev = 10*prev + node.Val
    dfs(node.Left,prev,res)
    dfs(node.Right,prev,res)
    
    return 
}
