func simplifiedFractions(n int) []string {
	ans := []string{}
	if n == 1 {
		return ans
	}
	for i := 2; i <= n ; i ++ {
		for j := 1; j < i; j ++ {
			if gcd(j,i) == 1 {
				stringYZ := strconv.Itoa(j) + "/" + strconv.Itoa(i)
				ans = append(ans,stringYZ)
			}
		}
	}

	return ans
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	for b != 0 {
		a,b = b,a%b
	} 
	return a

}
