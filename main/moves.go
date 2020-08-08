package main

import "fmt"

// ----------------------------------
// struktury do przechowywania danych
// ----------------------------------

type state struct {
	x, y int
	tab  [64]piece
}

// TEST
func (s state) curAddr() piece {
	return s.addr(s.x, s.y)
}
func (s state) setCur(x, y int) state {
	s.x, s.y = x, y
	return s
}

// TEST
func (s state) addr(x int, y int) piece {
	return s.tab[(y*8)+x]
}

// TEST
func (s state) set(p piece, x int, y int) state {
	s.tab[(y*8)+x] = p
	s.x, s.y = x, y
	return s
}

// TODO: mozna lepiej bez figury do wstawienia
func (s state) move(p piece, x int, y int) state {
	//fmt.Println("i moved", s.curAddr().whoami(), "from: ", s.x, s.y, " to: ", x, y)
	t := s.set(empty{true}, s.x, s.y)
	return t.set(s.curAddr(), x, y)
}

func (s state) emptyBoard() state {
	e := empty{true}
	for i := 0; i < len(s.tab); i++ {
		s.tab[i] = e
	}
	return s
}

func (s state) show() {
	fmt.Println("x: ", s.x, "y: ", s.y)
	fmt.Print(" ")
	for i := 0; i < 8; i++ {
		fmt.Print("   ", i, "   ")
	}
	fmt.Println()
	for y := 0; y < 8; y++ {
		fmt.Print(y, " ")
		for x := 0; x < 8; x++ {
			fmt.Print(s.addr(x, y).whoami())
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

// ===

// ----------------------------
// dodatkowe funkcje pomocnicze
// ----------------------------

func createEmptyBoard() state {
	var board = state{0, 0, [64]piece{}}
	board = board.emptyBoard()
	return board
}

func createSimpleBoard() state {
	var board = state{0, 0, [64]piece{}}
	for i := 0; i < 8; i++ {
		board = board.set(pawn{true, false}, i, 1)
		board = board.set(pawn{false, false}, i, 6)
	}

	board = board.set(rook{true}, 0, 0)
	board = board.set(rook{true}, 7, 0)
	board = board.set(rook{false}, 0, 7)
	board = board.set(rook{false}, 7, 7)
	board = board.set(knight{true}, 1, 0)
	board = board.set(knight{true}, 6, 0)
	board = board.set(knight{false}, 1, 7)
	board = board.set(knight{false}, 6, 7)
	board = board.set(bishop{true}, 2, 0)
	board = board.set(bishop{true}, 5, 0)
	board = board.set(bishop{false}, 2, 7)
	board = board.set(bishop{false}, 5, 7)
	board = board.set(queen{true}, 3, 0)
	board = board.set(queen{false}, 3, 7)
	board = board.set(king{true}, 4, 0)
	board = board.set(king{false}, 4, 7)

	var e empty
	for i := 0; i < 8; i++ {
		board = board.set(e, i, 2)
		board = board.set(e, i, 3)
		board = board.set(e, i, 4)
		board = board.set(e, i, 5)
	}
	return board
}

// ---

// ###########
// #  LINES  #
// ###########

// ---------------
//  up section
// ---------------

func isBorderUp(now state) bool {
	if now.y < 7 {
		return false
	}
	return true
}

// TEST
func isUpEmpty(now state) bool {
	return now.addr(now.x, now.y+1).isEmpty()
}

// TEST
func isUpEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepUp(now).curAddr().getColor()
}

// TEST
func checkStepUp(now state) bool {
	//fmt.Println("is border up: ", isBorderUp(now))
	if isBorderUp(now) {
		return true
	}
	//fmt.Println("is up empty: ", isUpEmpty(now))
	return !isUpEmpty(now)
}

func stepUp(now state) state {
	return now.move(now.curAddr(), now.x, now.y+1)
}

func up(now state) []state {
	possibleMoves := make([]state, 0, 8)
	//var possibleMoves []state
	//fmt.Println("jestem poza petla", checkStepUp(now))
	for checkStepUp(now) == false {
		//fmt.Println("jestem w petli")
		if isUpEnemyPiece(now) {
			possibleMoves = append(possibleMoves, stepUp(now))
			break
		}
		now = stepUp(now)
		possibleMoves = append(possibleMoves, now)

	}
	return possibleMoves
}

// ===

// -----------------
//  down section
// -----------------

func isBorderDown(now state) bool {
	if now.y > 0 {
		return false
	}
	return true
}

// TEST
func isDownEmpty(now state) bool {
	return now.addr(now.x, now.y-1).isEmpty()
}

// TEST
func isDownEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepDown(now).curAddr().getColor()
}

// TEST
func checkStepDown(now state) bool {
	//fmt.Println("is border Down: ", isBorderDown(now))
	if isBorderDown(now) {
		return true
	}
	//fmt.Println("is Down empty: ", isDownEmpty(now))
	return !isDownEmpty(now)
}

func stepDown(now state) state {
	return now.move(now.curAddr(), now.x, now.y-1)
}

func down(now state) []state {
	possibleMoves := make([]state, 0, 8)
	//var possibleMoves []state
	//fmt.Println("jestem poza petla", checkStepDown(now))
	for checkStepDown(now) == false {
		//fmt.Println("jestem w petli")
		if isDownEnemyPiece(now) {
			possibleMoves = append(possibleMoves, stepDown(now))
			break
		}
		now = stepDown(now)
		possibleMoves = append(possibleMoves, now)

	}
	return possibleMoves
}

// ===

// -----------------
//  left section
// -----------------
func isBorderLeft(now state) bool {
	if now.x > 0 {
		return false
	}
	return true
}

// TEST
func isLeftEmpty(now state) bool {
	return now.addr(now.x-1, now.y).isEmpty()
}

// TEST
func isLeftEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepLeft(now).curAddr().getColor()
}

