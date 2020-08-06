package main

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	s := state{0, 0, [64]piece{}}
	s = s.emptyBoard()
	s = s.set(king{true}, 0, 0)
	s = s.move(s.curAddr(), 2, 2)
	s2 := state{0, 0, [64]piece{}}
	s2 = s2.emptyBoard()
	s2 = s2.set(king{true}, 2, 2)
	if s != s2 {
		s.show()
		s2.show()
		t.Errorf("move nie dziala")

	}

}

// ###########
// #  LINES  #
// ###########

// ------------------
// bench  section
// ------------------
var result1 state

func BenchmarkStepLines(b *testing.B) {
	var test = state{4, 4, [64]piece{}}
	var r state
	for n := 0; n < b.N; n++ {
		r = stepRight(test)
		r = stepLeft(test)
		r = stepUp(test)
		r = stepDown(test)
	}
	result1 = r
}

var result2 []state

func BenchmarkLines(b *testing.B) {
	var test = state{4, 4, [64]piece{}}
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
	var test = state{0, 7, [64]piece{}}
	if isBorderUp(test) == false {
		t.Errorf("Jestem na 8 rzedzie a `isBorderUp` zwraca: %t", isBorderUp(test))
	}
}
func Test6RowIsBorderUp(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 6, [64]piece{}}
	if isBorderUp(test) == true {
		t.Errorf("Jestem na 7 rzedzie a `isBorderUp` zwraca: %t", isBorderUp(test))
	}
}
func TestStepUp(t *testing.T) {
	//tworze state na granicy
	var test1 = state{0, 0, [64]piece{}}
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
	var test1 = state{0, 4, [64]piece{}}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 0, 4)
	handTable := [3]state{
		state{0, 0, [64]piece{}},
		state{0, 0, [64]piece{}},
		state{0, 0, [64]piece{}}}

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

func Test7RowIsBorderDown(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 7, [64]piece{}}
	if isBorderDown(test) == false {
		t.Errorf("Jestem na 8 rzedzie a `isBorderDown` zwraca: %t", isBorderDown(test))
	}
}
func Test6RowIsBorderDown(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 6, [64]piece{}}
	if isBorderDown(test) == true {
		t.Errorf("Jestem na 7 rzedzie a `isBorderDown` zwraca: %t", isBorderDown(test))
	}
}
func TestStepDown(t *testing.T) {
	//tworze state na granicy
	var test1 = state{0, 0, [64]piece{}}
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
	var test1 = state{0, 4, [64]piece{}}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 0, 3)
	handTable := [3]state{
		state{0, 0, [64]piece{}},
		state{0, 0, [64]piece{}},
		state{0, 0, [64]piece{}}}

	for i := 0; i < len(handTable); i++ {
		handTable[i] = handTable[i].emptyBoard()
		handTable[i] = handTable[i].set(king{true}, 0, 2-i)
	}

	if len(down(test1)) != len(handTable) {
		t.Errorf("Dlugosci recznej i wygenerowanej tablicy nie sa rowne: %v != %v ", len(handTable), len(down(test1)))
		return
	}
	fmt.Println("tu")

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
// test  left section
// ----------------------

func Test0ColumnIsBorderLeft(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 0, [64]piece{}}
	if isBorderDown(test) == false {
		t.Errorf("Jestem na 0 rzedzie a `isBorderLeft` zwraca: %t", isBorderLeft(test))
	}
}
func Test1ColumnIsBorderLeft(t *testing.T) {
	//tworze state na granicy
	var test = state{1, 0, [64]piece{}}
	if isBorderLeft(test) == true {
		t.Errorf("Jestem na 7 rzedzie a `isBorderLeft` zwraca: %t", isBorderLeft(test))
	}
}
func TestStepLeft(t *testing.T) {
	//tworze state na granicy
	var a king
	var test1 = state{1, 0, [64]piece{}}
	test1.tab[1] = a
	var test2 = stepLeft(test1)
	test1.tab[0] = a
	test1.x--
	if test1 != test2 {
		t.Errorf("Manulane przesuniecie nie jest rowne stepDown: %v != %v ", test1, test2)
	}
}
func TestLeft(t *testing.T) {
	var test1 = state{4, 0, [64]piece{}}

	handTable := [4]state{
		state{3, 0, [64]piece{}},
		state{2, 0, [64]piece{}},
		state{1, 0, [64]piece{}},
		state{0, 0, [64]piece{}}}

	if len(left(test1)) != len(handTable) {
		t.Errorf("Dlugosci recznej i wygenerowanej tablicy nie sa rowne: %v != %v ", len(handTable), len(left(test1)))
		return
	}
	for i := 0; i < 4; i++ {
		if handTable[i] != left(test1)[i] {
			t.Errorf("Tablice nie sa rowne: %v != %v", handTable, left(test1))
		}
	}

}

