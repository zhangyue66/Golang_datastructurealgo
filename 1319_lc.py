func makeConnected(n int, connections [][]int) int {
    if len(connections) < n -1 {
		return -1
	}

	parent := []int{}
	for i:= 0 ; i < n ; i ++ {
		parent = append(parent,i)	
	}

	ans := n-1

	for _,connection := range connections {
		a := connection[0]
		b := connection[1]
        pa := Find(a,parent)
        pb := Find(b,parent)

		if pa != pb {
			Union(a,b,parent)
			ans -= 1
		}
	}
	return ans
}

func Find(a int,parent []int) int {
	if parent[a] == a {
		return a
	}
	res := Find(parent[a],parent)
	return res
}

func Union(a int, b int, parent []int) {
	parent_a := Find(a,parent)
	parent_b := Find(b,parent)

	if parent_a != parent_b {
		parent[parent_a] = parent_b
	}
}
