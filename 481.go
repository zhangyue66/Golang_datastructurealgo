func magicalString(n int) int {
    s := []int{1,2,2}
    if n == 0 {
        return 0
    }
    if n <= 3 {
        return 1
    }
    idx := 2
    
    for len(s) < n {
        num := 3-s[len(s)-1]
        
        for i := 0; i < s[idx] ; i ++ {
            s = append(s,num)
        }
        idx += 1
        
    }
    ans := 0
    
    for j := 0 ; j < n ; j++ {
        if s[j] == 1 {
            ans += 1
        }
    }
    
    return ans
}
