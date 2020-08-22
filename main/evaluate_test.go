package main

import (
	"fmt"
	"testing"
)

var a state

/*
func BenchmarkEvaluate5(b *testing.B) {
	var r state
	for n := 0; n < b.N; n++ {
		r, _ = evaluate(createSimpleBoard(), 5)
	}
	a = r
}
*/

func BenchmarkEvaluate3(b *testing.B) {
	var r state
	for n := 0; n < b.N; n++ {
		r, _ = evaluate(createSimpleBoard(), 3)
	}
	a = r
	r.show()
}

func BenchmarkEvaluate2(b *testing.B) {
	var r state
	for n := 0; n < b.N; n++ {
		r, _ = evaluate(createSimpleBoard(), 2)
	}
	a = r
	r.show()
}

func BenchmarkEvaluateAlfa3(b *testing.B) {
	var r state
	for n := 0; n < b.N; n++ {
		r, _ = createSimpleBoard().evaluateAlfaBeta(3, true)
	}
	a = r
	r.show()
}

func BenchmarkEvaluateAlfa2(b *testing.B) {
	var r state
	for n := 0; n < b.N; n++ {
		r, _ = createSimpleBoard().evaluateAlfaBeta(2, true)
	}
	a = r
	r.show()
}

var i int

func BenchmarkAnalyzeBoard(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = analyzeBoard(createSimpleBoard())
	}
	i = r

}

func TestAnalyzeBoard(t *testing.T) {
	if analyzeBoard(createSimpleBoard()) != 0 {
		t.Errorf("Ocena poczatkowej szachownicy nie jest rowna 0 jest: %d", analyzeBoard(createSimpleBoard()))
	}
	if analyzeBoard(createSimpleBoard()) != 0 {
		t.Errorf("Ocena pustej szachownicy nie jest rowna 0 jest: %d", analyzeBoard(createSimpleBoard()))
	}
}

func TestTake(t *testing.T) {
	board := createEmptyBoard()
	board = board.set(pawn{false, false}, 1, 1)
	board = board.set(queen{true}, 3, 1)
	board.show()
	var ocena int
	board, ocena = board.evaluateAlfaBeta(2, true)
	if ocena < 90 {
		t.Errorf("ocena nie jest poprawna %v - czyli pionek nie zostal zbity", ocena)
		board.show()
	}
}

func TestMateIn1(t *testing.T) {
	ocena := 0
	board := createEmptyBoard()
	board = board.set(rook{true}, 5, 1)
	board = board.set(rook{true}, 4, 2)
	board = board.set(king{false}, 0, 0)
	board = board.set(king{true}, 6, 6)
	board.player = true
	fmt.Println(board.curAddr())
	board, ocena = board.evaluateAlfaBeta(2, true)
	board.show()
	fmt.Println(ocena)
	//board.show()
}

func TestMateIn2(t *testing.T) {

}
