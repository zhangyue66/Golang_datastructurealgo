//graph := make(map[int][]int])
var graph = map[int][]int{}

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
    graph = make(map[int][]int)
    for i := 0; i < n ; i ++ {
        graph[i] = []int{}
	}
	for _,pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]],pre[0])
	}
    //fmt.Println(graph)

	yz := make([][]int,n)
	for i := range yz {
        yz[i] = []int{}
	}

	for i := 0; i < n ; i++ {
		bfs(i,yz,n)
	}


    ans := make([]bool,0)
    //fmt.Println(yz)
    //fmt.Println(ans)
    //[[0 1] [1]]
	for _,query := range queries {
        //fmt.Println(query)
		if len(yz[query[1]]) <= 1 {
            //fmt.Println("yz")
			ans = append(ans,false)
		} else {
            found := 0
			for _,value := range yz[query[1]] {
				if value == query[0] {
					ans = append(ans,true)
                    found = 1
                }
			}
            if found == 0 {
                ans=append(ans,false)
                
            } 
		}
	}

	return ans


}

func bfs(s int,arr [][]int,n int) {

    visited := make([]bool,n)

    queue := []int{}

	queue = append(queue,s)
	visited[s] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		arr[s] = append(arr[s],cur)

		for _,num := range graph[cur] {
			if visited[num] == false {
				queue = append(queue,num)
				visited[num] = true
			}
		}
	}
}
