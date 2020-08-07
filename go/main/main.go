package main

import (
	"fmt"
)

var allocSize int

// FIXME bedzie problem bo przeciez gora i dol sa na odwrot
func main() {
	var test1 = createSimpleBoard()
	x := test1.addr(4, 1).possibleMoves(test1)
	if x != nil {
		x[1].show()
	} else {
		fmt.Println("tablica jest pusta")
	}
	fmt.Println(len(x))

}
