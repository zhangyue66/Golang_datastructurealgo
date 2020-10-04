func generateTheString(n int) string {
	res := ""
	
	if n % 2 == 0 {
		for i := 0 ; i < n-1 ; i ++ {
			res += "a"
		}
		res += "b"
	} else {
		for i := 0 ; i < n ;i ++ {
			res += "a"
		}
	}
	
	return res
	
}
