package main

import "testing"

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

func TestAnalyzeBoard(t *testing.T) {
	if analyzeBoard(createSimpleBoard()) != 0 {
		t.Errorf("Ocena poczatkowej szachownicy nie jest rowna 0 jest: %d", analyzeBoard(createSimpleBoard()))
	}
	if analyzeBoard(createSimpleBoard()) != 0 {
		t.Errorf("Ocena pustej szachownicy nie jest rowna 0 jest: %d", analyzeBoard(createSimpleBoard()))
	}
}
