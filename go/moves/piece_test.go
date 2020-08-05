package main

import (
	"testing"
)

func BenchmarkKing(b *testing.B) {
	var test = state{v{4, 4}, [64]piece{}}
	var a king
	test.tab[addr(v{4, 4})] = a
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.tab[addr(v{4, 4})].possibleMoves(test)
	}
	result2 = r
}

func BenchmarkQueen(b *testing.B) {
	var test = state{v{4, 4}, [64]piece{}}
	var a queen
	test.tab[addr(v{4, 4})] = a
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.tab[addr(v{4, 4})].possibleMoves(test)
	}
	result2 = r
}

func BenchmarkRook(b *testing.B) {
	var test = state{v{4, 4}, [64]piece{}}
	var a rook
	test.tab[addr(v{4, 4})] = a
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.tab[addr(v{4, 4})].possibleMoves(test)
	}
	result2 = r
}

func BenchmarkBishop(b *testing.B) {
	var test = state{v{4, 4}, [64]piece{}}
	var a bishop
	test.tab[addr(v{4, 4})] = a
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.tab[addr(v{4, 4})].possibleMoves(test)
	}
	result2 = r
}

func BenchmarkKnight(b *testing.B) {
	var test = state{v{4, 4}, [64]piece{}}
	var a knight
	test.tab[addr(v{4, 4})] = a
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.tab[addr(v{4, 4})].possibleMoves(test)
	}
	result2 = r
}

func BenchmarkPawn(b *testing.B) {
	var test = state{v{4, 4}, [64]piece{}}
	var a pawn
	test.tab[addr(v{4, 4})] = a
	var r []state
	for n := 0; n < b.N; n++ {
		r = test.tab[addr(v{4, 4})].possibleMoves(test)
	}
	result2 = r
}

func Benchmark1Move(b *testing.B) {
	var board [64]piece

	var p pawn
	for i := 0; i < 8; i++ {
		board[addr(v{i, 1})] = p
		board[addr(v{i, 6})] = p
	}
	var r rook
	board[addr(v{0, 0})] = r
	board[addr(v{7, 0})] = r
	board[addr(v{0, 7})] = r
	board[addr(v{7, 7})] = r

	var kn knight
	board[addr(v{1, 0})] = kn
	board[addr(v{6, 0})] = kn
	board[addr(v{1, 7})] = kn
	board[addr(v{6, 7})] = kn

	var bi bishop
	board[addr(v{2, 0})] = bi
	board[addr(v{5, 0})] = bi
	board[addr(v{2, 7})] = bi
	board[addr(v{5, 7})] = bi

	var q queen
	board[addr(v{3, 0})] = q
	board[addr(v{3, 7})] = q

	var k king
	board[addr(v{4, 0})] = k
	board[addr(v{4, 7})] = k

	var test = state{v{}, board}

	var result []state
	for n := 0; n < b.N; n++ {
		for y := 0; y < 2; y++ {
			for x := 0; x < 8; x++ {
				test.pos = v{x, y}
				result = test.tab[addr(v{x, y})].possibleMoves(test)
				test.pos = v{x, 7 - y}
				result = test.tab[addr(v{x, 7 - y})].possibleMoves(test)
			}
		}

	}
	result2 = result
}
