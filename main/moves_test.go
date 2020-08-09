package main

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	board := createEmptyBoard()
	k := king{true}
	board = board.set(k, 4, 4)
	if board.addr(4, 4) != k {
		t.Errorf("@set nie zgadza sie z @addr")
	}
	if board.curAddr() != k {
		t.Errorf("@set nie zgadza sie z @curAddr")
	}
}

func TestAddr(t *testing.T) {
	board := createEmptyBoard()
	k := king{true}
	board = board.set(k, 7, 7)
	if board.addr(7, 7) != board.tab[63] {
		t.Errorf("@addr nie rowna sie @state.tab")
	}

}

func TestMove(t *testing.T) {
	s := state{0, 0, [64]piece{}, true}
	s = s.emptyBoard()
	s = s.set(king{true}, 0, 0)
	s = s.move(s.curAddr(), 2, 2)
	s2 := state{0, 0, [64]piece{}, true}
	s2 = s2.emptyBoard()
	s2 = s2.set(king{true}, 2, 2)
	if s != s2 {
		s.show()
		s2.show()
		t.Errorf("move nie dziala")

	}

}
func TestCheckStepUp(t *testing.T) {
	var board = createSimpleBoard()
	board = board.setCur(4, 1)
	if checkStepUp(board) != false {
		t.Errorf("pionek nie moze sie ruszyc na poczatku ")
	}
	board = board.setCur(4, 0)
	if checkStepUp(board) != true {
		t.Errorf("krol na poczatku moze sie ruszyc do gory")
	}
}
func TestCheckStep(t *testing.T) {
	var board = createSimpleBoard()
	board = board.setCur(4, 0)
	if len(board.curAddr().possibleMoves(board)) != 0 {
		t.Errorf("Krol na poczatku meczu ma ruchy ")
	}
	board = board.setCur(4, 1)
	if len(board.curAddr().possibleMoves(board)) != 1 {
		t.Errorf("Pionek na poczatku ma wiecej niz jeden ruch %v ", len(board.curAddr().possibleMoves(board)))
	}
}

// ###########
// #  LINES  #
// ###########

// ------------------
//   bench  section
// ------------------
var result1 state

func BenchmarkStepLines(b *testing.B) {
	var test = state{4, 4, [64]piece{}, true}
	var r state
	for n := 0; n < b.N; n++ {
		//r = stepRight(test)
		//r = stepLeft(test)
		r = stepUp(test)
		//r = stepDown(test)
	}
	result1 = r
}

var result2 []state

func BenchmarkLines(b *testing.B) {
	var test = state{4, 4, [64]piece{}, true}
	test = test.emptyBoard()
	test = test.set(king{true}, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = lines(test)
	}
	result2 = r
}

// ---

// -----------------------
// testing  up section
// -----------------------
func Test7RowIsBorderUp(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 7, [64]piece{}, true}
	if isBorderUp(test) == false {
		t.Errorf("Jestem na 8 rzedzie a `isBorderUp` zwraca: %t", isBorderUp(test))
	}
}
func Test6RowIsBorderUp(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 6, [64]piece{}, true}
	if isBorderUp(test) == true {
		t.Errorf("Jestem na 7 rzedzie a `isBorderUp` zwraca: %t", isBorderUp(test))
	}
}
func TestStepUp(t *testing.T) {
	//tworze state na granicy
	var test1 = state{0, 0, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 0, 0)
	var test2 = stepUp(test1)
	test1 = test1.set(empty{true}, 0, 0)
	test1 = test1.set(king{true}, 0, 1)
	test1 = test1.set(empty{true}, 0, 0)
	test1.y++

	if test1 != test2 {
		test1.show()
		test2.show()
		t.Errorf("Manulane przesuniecie nie jest rowne stepUp")
	}
}
func TestUp(t *testing.T) {
	var test1 = state{0, 4, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 0, 4)
	handTable := [3]state{
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true}}

	for i := 0; i < len(handTable); i++ {
		handTable[i] = handTable[i].emptyBoard()
		handTable[i] = handTable[i].set(king{true}, 0, i+5)
	}

	if len(up(test1)) != len(handTable) {
		t.Errorf("Dlugosci recznej i wygenerowanej tablicy nie sa rowne: %v != %v ", len(handTable), len(up(test1)))
		return
	}
	for i := 0; i < 3; i++ {
		if handTable[i] != up(test1)[i] {

			fmt.Println("Hand")
			handTable[i].show()
			fmt.Println()
			fmt.Println("Auto")
			up(test1)[i].show()

			t.Errorf("Tablice nie sa rowne")
		}
	}

}

// ---

// ----------------------
// test  down section
// ----------------------

