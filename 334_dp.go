func increasingTriplet(nums []int) bool {
    if len(nums) <= 2 {
        return false
    }
    
    dp := make([]int,len(nums))
    
    for i:=0 ; i < len(dp) ; i ++ {
        dp[i] = 1
    }
    
    for i := 1 ; i < len(nums) ; i ++ {
        for j := 0 ; j <= i ; j ++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i],dp[j] + 1)
            }
        }
    }
    
    for _,value := range dp {
        if value >= 3 {
            return true
        }
    }
    return false
}

func max(a,b int) int {
    if a < b {
        return b
    }
    return a
}
