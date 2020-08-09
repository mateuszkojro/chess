package main

import "fmt"

func evaluate(now state, depth int) int {
	if depth > 0 {
		depth--
		possibleMoves := make([]state, 0, 32)
		for i := 0; i < 64; i++ {
			possibleMoves = now.tab[i].possibleMoves(now)
			for j := 0; j < len(possibleMoves); j++ {
				return evaluate(possibleMoves[j], depth)
			}
		}

	}
	return analyzeBoard(now)
}

func analyzeBoard(now state) int {
	noMoves := 0
	relValue := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if now.addr(x, y).isEmpty() == false {
				fmt.Println("cord: ", x, y, "piece: ", now.addr(x, y).whoami(), "color: ", now.addr(x, y).getColor(), "value: ", now.addr(x, y).value(), "possible moves: ", len(now.addr(x, y).possibleMoves(now)))
				if now.addr(x, y).getColor() == now.player {
					noMoves += len(now.addr(x, y).possibleMoves(now))
					relValue += now.addr(x, y).value()
				} else {
					noMoves -= len(now.addr(x, y).possibleMoves(now))
					relValue -= now.addr(x, y).value()
				}
			}
		}
	}
	return noMoves + relValue
}
