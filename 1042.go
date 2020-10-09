func gardenNoAdj(N int, paths [][]int) []int {
	graph := make([][]int, N)
	for _,path := range paths {
		graph[path[0]-1] = append(graph[path[0]-1],path[1]-1)
		graph[path[1]-1] = append(graph[path[1]-1],path[0]-1)
	}
    //[[1 2] [0 2] [1 0]]
    
    res := make([]int,N)

	for i := 0 ; i < N ; i ++ {
		colors := [5]int{}
		for _,j := range graph[i] {
			colors[res[j]] = 1
		}
		for c := 1; c<=4; c++ {
			if colors[c] == 0 {
				res[i] = c
				break
			}
		}
	}
	return res
}
