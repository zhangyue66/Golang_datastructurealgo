import "unicode"

func decodeAtIndex(S string, K int) string {
    size := 0
    
    reverse := ""
    
    
    for _,s := range S {
        reverse = string(s)+reverse
        
        if unicode.IsLetter(s) {
            size += 1
        } else {
            size *= int(s-'0')
        }
    }
    //fmt.Println(size,reverse,K) //10404 3edoc2teel
    
    for _,s := range reverse {
        //fmt.Println(string(s))
        K = K % size
        //fmt.Println(K)
        if K == 0 && unicode.IsLetter(s) {
            
            //fmt.Println(string(s))
            return string(s)
            
        }
        
        if !unicode.IsLetter(s) {
            
            size = size/int(s-'0')
            
        } else {
            size -= 1
        }
    }
    
    return ""
    
}
