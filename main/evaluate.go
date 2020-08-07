package main

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
	var eval int32
	for i := 0; i < 64; i++ {
		x := now.tab[i]
		for j := 0; j < len(x.possibleMoves(now)); j++ {
			eval++
		}
	}
	return int(eval)
}
