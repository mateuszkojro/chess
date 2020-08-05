package main

/*
dla bialych
if num > 0{
	if = 10
	if = 40
	if = 41
	if = 50
	if = 90
	if = 1000
}
*/

// TODO napisac testy dla wszystkich figur

// ------------------------
// interface for all pieces
// ------------------------

type piece interface {
	possibleMoves(now state) []state
	getColor(now state) bool
}

// ----
// king
// ----

type king struct {
	color bool
}

func (k king) possibleMoves(now state) []state {
	var possibleMoves []state
	if !checkStepLeftUp(now) {
		possibleMoves = append(possibleMoves, stepLeftUp(now))
	}
	if !checkStepUp(now) {
		possibleMoves = append(possibleMoves, stepUp(now))
	}
	if !checkStepRightUp(now) {
		possibleMoves = append(possibleMoves, stepRightUp(now))
	}
	if !checkStepLeft(now) {
		possibleMoves = append(possibleMoves, stepLeft(now))
	}
	if !checkStepRight(now) {
		possibleMoves = append(possibleMoves, stepRight(now))
	}
	if !checkStepLeftDown(now) {
		possibleMoves = append(possibleMoves, stepLeftDown(now))
	}
	if !checkStepDown(now) {
		possibleMoves = append(possibleMoves, stepDown(now))
	}
	if !checkStepRightDown(now) {
		possibleMoves = append(possibleMoves, stepRightDown(now))
	}
	return possibleMoves
}
func (k king) getColor(now state) bool {
	return k.color
}

// -----
// queen
// -----

type queen struct {
	color bool
}

func (q queen) possibleMoves(now state) []state {
	var possibleMoves []state
	possibleMoves = append(possibleMoves, crosses(now)...)
	possibleMoves = append(possibleMoves, lines(now)...)
	return possibleMoves
}
func (q queen) getColor(now state) bool {
	return q.color
}

// ----
// rook
// ----

type rook struct {
	color bool
}

func (r rook) possibleMoves(now state) []state {
	return lines(now)
}
func (r rook) getColor(now state) bool {
	return r.color
}

// ------
// bishop
// ------

type bishop struct {
	color bool
}

func (b bishop) possibleMoves(now state) []state {
	return crosses(now)
}
func (b bishop) getColor(now state) bool {
	return b.color
}

// ------
// knight
// ------

type knight struct {
	color bool
}

func (k knight) possibleMoves(now state) []state {
	var possibleMoves []state
	return append(possibleMoves, now)
}
func (k knight) getColor(now state) bool {
	return k.color
}

// ----
// pawn
// ----

type pawn struct {
	color bool
	moved bool
}

// FIXME sprawddzam zawsze 2 ruchy do przodu nic wiecej
func (p pawn) possibleMoves(now state) []state {
	var possibleMoves []state
	if !checkStepUp(now) {
		now = stepUp(now)
		possibleMoves = append(possibleMoves, now)
		/*
			if !p.moved {
				possibleMoves = append(possibleMoves, stepUp(now))
			}
		*/
	}

	return possibleMoves
}
func (p pawn) getColor(now state) bool {
	return p.color
}
