package main

func main() {

}

func isWinner(player1 []int, player2 []int) int {

	handle := func(player1 []int) int {
		var p int
		ten := -1
		for i := 0; i < len(player1); i++ {
			p += player1[i]

			if ten != -1 && i-ten <= 2 {
				p += player1[i]
			}

			if player1[i] == 10 {
				ten = i
			}
		}
		return p
	}
	p1, p2 := handle(player1), handle(player2)
	if p1 < p2 {
		return 2
	} else if p1 > p2 {
		return 1
	}
	return 0
}
