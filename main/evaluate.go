package main

// potrzebuje jakis min max

func analyzeBoard(now state) int {
	noMoves := 0
	relValue := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if now.addr(x, y).isEmpty() == false {
				now = now.setCur(x, y)
				//fmt.Println("cord: ", x, y, "piece: ", now.addr(x, y).whoami(), "color: ", now.addr(x, y).getColor(), "value: ", now.addr(x, y).value(), "possible moves: ", len(now.addr(x, y).possibleMoves(now)))
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
				possibleMoves := now.curAddr().possibleMoves(now)
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
