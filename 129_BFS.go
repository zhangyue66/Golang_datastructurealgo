func sumNumbers(root *TreeNode) int {
    res := 0
    if root == nil {
        return res
    }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        
        if cur.Left == nil && cur.Right == nil {
            res += cur.Val
        }
        
        if cur.Left != nil {
            cur.Left.Val += cur.Val*10
            queue = append(queue,cur.Left)
            
        }
        
        if cur.Right != nil {
            cur.Right.Val += cur.Val*10
            queue = append(queue,cur.Right)
        }
    }
    return res
}
