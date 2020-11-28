func isValid(s string) bool {
    stack := []rune{}
    for _,char := range s {
        stack = append(stack,char)
        
        
        if len(stack) >= 3 {
            a := stack[len(stack)-3:]
            b  := []rune{97,98,99}
            if  compare(a,b) == true {
                stack = stack[0:len(stack)-3]
            }
        }
    }
    if len(stack) == 0{
        return true
    }
    return false
}

func compare(a []rune, b []rune) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