// TEST
func checkStepLeft(now state) bool {
	//fmt.Println("is border Left: ", isBorderLeft(now))
	if isBorderLeft(now) {
		return true
	}
	//fmt.Println("is Left empty: ", isLeftEmpty(now))
	return !isLeftEmpty(now)
}

func stepLeft(now state) state {
	return now.move(now.curAddr(), now.x-1, now.y)
}

func left(now state) []state {
	possibleMoves := make([]state, 0, 8)
	//var possibleMoves []state
	//fmt.Println("jestem poza petla", checkStepLeft(now))
	for checkStepLeft(now) == false {
		//fmt.Println("jestem w petli")
		if isLeftEnemyPiece(now) {
			possibleMoves = append(possibleMoves, stepLeft(now))
			break
		}
		now = stepLeft(now)
		possibleMoves = append(possibleMoves, now)

	}
	return possibleMoves
}

// ===

// ------------------
//  right section
// ------------------
func isBorderRight(now state) bool {
	if now.x < 7 {
		return false
	}
	return true
}

// TEST
func isRightEmpty(now state) bool {
	return now.addr(now.x+1, now.y).isEmpty()
}

// TEST
func isRightEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepRight(now).curAddr().getColor()
}

// TEST
func checkStepRight(now state) bool {
	//fmt.Println("is border Right: ", isBorderRight(now))
	if isBorderRight(now) {
		return true
	}
	//fmt.Println("is Right empty: ", isRightEmpty(now))
	return !isRightEmpty(now)
}

func stepRight(now state) state {
	return now.move(now.curAddr(), now.x+1, now.y)
}

func right(now state) []state {
	possibleMoves := make([]state, 0, 8)
	//var possibleMoves []state
	//fmt.Println("jestem poza petla", checkStepRight(now))
	for checkStepRight(now) == false {
		//fmt.Println("jestem w petli")
		if isRightEnemyPiece(now) {
			possibleMoves = append(possibleMoves, stepRight(now))
			break
		}
		now = stepRight(now)
		possibleMoves = append(possibleMoves, now)

	}
	return possibleMoves
}

// ===

// --------------
// combined lines
// --------------

// FIXME: mozna zrobic szybciej jezeli unikniemy realokacji np przez zrobienie statycznych tablic
func lines(now state) []state {
	//var possibleMoves []state
	possibleMoves := make([]state, 0, 16)
	possibleMoves = append(possibleMoves, up(now)...)
	possibleMoves = append(possibleMoves, down(now)...)
	possibleMoves = append(possibleMoves, left(now)...)
	possibleMoves = append(possibleMoves, right(now)...)

	return possibleMoves
}

// ---

// ############
// #  CROSES  #
// ############

// ---------------------
//  left up section
// ---------------------

func isBorderLeftUp(now state) bool {
	if now.y < 7 && now.x > 0 {
		return false
	}
	return true
}

// TEST
func isLeftUpEmpty(now state) bool {
	return now.addr(now.x-1, now.y+1).isEmpty()
}

// TEST
func isLeftUpEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepLeftUp(now).curAddr().getColor()
}

// TEST
func checkStepLeftUp(now state) bool {
	//fmt.Println("is border LeftUp: ", isBorderLeftUp(now))
	if isBorderLeftUp(now) {
		return true
	}
	//fmt.Println("is LeftUp empty: ", isLeftUpEmpty(now))
	return !isLeftUpEmpty(now)
}

func stepLeftUp(now state) state {
	return now.move(now.curAddr(), now.x-1, now.y+1)
}

func leftUp(now state) []state {
	possibleMoves := make([]state, 0, 8)
	//var possibleMoves []state
	//fmt.Println("jestem poza petla", checkStepLeftUp(now))
	for checkStepLeftUp(now) == false {
		//fmt.Println("jestem w petli")
		if isLeftUpEnemyPiece(now) {
			possibleMoves = append(possibleMoves, stepLeftUp(now))
			break
		}
		now = stepLeftUp(now)
		possibleMoves = append(possibleMoves, now)

	}
	return possibleMoves
}

