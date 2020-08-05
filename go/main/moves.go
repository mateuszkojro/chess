package main

// ----------------------------------
// struktury do przechowywania danych
// ----------------------------------

type figure struct {
	symbol byte
	name   string
	color  bool
}
type state struct {
	x, y int
	tab  [64]piece
}

// TEST

func (s state) curAddr() piece {
	return s.addr(s.x, s.y)
}

// TEST
/*
func (s state) addr(position v) piece {
	return s.tab[position.x+(position.y*8)]
}
*/
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

// TEST
func (s state) move(p piece, x int, y int) state {
	return (s.set(empty{true, true}, s.x, s.y).set(p, x, y))
}

// ===

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
	return stepUp(now).curAddr().isEmpty()
}

// TEST
func isUpEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepUp(now).curAddr().getColor()
}

// TEST
// nie jestem tego pewnien bo nie pamietam jak wyglda test podczas ruchcu wartosci moga byc odwrocone
func checkStepUp(now state) bool {
	if isBorderUp(now) {
		return false
	} else if isUpEmpty(now) {
		return false
	}
	return true
}

func stepUp(now state) state {
	return now.move(now.curAddr(), now.x, now.y+1)
}

func up(now state) []state {
	var possibleMoves []state
	for !checkStepUp(now) {
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
	return stepDown(now).curAddr().isEmpty()
}

// TEST
func isDownEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepDown(now).curAddr().getColor()
}

// TEST
func checkStepDown(now state) bool {
	if isBorderDown(now) {
		return false
	} else if isDownEmpty(now) {
		return false
	}
}

func stepDown(now state) state {
	return now.move(now.curAddr(), now.x, now.y-1)
}

func down(now state) []state {
	var possibleMoves []state
	for !checkStepDown(now) {
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
	return stepLeft(now).curAddr().isEmpty()
}

// TEST
func isLeftEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepLeft(now).curAddr().getColor()
}

// TEST
func checkStepLeft(now state) bool {
	if isBorderLeft(now) {
		return false
	} else if isLeftEmpty(now) {
		return false
	}
}

func stepLeft(now state) state {
	return now.move(now.curAddr(), now.x-1, now.y)
}

func Left(now state) []state {
	var possibleMoves []state
	for !checkStepLeft(now) {
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
// -----------------
func isBorderRight(now state) bool {
	if now.x < 7 {
		return false
	}
	return true
}

// TEST
func isRightEmpty(now state) bool {
	return stepRight(now).curAddr().isEmpty()
}

// TEST
func isRightEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepRight(now).curAddr().getColor()
}

// TEST
func checkStepRight(now state) bool {
	if isBorderRight(now) {
		return false
	} else if isRightEmpty(now) {
		return false
	}
}

func stepRight(now state) state {
	return now.move(now.curAddr(), now.x+1, now.y)
}

func right(now state) []state {
	var possibleMoves []state
	for !checkStepRight(now) {
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

// ----
// combined s
// ----

// FIXME: mozna zrobic szybciej jezeli unikniemy realokacji np przez zrobienie statycznych tablic
func lines(now state) []state {
	var possibleMoves []state
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
	if now.y < 7 || now.x > 0 {
		return false
	}
	return true
}

// TEST
func isLeftUpEmpty(now state) bool {
	return stepLeftUp(now).curAddr().isEmpty()
}

// TEST
func isLeftUpEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepLeftUp(now).curAddr().getColor()
}

// TEST
func checkStepLeftUp(now state) bool {
	if isBorderLeftUp(now) {
		return false
	} else if isLeftUpEmpty(now) {
		return false
	}
}

func stepLeftUp(now state) state {
	return now.move(now.curAddr(), now.x-1, now.y+1)
}

func LeftUp(now state) []state {
	var possibleMoves []state
	for !checkStepLeftUp(now) {
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
	if now.y < 7 || now.x < 7 {
		return false
	}
	return true
}

// TEST
func isRightUpEmpty(now state) bool {
	return stepRightUp(now).curAddr().isEmpty()
}

// TEST
func isRightUpEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepRightUp(now).curAddr().getColor()
}

// TEST
func checkStepRightUp(now state) bool {
	if isBorderRightUp(now) {
		return false
	} else if isRightUpEmpty(now) {
		return false
	}
}

func stepRightUp(now state) state {
	return now.move(now.curAddr(), now.x+1, now.y+1)
}

func rightUp(now state) []state {
	var possibleMoves []state
	for !checkStepRightUp(now) {
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
	if now.y > 0 || now.x > 0 {
		return false
	}
	return true
}

// TEST
func isLeftDownEmpty(now state) bool {
	return stepLeftDown(now).curAddr().isEmpty()
}

// TEST
func isLeftDownEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepLeftDown(now).curAddr().getColor()
}

// TEST
func checkStepLeftDown(now state) bool {
	if isBorderLeftDown(now) {
		return false
	} else if isLeftDownEmpty(now) {
		return false
	}
}

func stepLeftDown(now state) state {
	return now.move(now.curAddr(), now.x-1, now.y-1)
}

func LeftDown(now state) []state {
	var possibleMoves []state
	for !checkStepLeftDown(now) {
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
	if now.y > 0 || now.x < 7 {
		return false
	}
	return true
}

// TEST
func isRightDownEmpty(now state) bool {
	return stepRightDown(now).curAddr().isEmpty()
}

// TEST
func isRightDownEnemyPiece(now state) bool {
	return now.curAddr().getColor() != stepRightDown(now).curAddr().getColor()
}

// TEST
func checkStepRightDown(now state) bool {
	if isBorderRightDown(now) {
		return false
	} else if isRightDownEmpty(now) {
		return false
	}
}

func stepRightDown(now state) state {
	return now.move(now.curAddr(), now.x+1, now.y-1)
}

func rightDown(now state) []state {
	var possibleMoves []state
	for !checkStepRightDown(now) {
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
	var possibleMoves []state
	possibleMoves = append(possibleMoves, leftUp(now)...)
	possibleMoves = append(possibleMoves, rightUp(now)...)
	possibleMoves = append(possibleMoves, leftDown(now)...)
	possibleMoves = append(possibleMoves, rightDown(now)...)

	return possibleMoves
}

// ---
