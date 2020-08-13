package main

import "fmt"

// potrzebuje jakis min max

func analyzeBoard(now state) int {
	//fmt.Println("-------- new board ----------")
	//now.show()

	val := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if now.addr(x, y).isEmpty() == false {
				now = now.setCur(x, y)
				//fmt.Println("cord: ", x, y, "piece: ", now.addr(x, y).whoami(), "color: ", now.addr(x, y).getColor(), "value: ", now.addr(x, y).value(), "possible moves: ", len(now.addr(x, y).possibleMoves(now)))

				if now.curAddr().getColor() == true {
					//fmt.Println("+")
					val += len(now.curAddr().possibleMoves(now))
					val += (now.curAddr().value() * 10)
				} else {
					//fmt.Println("-")
					val -= len(now.curAddr().possibleMoves(now))
					val -= (now.curAddr().value() * 10)
				}

			}
		}
	}

	//fmt.Println(val)
	return val
}

func evaluate(now state, depth int) (state, int) {
	if depth > 0 {
		depth--
		//var possibleMoves []state
		//possibleMoves := make([]state, 0, 32)
		max := now
		maxOcena := 0
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				now = now.setCur(x, y)
				possibleMoves := now.moves()
				for j := 0; j < len(possibleMoves); j++ {
					//fmt.Println(len(possibleMoves))
					res, ocena := evaluate(possibleMoves[j], depth)
					if ocena > maxOcena {
						max = res
						maxOcena = ocena
					}
				}
			}
		}
		return max, maxOcena
	}
	return now, analyzeBoard(now)
}

func progress(now, to int) {
	fmt.Print("\r")
	prog := now
	fmt.Print(prog)
	fmt.Print("/ ", to, " pola ")
	fmt.Print("|")
	for i := 0; i < prog; i++ {
		fmt.Print("#")
	}
	for i := 0; i < to-(prog); i++ {
		fmt.Print(" ")
	}
	fmt.Print("|")
}

func (s state) evaluateAlfaBeta(depth int) (state, int) {
	alfa := 200
	beta := -200
	var maks int // := alfaBeta(s.moves()[0], depth, alfa, beta, s.player)
	var val int
	var res state

	//Refactor this for fucks sake
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if s.addr(x, y).isEmpty() == true {

				s = s.setCur(x, y)
				moves := s.moves()
				for i := 0; i < len(moves); i++ {
					val = alfaBeta(moves[i], depth, alfa, beta, s.player)
					maks = max(val, maks)
					if val == maks {
						res = moves[i]
					}
				}
			}

		}
	}

	return res, maks
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func alfaBeta(node state, depth int, alfa int, beta int, player bool) int {
	node.player = player

	//fmt.Println(analyzeBoard(node))

	if depth == 0 /*|| (analyzeBoard(node) > 200 || analyzeBoard(node) < -200)*/ {
		//fmt.Println("alfa beta:", node.x, node.y, "glebokosc", depth, node.player, "wartosc", analyzeBoard(node))
		//node.show()
		return analyzeBoard(node) // value
	}
	if player == false {
		value := -1000
		moves := node.moves()
		for i := 0; i < len(moves); i++ {
			// w sumie nie wiem czy ! player jest correct
			value = max(value, alfaBeta(moves[i], depth-1, alfa, beta, false))
			alfa = max(alfa, value)
			if alfa >= beta {
				//break
			}
		}
		return value
	}

	value := 1000
	moves := node.moves()
	for i := 0; i < len(moves); i++ {
		value = min(value, alfaBeta(moves[i], depth-1, alfa, beta, true))
		beta = min(beta, value)
		if beta <= alfa {
			//break
		}
	}
	return value

}