// ---

// ----------------------
// test  right section
// ----------------------

func Test7ColumnIsBorderLeft(t *testing.T) {
	//tworze state na granicy
	var test = state{7, 0, [64]piece{}}
	if isBorderDown(test) == false {
		t.Errorf("Jestem na 0 rzedzie a `isBorderLeft` zwraca: %t", isBorderLeft(test))
	}
}
func Test6ColumnIsBorderLeft(t *testing.T) {
	//tworze state na granicy
	var test = state{6, 0, [64]piece{}}
	if isBorderLeft(test) == true {
		t.Errorf("Jestem na 7 rzedzie a `isBorderLeft` zwraca: %t", isBorderLeft(test))
	}
}
func TestStepRight(t *testing.T) {
	//tworze state na granicy
	var test1 = state{0, 0, [64]piece{}}
	var a king
	test1.tab[0] = a
	var test2 = stepRight(test1)
	test1.tab[1] = a
	test1.x++
	if test1 != test2 {
		t.Errorf("Manulane przesuniecie nie jest rowne stepRight: %v != %v ", test1, test2)
	}
}
func TestRight(t *testing.T) {
	var test1 = state{3, 0, [64]piece{}}

	handTable := [4]state{
		state{4, 0, [64]piece{}},
		state{5, 0, [64]piece{}},
		state{6, 0, [64]piece{}},
		state{7, 0, [64]piece{}}}

	if len(right(test1)) != len(handTable) {
		t.Errorf("Dlugosci recznej i wygenerowanej tablicy nie sa rowne: %v != %v ", len(handTable), len(right(test1)))
		return
	}
	for i := 0; i < 4; i++ {
		if handTable[i] != right(test1)[i] {
			t.Errorf("Tablice nie sa rowne: %v != %v", handTable, right(test1))
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
	var test = state{4, 4, [64]piece{}}
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
	var test = state{4, 4, [64]piece{}}
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
	var test = state{0, 7, [64]piece{}}
	if isBorderLeftUp(test) == false {
		t.Errorf("Jestem na 8 rzedzie a `isBorderLeftUp` zwraca: %t", isBorderLeftUp(test))
	}
}
func Test1Col7RowIsBorderLeftUp(t *testing.T) {
	//tworze state na granicy
	var test = state{1, 7, [64]piece{}}
	if isBorderLeftUp(test) == false {
		t.Errorf("Jestem na 7 rzedzie a `isBorderLeftUp` zwraca: %t", isBorderLeftUp(test))
	}
}

func Test1Col6RowIsBorderLeftUp(t *testing.T) {
	//tworze state na granicy
	var test = state{1, 6, [64]piece{}}
	if isBorderLeftUp(test) == true {
		t.Errorf("Jestem na 8 rzedzie a `isBorderLeftUp` zwraca: %t", isBorderLeftUp(test))
	}
}
func Test0Col6RowIsBorderLeftUp(t *testing.T) {
	//tworze state na granicy
	var test = state{0, 6, [64]piece{}}
	if isBorderLeftUp(test) == false {
		t.Errorf("Jestem na 7 rzedzie a `isBorderLeftUp` zwraca: %t", isBorderUp(test))
	}
}
func TestStepLeftUp(t *testing.T) {
	//tworze state na granicy
	var test1 = state{1, 1, [64]piece{}}
	var a king
	test1.tab[9] = a
	var test2 = stepLeftUp(test1)
	test1.tab[16] = a
	test1.y++
	test1.x--
	if test1 != test2 {
		t.Errorf("Manulane przesuniecie nie jest rowne stepUp: %v != %v ", test1, test2)
	}
}
func TestLeftUp(t *testing.T) {
	var test1 = state{3, 4, [64]piece{}}

	handTable := [3]state{
		state{2, 5, [64]piece{}},
		state{1, 6, [64]piece{}},
		state{0, 7, [64]piece{}}}

	if len(leftUp(test1)) != len(handTable) {
		t.Errorf("Dlugosci recznej i wygenerowanej tablicy nie sa rowne: %v != %v ", len(handTable), len(leftUp(test1)))
		return
	}
	for i := 0; i < 3; i++ {
		if handTable[i] != leftUp(test1)[i] {
			t.Errorf("Tablice nie sa rowne: %v != %v", handTable, leftUp(test1))
		}
	}

}

// ---
