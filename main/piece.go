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
	getColor() bool
	isEmpty() bool
	whoami() string
}

type empty struct {
	color bool
}

func (e empty) possibleMoves(now state) []state {
	return nil
}
func (e empty) getColor() bool {
	return e.color
}

func (e empty) isEmpty() bool {
	return true
}

func (e empty) whoami() string {
	return ("   .  ")
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
func (k king) getColor() bool {
	return k.color
}

func (k king) isEmpty() bool {
	return false
}

func (k king) whoami() string {
	return (" king ")
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
func (q queen) getColor() bool {
	return q.color
}

func (q queen) isEmpty() bool {
	return false
}

func (q queen) whoami() string {
	return (" queen")
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
func (r rook) getColor() bool {
	return r.color
}
func (r rook) isEmpty() bool {
	return false
}

func (r rook) whoami() string {
	return (" rook ")
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
func (b bishop) getColor() bool {
	return b.color
}
func (b bishop) isEmpty() bool {
	return false
}

func (b bishop) whoami() string {
	return ("bishop")
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
func (k knight) getColor() bool {
	return k.color
}
func (k knight) isEmpty() bool {
	return false
}

func (k knight) whoami() string {
	return ("knight")
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
func (p pawn) getColor() bool {
	return p.color
}
func (p pawn) isEmpty() bool {
	return false
}

func (p pawn) whoami() string {
	return (" pawn ")
}
