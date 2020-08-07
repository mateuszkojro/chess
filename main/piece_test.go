package main

import (
	"testing"
)

func BenchmarkKing(b *testing.B) {
	var test = createSimpleBoard()
	var a king
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkQueen(b *testing.B) {
	var test = createSimpleBoard()
	var a queen
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkRook(b *testing.B) {
	var test = createSimpleBoard()
	var a rook
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkBishop(b *testing.B) {
	var test = createSimpleBoard()
	var a bishop
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkKnight(b *testing.B) {
	var test = createSimpleBoard()
	var a knight
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkPawn(b *testing.B) {
	var test = createSimpleBoard()
	var a pawn
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

// HACK its far from correct
func Benchmark1Move(b *testing.B) {
	var board = createSimpleBoard()

	var result []state
	for n := 0; n < b.N; n++ {
		for y := 0; y < 2; y++ {
			for x := 0; x < 8; x++ {
				result = board.addr(x, y).possibleMoves(board)
				result = board.addr(x, 7-y).possibleMoves(board)
			}
		}

	}
	result2 = result
}

var a int

func BenchmarkEvaluate5(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		a = evaluate(createSimpleBoard(), 5)
	}
	a = r
}

func BenchmarkAnalyzeBoard(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		a = analyzeBoard(createSimpleBoard())
	}
	a = r
}