package main

import (
	"fmt"
	"math/rand"
	"time"
)

var allocSize int

// FIXME bedzie problem bo przeciez gora i dol sa na odwrot
func main() {
	rand.Seed(time.Now().UnixNano())
	var test1 = createSimpleBoard()
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			fmt.Println(len(test1.setCur(x, y).curAddr().possibleMoves(test1.setCur(x, y))), test1.setCur(x, y).curAddr().whoami(), x, y)
		}
	}
	fmt.Println(evaluate(test1, 2))
	//fmt.Println(len(x))
}
