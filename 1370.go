func sortString(s string) string {
    dict := map[string] int{}
	
	letters := []string{}

    for _,ru := range s {
        _,ok := dict[string(ru)]
		if !ok {
            letters = append(letters,string(ru))
		}
        dict[string(ru)] +=1

	}
    sort.Strings(letters)
    
    temp := reverse(letters)
    //fmt.Println(letters)
    letters = append(letters,temp...)
    
    ans := ""
    change := true
    for change {
        for _,letter := range letters {
            if dict[letter] > 0 {
                ans += letter
                dict[letter] -= 1
            }
            if len(ans) == len(s) {
                change = false
                
            }
            if change == false {
                //fmt.Println(dict,letters)
                break
            }
        }

    }
    
    return ans

}

func reverse(s []string)  []string{
    ans := make([]string,len(s))
    for i,j := 0,len(s)-1; i < len(s) ; i,j = i+1,j-1 {
        ans[j] = string(s[i])
    }
    return ans
}
