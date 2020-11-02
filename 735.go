func asteroidCollision(asteroids []int) []int {
	ans := []int{}
	marker := 0
	for _,aster := range asteroids {
        for len(ans) != 0 && aster < 0 && ans[len(ans)-1] > 0 {
            if ans[len(ans)-1] < -aster {
				ans = ans[:len(ans)-1]
				continue
            }else if ans[len(ans)-1] == -aster {
				ans = ans[:len(ans)-1]
			}
			marker = 1
			break
		}
		if marker == 0 {
			ans = append(ans,aster)
			
        } else if marker == 1 {
            marker = 0
        }
	}
	return ans
}