// ---

// ----------------------
//  right up section
// ----------------------

func isBorderRightUp(now state) bool {
	if now.y < 7 && now.x < 7 {
		return false
	}
	return true
}

// TEST
func isRightUpEmpty(now state) bool {
	return now.addr(now.x+1, now.y+1).isEmpty()
}

// TEST
func isRightUpEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepRightUp(now).curAddr().getColor()
}

// TEST
func checkStepRightUp(now state) bool {
	//fmt.Println("is border RightUp: ", isBorderRightUp(now))
	if isBorderRightUp(now) {
		return true
	}
	//fmt.Println("is RightUp empty: ", isRightUpEmpty(now))
	return !isRightUpEmpty(now)
}

func stepRightUp(now state) state {
	return now.move(now.curAddr(), now.x+1, now.y+1)
}

func rightUp(now state) []state {
	possibleMoves := make([]state, 0, 8)
	//var possibleMoves []state
	//fmt.Println("jestem poza petla", checkStepRightUp(now))
	for checkStepRightUp(now) == false {
		//fmt.Println("jestem w petli")
		if isRightUpEnemyPiece(now) {
			possibleMoves = append(possibleMoves, stepRightUp(now))
			break
		}
		now = stepRightUp(now)
		possibleMoves = append(possibleMoves, now)

	}
	return possibleMoves
}

// ---
// ---------------------
//  left Down section
// ---------------------
func isBorderLeftDown(now state) bool {
	if now.y > 0 && now.x > 0 {
		return false
	}
	return true
}

// TEST
func isLeftDownEmpty(now state) bool {
	return now.addr(now.x-1, now.y-1).isEmpty()
}

// TEST
func isLeftDownEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepLeftDown(now).curAddr().getColor()
}

// TEST
func checkStepLeftDown(now state) bool {
	//fmt.Println("is border LeftDown: ", isBorderLeftDown(now))
	if isBorderLeftDown(now) {
		return true
	}
	//fmt.Println("is LeftDown empty: ", isLeftDownEmpty(now))
	return !isLeftDownEmpty(now)
}

func stepLeftDown(now state) state {
	return now.move(now.curAddr(), now.x-1, now.y-1)
}

func leftDown(now state) []state {
	possibleMoves := make([]state, 0, 8)
	//var possibleMoves []state
	//fmt.Println("jestem poza petla", checkStepLeftDown(now))
	for checkStepLeftDown(now) == false {
		//fmt.Println("jestem w petli")
		if isLeftDownEnemyPiece(now) {
			possibleMoves = append(possibleMoves, stepLeftDown(now))
			break
		}
		now = stepLeftDown(now)
		possibleMoves = append(possibleMoves, now)

	}
	return possibleMoves
}

// ---

// ------------------------
//  Right Down section
// ------------------------

func isBorderRightDown(now state) bool {
	if now.y > 0 && now.x < 7 {
		return false
	}
	return true
}

// TEST
func isRightDownEmpty(now state) bool {
	return now.addr(now.x+1, now.y-1).isEmpty()
}

// TEST
func isRightDownEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepRightDown(now).curAddr().getColor()
}

// TEST
func checkStepRightDown(now state) bool {
	//fmt.Println("is border RightDown: ", isBorderRightDown(now))
	if isBorderRightDown(now) {
		return true
	}
	//fmt.Println("is RightDown empty: ", isRightDownEmpty(now))
	return !isRightDownEmpty(now)
}

func stepRightDown(now state) state {
	return now.move(now.curAddr(), now.x+1, now.y-1)
}

func rightDown(now state) []state {
	possibleMoves := make([]state, 0, 8)
	//var possibleMoves []state
	//fmt.Println("jestem poza petla", checkStepRightDown(now))
	for checkStepRightDown(now) == false {
		//fmt.Println("jestem w petli")
		if isRightDownEnemyPiece(now) {
			possibleMoves = append(possibleMoves, stepRightDown(now))
			break
		}
		now = stepRightDown(now)
		possibleMoves = append(possibleMoves, now)

	}
	return possibleMoves
}

// ---

// ----
// combined
// ----

// FIXME: mozna zrobic szybciej jezeli unikniemy realokacji np przez zrobienie statycznych tablic
func crosses(now state) []state {
	//var possibleMoves []state
	possibleMoves := make([]state, 0, 16)
	possibleMoves = append(possibleMoves, leftUp(now)...)
	possibleMoves = append(possibleMoves, rightUp(now)...)
	possibleMoves = append(possibleMoves, leftDown(now)...)
	possibleMoves = append(possibleMoves, rightDown(now)...)

	return possibleMoves
}

// ---
