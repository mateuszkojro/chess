package main

import (
	"testing"
)

func BenchmarkKing(b *testing.B) {
	var test = state{4, 4, [64]piece{}}
	var a king
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkQueen(b *testing.B) {
	var test = state{4, 4, [64]piece{}}
	var a queen
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkRook(b *testing.B) {
	var test = state{4, 4, [64]piece{}}
	var a rook
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkBishop(b *testing.B) {
	var test = state{4, 4, [64]piece{}}
	var a bishop
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkKnight(b *testing.B) {
	var test = state{4, 4, [64]piece{}}
	var a knight
	test.set(a, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.curAddr().possibleMoves(test)
	}
	result2 = r
}

func BenchmarkPawn(b *testing.B) {
	var test = state{4, 4, [64]piece{}}
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
	var board = state{0, 0, [64]piece{}}
	var p pawn
	for i := 0; i < 8; i++ {
		board = board.set(p, i, 1)
		board = board.set(p, i, 6)
	}
	var r rook
	var k king
	var kn knight
	var bi bishop
	var q queen

	board = board.set(r, 0, 0)
	board = board.set(r, 7, 0)
	board = board.set(r, 0, 7)
	board = board.set(r, 7, 7)
	board = board.set(kn, 1, 0)
	board = board.set(kn, 6, 0)
	board = board.set(kn, 1, 7)
	board = board.set(kn, 6, 7)
	board = board.set(bi, 2, 0)
	board = board.set(bi, 5, 0)
	board = board.set(bi, 2, 7)
	board = board.set(bi, 5, 7)
	board = board.set(q, 3, 0)
	board = board.set(q, 3, 7)
	board = board.set(k, 4, 0)
	board = board.set(k, 4, 7)

	var e empty
	for i := 0; i < 8; i++ {
		board = board.set(e, i, 2)
		board = board.set(e, i, 3)
		board = board.set(e, i, 4)
		board = board.set(e, i, 5)
	}

	var result []state
	for n := 0; n < b.N; n++ {
		for y := 0; y < 2; y++ {
			for x := 0; x < 8; x++ {
				result = board.addr(x, y).possibleMoves(board)
				result = board.addr(x, 8-y).possibleMoves(board)
			}
		}

	}
	result2 = result
}
