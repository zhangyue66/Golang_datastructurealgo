func findKthPositive(arr []int, k int) int {
	cnt := 0
    for i := 1;i<=arr[len(arr)-1];i++ {
		if contain(arr,i) == false {
			cnt++
		}
		if cnt == k {
			return i
		}
	}
    fmt.Println(cnt)
    return arr[len(arr)-1]+(k-cnt)

}

func contain(arr []int,target int) bool{
	for _,value := range arr {
		if value == target {
			return true
		}

	}
	return false
}
