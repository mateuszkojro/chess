package main

/*
func evaluate(now state, depth int) int {
	for piece := range pieces {
		for state := range piece.possibleMoves(now) {
			if depth > 0 {
				depth--
				return evaluate(state, depth)
			}
			return analyzeBoard(state)
		}
	}
}
*/

func analyzeBoard(now state) int {
	return 1
}
