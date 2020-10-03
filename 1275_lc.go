//var ttt = [3][3]string{}

func tictactoe(moves [][]int) string {
	ttt := [3][3]string{}
	if len(moves) <= 3 {
		return "Pending"
	}
	
	for i,move := range moves {
		if i % 2 == 0 {
			ttt[move[0]][move[1]] = "X"
		} else {
			ttt[move[0]][move[1]] = "O"
		}
	}
    fmt.Println(ttt)
    resA := check_win(ttt,"X")
    resB := check_win(ttt,"O")
	if resA == true  {
		return "A"
	} else if resB == true  {
		return "B"
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

func check_win(ttt [3][3]string,letter string) bool {
    if (ttt[0][0]== letter && ttt[0][1] == letter && ttt[0][2] == letter) || (ttt[1][0] == letter && ttt[1][1] == letter && ttt[1][2] == letter) || (ttt[2][0] == letter && ttt[2][1] == letter && ttt[2][2] == letter) {
		return true
	}
    if (ttt[0][0] == letter && ttt[1][0] == letter && ttt[2][0] == letter) || (ttt[0][1] == letter && ttt[1][1] == letter && ttt[2][1] == letter) || (ttt[0][2] == letter && ttt[1][2] == letter &&  ttt[2][2] == letter) {
		return true
	}
    if (ttt[0][0] == letter && ttt[1][1] == letter && ttt[2][2] == letter) {
		return true
	}
    if (ttt[0][2] == letter && ttt[1][1] == letter && ttt[2][0] == letter ){
		return true
	}
	return false
}
