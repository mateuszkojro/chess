package main

// ----------------------------------
// struktury do przechowywania danych
// ----------------------------------

type boardHighLevel struct {
	position v
	tab      map[v]byte
}
type v struct {
	x int
	y int
}
type figure struct {
	symbol byte
	name   string
	color  bool
}
type state struct {
	pos v
	tab [64]piece
}

func (s state) curAddr() piece {
	return s.addr(s.pos)
}
func (s state) addr(position v) piece {
	return s.tab[position.x + (position.y * 8)]
}

func (s state) addrd(x int, y int) piece {
	return s.tab[(y * 8) + x]
}
func (s state) set (p piece , x int , y int) piece {
	s.tab[(y * 8) + x] = p
	s.pos = v{x,y}
	return s
}
// ===

// ###########
// #  LINES  #
// ###########

// ---------------
//  up section
// ---------------

func isBorderUp(now state) bool {
	if now.pos.y < 7 {
		return false
	}
	return true
}
func stepUp(now state) state {
	UpState := now
	UpState.tab[addr(v{now.pos.x, now.pos.y + 1})] = now.tab[addr(v{now.pos.x, now.pos.y})]
	UpState.pos = v{now.pos.x, now.pos.y + 1}
	return UpState
}

// FIXME wyjebie blad jezeli nademna jest puste pole

func isUpMyPiece(now state) bool {
	x := stepUp(now)
	return now.tab[addr(now.pos)].getColor(now) == x.tab[addr(x.pos)].getColor(x)
}

// FIXME wyjebie blad jezeli nademna jest puste pole
// FIXME zrobic tak jak w isUpMyPiece

func isUpEnemyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) != now.tab[addrd(now.pos.x, now.pos.y+1)].getColor(now)
}

func checkStepUp(now state) bool {
	if isBorderUp(now){
		return false
	} else if isUpEmpty(now) {
		return false
	}
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
	if now.pos.y > 0 {
		return false
	}
	return true
}

func isDownMyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) == now.tab[addrd(now.pos.x, now.pos.y-1)].getColor(now)
}

func isDownEnemyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) != now.tab[addrd(now.pos.x, now.pos.y-1)].getColor(now)
}

// TODO: sprawdzam tylko czy nie jestem na skraju
func checkStepDown(now state) bool {
	if isDownEmpty(now) {
		return false
	}
	return isBorderDown(now) || isDownMyPiece(now)
}

func stepDown(now state) state {
	UpState := now
	UpState.tab[addr(v{now.pos.x, now.pos.y - 1})] = now.tab[addr(v{now.pos.x, now.pos.y})]
	UpState.pos = v{now.pos.x, now.pos.y - 1}
	return UpState
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
	if now.pos.x > 0 {
		return false
	}
	return true
}
func isLeftMyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) == now.tab[addrd(now.pos.x-1, now.pos.y)].getColor(now)
}

func isLeftEnemyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) != now.tab[addrd(now.pos.x-1, now.pos.y)].getColor(now)
}

// TODO: sprawdzam tylko czy nie jestem na skraju
func checkStepLeft(now state) bool {
	if isLeftEmpty(now) {
		return false
	}
	return isBorderLeft(now) || isLeftMyPiece(now)
}

func stepLeft(now state) state {
	UpState := now
	UpState.tab[addr(v{now.pos.x - 1, now.pos.y})] = now.tab[addr(v{now.pos.x, now.pos.y})]
	UpState.pos = v{now.pos.x - 1, now.pos.y}
	return UpState
}

func left(now state) []state {
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

func isRightMyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) == now.tab[addrd(now.pos.x+1, now.pos.y)].getColor(now)
}

func isRightEnemyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) != now.tab[addrd(now.pos.x+1, now.pos.y)].getColor(now)
}

// TODO: sprawdzam tylko czy nie jestem na skraju
func checkStepRight(now state) bool {
	if isRightEmpty(now) {
		return false
	}
	return isBorderRight(now) || isRightMyPiece(now)
}

func stepRight(now state) state {
	UpState := now
	UpState.tab[addr(v{now.pos.x + 1, now.pos.y})] = now.tab[addr(v{now.pos.x, now.pos.y})]
	UpState.pos = v{now.pos.x + 1, now.pos.y}
	return UpState
}

func isBorderRight(now state) bool {
	if now.pos.x < 7 {
		return false
	}
	return true
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
	return isBorderLeft(now) || isBorderUp(now)
}
func stepLeftUp(now state) state {
	LeftUpState := now
	LeftUpState.tab[addr(v{now.pos.x - 1, now.pos.y + 1})] = now.tab[addr(v{now.pos.x, now.pos.y})]
	LeftUpState.pos = v{now.pos.x - 1, now.pos.y + 1}
	return LeftUpState
}