func Test0RowIsBorderDown(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 0, [64]piece{}, true}
	if isBorderDown(test) == false {
		t.Errorf("Jestem na 8 rzedzie a `isBorderDown` zwraca: %t", isBorderDown(test))
	}
}
func Test1RowIsBorderDown(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 1, [64]piece{}, true}
	if isBorderDown(test) == true {
		t.Errorf("Jestem na 7 rzedzie a `isBorderDown` zwraca: %t", isBorderDown(test))
	}
}
func TestStepDown(t *testing.T) {
	//tworze state na granicy
	var test1 = state{0, 0, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 0, 1)
	var test2 = stepDown(test1)
	test1 = test1.set(empty{true}, 0, 1)
	test1 = test1.set(king{true}, 0, 0)
	test1 = test1.set(empty{true}, 0, 1)
	test1.y--

	if test1 != test2 {
		test1.show()
		test2.show()
		t.Errorf("Manulane przesuniecie nie jest rowne stepDown")
	}
}
func TestDown(t *testing.T) {
	var test1 = state{0, 4, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 0, 3)
	handTable := [3]state{
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true}}

	for i := 0; i < len(handTable); i++ {
		handTable[i] = handTable[i].emptyBoard()
		handTable[i] = handTable[i].set(king{true}, 0, 2-i)
	}

	if len(down(test1)) != len(handTable) {
		t.Errorf("Dlugosci recznej i wygenerowanej tablicy nie sa rowne: %v != %v ", len(handTable), len(down(test1)))
		return
	}

	for i := 0; i < 3; i++ {
		if handTable[i] != down(test1)[i] {
			fmt.Println("Hand")
			handTable[i].show()
			fmt.Println()
			fmt.Println("Auto")
			down(test1)[i].show()
			t.Errorf("Tablice nie sa rowne")
		}
	}

}

// ---

// ----------------------
// test left section
// ----------------------

func Test0ColumnIsBorderLeft(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 0, [64]piece{}, true}
	if isBorderDown(test) == false {
		t.Errorf("Jestem na 0 rzedzie a `isBorderLeft` zwraca: %t", isBorderLeft(test))
	}
}
func Test1ColumnIsBorderLeft(t *testing.T) {
	//tworze state na granicy
	var test = state{1, 0, [64]piece{}, true}
	if isBorderLeft(test) == true {
		t.Errorf("Jestem na 7 rzedzie a `isBorderLeft` zwraca: %t", isBorderLeft(test))
	}
}
func TestStepLeft(t *testing.T) {
	//tworze state na granicy
	var test1 = state{0, 0, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 1, 0)
	var test2 = stepLeft(test1)
	test1 = test1.set(empty{true}, 1, 0)
	test1 = test1.set(king{true}, 0, 0)
	test1 = test1.set(empty{true}, 1, 0)
	test1.x--

	if test1 != test2 {
		test1.show()
		test2.show()
		t.Errorf("Manulane przesuniecie nie jest rowne stepLeft")
	}
}
func TestLeft(t *testing.T) {
	var test1 = state{0, 4, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 3, 0)
	handTable := [3]state{
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true}}

	for i := 0; i < len(handTable); i++ {
		handTable[i] = handTable[i].emptyBoard()
		handTable[i] = handTable[i].set(king{true}, 2-i, 0)
	}

	if len(left(test1)) != len(handTable) {
		t.Errorf("Dlugosci recznej i wygenerowanej tablicy nie sa rowne: %v != %v ", len(handTable), len(left(test1)))
		return
	}

	for i := 0; i < 3; i++ {
		if handTable[i] != left(test1)[i] {
			fmt.Println("Hand")
			handTable[i].show()
			fmt.Println()
			fmt.Println("Auto")
			left(test1)[i].show()
			t.Errorf("Tablice nie sa rowne")
		}
	}

}

// ---

// ----------------------
// test  right section
// ----------------------

