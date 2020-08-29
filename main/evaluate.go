package main

import "fmt"

// potrzebuje jakis min max
//TODO mozna dodac zeby wchodzilo glebiej jezeli zauwazy szach to by pozwolilo oszukac troche glebokosc tam gdzie sie liczy
func analyzeBoard(now state) int {

	val := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if now.addr(x, y).isEmpty() == false {
				now = now.setCur(x, y)

				if now.curAddr().getColor() == true {
					val += len(now.curAddr().possibleMoves(now))
					val += (now.curAddr().value() * 10)
				} else {
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
	prog := now + 1
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

func (s state) evaluateAlfaBeta(depth int, color bool) (state, int) {
	fmt.Println("depriciated: evalautaAlfaBeta")
	alfa := 200000
	beta := -200000
	var maks int // := alfaBeta(s.moves()[0], depth, alfa, beta, s.player)
	var val int
	var res state
	s.player = color
	//Refactor this for fucks sake
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			progress(y*8+x, 64)
			if s.addr(x, y).isEmpty() == false && s.addr(x, y).getColor() == color {
				//fmt.Println(x, y, s.addr(x, y).whoami())
				s = s.setCur(x, y)
				moves := s.moves()
				res = moves[0]
				for i := 0; i < len(moves); i++ {

					val = alfaBeta(moves[i], depth, alfa, beta, color)
					if s.player == true {
						maks = max(val, maks)
					} else {
						maks = min(val, maks)
					}

					if val == maks {
						res = moves[i]
					}
				}
			}

		}
	}
	fmt.Print('\a')
	return res, maks
}

type info struct {
	ocena int
	move  state
}

func (s state) evaluateAlfaBeta_dev(depth int, color bool) (state, int) {
	alfa := 2000
	beta := -2000
	s.player = color
	var possible []info
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			progress(y*8+x, 64)
			s = s.setCur(x, y)
			if !s.addr(x, y).isEmpty() || !s.addr(x, y).getColor() {
				moves := s.moves()
				for i := 0; i < len(moves); i++ {
					ocena := alfaBeta(moves[i], depth, alfa, beta, color)
					possible = append(possible, info{ocena, moves[i]})
				}
			}
		}
	}
	if color {
		max := info{beta * 100, state{0, 0, [64]piece{}, color}}
		for i := 0; i < len(possible); i++ {
			if possible[i].ocena > max.ocena {
				max = possible[i]
			}
		}
		return max.move, max.ocena
	}
	min := info{alfa * 100, state{0, 0, [64]piece{}, color}}
	for i := 0; i < len(possible); i++ {
		if possible[i].ocena < min.ocena {
			min = possible[i]
		}
	}
	return min.move, min.ocena

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
	ocena := analyzeBoard(node)
	if depth == 0 || (ocena > 2000 || ocena < -2000) {
		//fmt.Println("alfa beta:", node.x, node.y, "glebokosc", depth, node.player, "wartosc", analyzeBoard(node))
		//node.show()
		return analyzeBoard(node) // value
	}
	if player == false {
		value := -200000
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

	value := 200000
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
