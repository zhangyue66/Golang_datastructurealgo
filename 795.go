func numSubarrayBoundedMax(A []int, L int, R int) int {
    res,dp := 0,0
    prev := -1
    
    for i,_ := range A {
        if A[i] < L && i > 0 {
            res += dp
        }
        if A[i] > R {
            prev = i
            dp = 0
        }
        if A[i] >= L && A[i] <= R {
            dp = i - prev
            res += dp
        }
    }
    return res
    
}