func Test7ColumnIsBorderLeft(t *testing.T) {
	//tworze state na granicy
	var test = state{7, 0, [64]piece{}, true}
	if isBorderDown(test) == false {
		t.Errorf("Jestem na 0 rzedzie a `isBorderLeft` zwraca: %t", isBorderLeft(test))
	}
}
func Test6ColumnIsBorderLeft(t *testing.T) {
	//tworze state na granicy
	var test = state{6, 0, [64]piece{}, true}
	if isBorderLeft(test) == true {
		t.Errorf("Jestem na 7 rzedzie a `isBorderLeft` zwraca: %t", isBorderLeft(test))
	}
}
func TestStepRight(t *testing.T) {
	//tworze state na granicy
	var test1 = state{0, 0, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 1, 0)
	var test2 = stepRight(test1)
	test1 = test1.set(empty{true}, 1, 0)
	test1 = test1.set(king{true}, 2, 0)
	test1 = test1.set(empty{true}, 1, 0)
	test1.x++

	if test1 != test2 {
		test1.show()
		test2.show()
		t.Errorf("Manulane przesuniecie nie jest rowne stepRight")
	}
}
func TestRight(t *testing.T) {
	var test1 = state{0, 4, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 4, 0)
	handTable := [3]state{
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true}}

	for i := 0; i < len(handTable); i++ {
		handTable[i] = handTable[i].emptyBoard()
		handTable[i] = handTable[i].set(king{true}, 5+i, 0)
	}

	if len(right(test1)) != len(handTable) {
		t.Errorf("Dlugosci recznej i wygenerowanej tablicy nie sa rowne: %v != %v ", len(handTable), len(right(test1)))
		return
	}

	for i := 0; i < 3; i++ {
		if handTable[i] != right(test1)[i] {
			fmt.Println("Hand")
			handTable[i].show()
			fmt.Println()
			fmt.Println("Auto")
			right(test1)[i].show()
			t.Errorf("Tablice nie sa rowne")
		}
	}

}

// ############
// #  CROSES  #
// ############

// ------------------
// bench  section
// ------------------

func BenchmarkStepCross(b *testing.B) {
	var test = state{4, 4, [64]piece{}, true}
	var r state
	for n := 0; n < b.N; n++ {
		r = stepRightUp(test)
		r = stepLeftUp(test)
		r = stepRightDown(test)
		r = stepLeftDown(test)
	}
	result1 = r
}

var result4 []state

func BenchmarkCross(b *testing.B) {
	var test = createEmptyBoard()
	test = test.set(king{true}, 4, 4)
	var r []state
	for n := 0; n < b.N; n++ {
		r = crosses(test)
	}
	result2 = r
}

// ---

// -----------------------
// testing cross section
// -----------------------

func Test0Col7RowIsBorderLeftUp(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 7, [64]piece{}, true}
	if isBorderLeftUp(test) == false {
		t.Errorf("Jestem na 7 rzedzie i 0 kolumnie a `isBorderLeftUp` zwraca: %t", isBorderLeftUp(test))
	}
}

// TEST
func Test1Col7RowIsBorderLeftUp(t *testing.T) {
	//tworze state na granicy
	var test = state{1, 7, [64]piece{}, true}
	if isBorderLeftUp(test) == false {
		t.Errorf("Jestem na 7 rzedzie a `isBorderLeftUp` zwraca: %t", isBorderLeftUp(test))
	}
}
func Test0Col6RowIsBorderLeftUp(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 6, [64]piece{}, true}
	if isBorderLeftUp(test) == false {
		t.Errorf("Jestem na 0 kolumnie a `isBorderLeftUp` zwraca: %t", isBorderLeftUp(test))
	}
}

func Test1Col6RowIsBorderLeftUp(t *testing.T) {
	//tworze state na granicy
	var test = state{1, 6, [64]piece{}, true}
	if isBorderLeftUp(test) == true {
		t.Errorf("Jestem na 6 rzedzie a `isBorderLeftUp` zwraca: %t", isBorderLeftUp(test))
	}
}
func TestStepLeftUp(t *testing.T) {
	//tworze state na granicy
	var test1 = state{0, 0, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 1, 1)
	var test2 = stepLeftUp(test1)
	test1 = test1.set(empty{true}, 1, 1)
	test1 = test1.set(king{true}, 0, 2)
	test1 = test1.set(empty{true}, 1, 1)
	test1.x--
	test1.y++

	if test1 != test2 {
		test1.show()
		test2.show()
		t.Errorf("Manulane przesuniecie nie jest rowne stepLeftUp")
	}
}
func TestLeftUp(t *testing.T) {
	var test1 = state{0, 4, [64]piece{}, true}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 3, 4)
	handTable := [3]state{
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true},
		state{0, 0, [64]piece{}, true}}

	for i := 0; i < len(handTable); i++ {
		handTable[i] = handTable[i].emptyBoard()
		handTable[i] = handTable[i].set(king{true}, 2-i, 5+i)
	}

	if len(leftUp(test1)) != len(handTable) {
		t.Errorf("Dlugosci recznej i wygenerowanej tablicy nie sa rowne: %v != %v ", len(handTable), len(leftUp(test1)))
		return
	}

	for i := 0; i < 3; i++ {
		if handTable[i] != leftUp(test1)[i] {
			fmt.Println("Hand")
			handTable[i].show()
			fmt.Println()
			fmt.Println("Auto")
			leftUp(test1)[i].show()
			t.Errorf("Tablice nie sa rowne")
		}
	}

}

// ---
