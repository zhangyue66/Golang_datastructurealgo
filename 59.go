func generateMatrix(n int) [][]int {
    
    ans := make([][]int,n)
    for i := range ans {
        ans[i] = make([]int,n)
    }
    
    dr := [4]int{0,1,0,-1}
    dc := [4]int{1,0,-1,0}
    
    direction := 0
    
    i,j := 0,0
    
    for num := 1; num <= n*n ; num ++ {
        ans[i][j] = num
        
        ni,nj := i+dr[direction] , j + dc[direction]
        
        if (ni >= 0 && ni < n) && ( nj >= 0 && nj < n) && ( ans[ni][nj] == 0) {
            i,j = ni,nj
        } else {
            direction = (direction + 1) % 4
            i += dr[direction]
            j += dc[direction]
        }
        
        
    }
    
    return ans
}
