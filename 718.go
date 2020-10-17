import "math"

func findLength(A []int, B []int) int {
	res := 0.0

	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}

	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1

			}

			res = math.Max(res, float64(dp[i][j]))
		}
	}

	return int(res)

}
