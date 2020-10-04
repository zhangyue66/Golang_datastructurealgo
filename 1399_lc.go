func countLargestGroup(n int) int {
	res := [][]int{}
	set := make(map[int]bool)

	for i := 1; i <= n; i++ {
		yz := addNumber(i)
		_,ok := set[yz]
		if !ok {
			set[yz] = true
            res = append(res,[]int{yz})
		} else {
			res[yz-1] = append(res[yz-1],yz)
		}
	}
    //fmt.Println(res)
	var max_len float64
	for _,rs := range res {
        max_len = math.Max(float64(max_len),float64(len(rs)))
	}
	ans := 0
	for _,rs := range res {
        if len(rs) == int(max_len) {
			ans++
		}
	}
	return ans
}

func addNumber(n int) int {
	str := strconv.Itoa(n)

	cnt := 0

	for _,char := range str {
        temp,_ := strconv.Atoi(string(char))
		 cnt += temp
	}
    //fmt.Println(cnt)
	return cnt

}
