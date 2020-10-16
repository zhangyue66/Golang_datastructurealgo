func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
    ans := make([][]int,R*C)
    ans[0] = []int{r0,c0}
    
    cnt := 1
    
    dr := []int{0,1,0,-1} // east ,south,west ,north
    dc := []int{1,0,-1,0}
        
    direction := 0
    step := 0

    
    for cnt != R*C {
       
        if direction == 0 || direction == 2 {
            step += 1
        }
        // fmt.Println(direction)
        // fmt.Println(step)

        for num := 0 ; num < step ; num ++ {
            //fmt.Println(r0,c0)
            r0 +=  dr[direction]
            c0 +=  dc[direction]
            
            
            if r0 >= 0 && r0 < R && c0 >= 0 && c0 < C {          
                
                ans[cnt] = []int{r0,c0}
                cnt ++
            }
            if cnt == R*C {
                break
            }       
        }
        direction = (direction + 1) % 4       
    }
    
    return ans
}
