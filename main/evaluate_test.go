package main

import "testing"

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
}

func BenchmarkEvaluate2(b *testing.B) {
	var r state
	for n := 0; n < b.N; n++ {
		r, _ = evaluate(createSimpleBoard(), 2)
	}
	a = r
}

func BenchmarkEvaluateAlfa3(b *testing.B) {
	var r state
	for n := 0; n < b.N; n++ {
		r, _ = createSimpleBoard().evaluateAlfaBeta(3)
	}
	a = r
}

func BenchmarkEvaluateAlfa2(b *testing.B) {
	var r state
	for n := 0; n < b.N; n++ {
		r, _ = createSimpleBoard().evaluateAlfaBeta(2)
	}
	a = r
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
