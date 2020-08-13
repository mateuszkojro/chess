package main

import (
	"fmt"
	//	"runtime"
)

var allocSize int

// FIXME bedzie problem bo przeciez gora i dol sa na odwrot
func main() {
	//	runtime.GOMAXPROCS(8)
	//runtime.GOMAXPROCS(1)

	/*
		rand.Seed(time.Now().UnixNano())
		var test1 = createSimpleBoard()
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				fmt.Println(len(test1.setCur(x, y).curAddr().possibleMoves(test1.setCur(x, y))), test1.setCur(x, y).curAddr().whoami(), x, y)
			}
		}
		test1.show()
		test1 = createEmptyBoard()
		test1 = test1.set(queen{true}, 4, 4)
		fmt.Println(len(test1.curAddr().possibleMoves(test1)))
		fmt.Println(evaluate(test1, 2))
	*/

	fmt.Println("analyzeBoard: ")
	now := createSimpleBoard()
	now.player = !now.player
	now, ocena := now.evaluateAlfaBeta(2)
	//now.show()
	now.player = !now.player
	now, ocena = now.evaluateAlfaBeta(2)

	count := 5

	for i := 0; i < count; i++ {
		progress(i, count)

		now, ocena = now.evaluateAlfaBeta(2)
		now.player = !now.player
		now, ocena = now.evaluateAlfaBeta(2)
		now.player = !now.player
	}

	//fmt.Println(now.player)
	//now, ocena = now.evaluateAlfaBeta(2)
	//now.player = !now.player
	//now, ocena = now.evaluateAlfaBeta(2)
	now.show()
	fmt.Println("ocena: ", ocena)
	//fmt.Println(len(x))
}
