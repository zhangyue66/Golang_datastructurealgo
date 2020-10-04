import "sort"

func canBeEqual1(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)

    for i := 0 ; i < len(target); i ++ {
        if target[i] != arr[i]{
            return false
        }
           
    }
    
    return true
}

func canBeEqual(target []int, arr []int) bool {
	dict1 := make(map[int]int)
	dict2 := make(map[int]int)

	for _,number := range target {
		_,ok := dict1[number]
		if !ok {
			dict1[number] = 1
		} else {
			dict1[number]++
		}
	}

	for _,num := range arr {
		_,ok := dict2[num]
		if !ok {
			dict2[num] = 1
		} else {
			dict2[num] ++
		}
	}


	key,value := range dict1 {
		if dict2[key] != value {
			return false
		}
	}

	return true
}