// FIXME wyjebie blad jezeli nademna jest puste pole
// FIXME zrobic tak jak w isUpMyPiece

func isLeftUpMyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) == now.tab[addrd(now.pos.x-1, now.pos.y+1)].getColor(now)
}

// FIXME wyjebie blad jezeli nademna jest puste pole
// FIXME zrobic tak jak w isUpMyPiece

func isLeftUpEnemyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) != now.tab[addrd(now.pos.x-1, now.pos.y+1)].getColor(now)
}

func checkStepLeftUp(now state) bool {
	if isLeftUpEmpty(now) {
		return false
	}
	return isBorderLeftUp(now) || isLeftUpMyPiece(now)
}

func leftUp(now state) []state {
	var possibleMoves []state
	for !checkStepLeftUp(now) {
		now = stepLeftUp(now)
		possibleMoves = append(possibleMoves, now)
		if isLeftUpEnemyPiece(now) {
			break
		}
	}
	return possibleMoves
}

// ---

// ----------------------
//  right up section
// ----------------------

func isBorderRightUp(now state) bool {
	return isBorderRight(now) || isBorderUp(now)
}
func stepRightUp(now state) state {
	RightUpState := now
	RightUpState.tab[addr(v{now.pos.x + 1, now.pos.y + 1})] = now.tab[addr(v{now.pos.x, now.pos.y})]
	RightUpState.pos = v{now.pos.x + 1, now.pos.y + 1}
	return RightUpState
}
func isRightUpEnemyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) == now.tab[addrd(now.pos.x+1, now.pos.y+1)].getColor(now)
}

// FIXME wyjebie blad jezeli nademna jest puste pole
// FIXME zrobic tak jak w isUpMyPiece

func isRightUpMyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) != now.tab[addrd(now.pos.x+1, now.pos.y+1)].getColor(now)
}

func checkStepRightUp(now state) bool {}
	if (isRightUpEmpty(now)){
		return false
	}
	return isBorderRightUp(now) || isRightUpMyPiece(now)
}

// FIXME: to chyba mozna jakos uogulnic dla wszytskich funkcji ale jescze nie wiem jak
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
ccccccccc
// ---------------------
//  left Down section
// ---------------------

func isBorderLeftDown(now state) bool {
	return isBorderLeft(now) || isBorderDown(now)
}
func stepLeftDown(now state) state {
	LeftDownState := now
	LeftDownState.tab[addr(v{now.pos.x - 1, now.pos.y - 1})] = now.tab[addr(v{now.pos.x, now.pos.y})]
	LeftDownState.pos = v{now.pos.x - 1, now.pos.y - 1}
	return LeftDownState
}

// FIXME wyjebie blad jezeli nademna jest puste pole
// FIXME zrobic tak jak w isUpMyPiece

func isLeftDownMyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) == now.tab[addrd(now.pos.x-1, now.pos.y-1)].getColor(now)
}

// FIXME wyjebie blad jezeli nademna jest puste pole
// FIXME zrobic tak jak w isUpMyPiece

func isLeftDownEnemyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) != now.tab[addrd(now.pos.x-1, now.pos.y-1)].getColor(now)
}

func checkStepLeftDown(now state) bool {
	if isLeftDownEmpty(now){
		return false
	}
	return isBorderLeftDown(now) || isLeftDownMyPiece(now)
}

func leftDown(now state) []state {
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
	return isBorderRight(now) || isBorderDown(now)
}
func stepRightDown(now state) state {
	RightDownState := now
	RightDownState.tab[addr(v{now.pos.x + 1, now.pos.y - 1})] = now.tab[addr(v{now.pos.x, now.pos.y})]
	RightDownState.pos = v{now.pos.x + 1, now.pos.y - 1}
	return RightDownState
}

// FIXME wyjebie blad jezeli nademna jest puste pole
// FIXME zrobic tak jak w isUpMyPiece

func isRightDownMyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) == now.tab[addrd(now.pos.x, now.pos.y-1)].getColor(now)
}

// FIXME wyjebie blad jezeli nademna jest puste pole
// FIXME zrobic tak jak w isUpMyPiece

func isRightDownEnemyPiece(now state) bool {
	return now.tab[addr(now.pos)].getColor(now) != now.tab[addrd(now.pos.x, now.pos.y-1)].getColor(now)
}

func checkStepRightDown(now state) bool {
	return isBorderRightDown(now) || isRightDownMyPiece(now)
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
