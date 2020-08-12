package main

import (
	"fmt"
)

var allocSize int

// FIXME bedzie problem bo przeciez gora i dol sa na odwrot
func main() {
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
	now, ocena := evaluate(createSimpleBoard(), 5)
	now.show()
	fmt.Println("ocena: ", ocena)
	//fmt.Println(len(x))
}
