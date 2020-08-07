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

	fmt.Println(evaluate(test1, 1))
	//fmt.Println(len(x))

}
