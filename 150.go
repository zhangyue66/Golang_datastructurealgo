func evalRPN(tokens []string) int {
    stack := []int{}
    for _,token := range tokens {
        if token == "+" {
            num2 := stack[len(stack)-1]
            num1 := stack[len(stack)-2]
            
            // num1,err = strconv.Atoi(num1)
            // num2,err = strconv.Atoi(num2)
            
            stack = stack[0:len(stack)-2]
            
            stack = append(stack,num1+num2)
            
        }else if token == "-" {
            num2 := stack[len(stack)-1]
            num1 := stack[len(stack)-2]
            stack = stack[0:len(stack)-2]
            stack = append(stack,num1-num2)
        } else if token == "*" {
            num2 := stack[len(stack)-1]
            num1 := stack[len(stack)-2]
            stack = stack[0:len(stack)-2]
            stack = append(stack,num1*num2)
        } else if token == "/" {
            num2 := stack[len(stack)-1]
            num1 := stack[len(stack)-2]
            stack = stack[0:len(stack)-2]
            stack = append(stack,int(num1/num2))
        } else {
            //fmt.Println(token)
            num,_ := strconv.Atoi(token)
            //fmt.Println(err)
            
            stack = append(stack,int(num))
        }
    }
    return stack[0]
}
