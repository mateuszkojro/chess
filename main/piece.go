package main

func colorText(text string) string {
	colorReset := "\033[0m"
	color := "\033[31m"
	return string(color) + text + string(colorReset)
}

// TODO napisac testy dla wszystkich figur

// ------------------------
// interface for all pieces
// ------------------------

type piece interface {
	possibleMoves(now state) []state
	getColor() bool
	isEmpty() bool
	whoami() string
	value() int
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

func (e empty) value() int {
	return 0
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
	text := (" king ")
	if k.getColor() == true {
		return text
	}
	return colorText(text)
}
func (k king) value() int {
	return 200
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
	text := (" queen")
	if q.getColor() == true {
		return text
	}
	return colorText(text)
}
func (q queen) value() int {
	return 9
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
	text := (" rook ")
	if r.getColor() == true {
		return text
	}
	return colorText(text)
}
func (r rook) value() int {
	return 5
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
	text := ("bishop")
	if b.getColor() == true {
		return text
	}
	return colorText(text)
}
func (b bishop) value() int {
	return 3
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
	text := ("knight")
	if k.getColor() == true {
		return text
	}
	return colorText(text)
}
func (k knight) value() int {
	return 3
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
	if p.getColor() == true {
		if !checkStepUp(now) {
			now = stepUp(now)
			possibleMoves = append(possibleMoves, now)
			/*
				if !p.moved {
					possibleMoves = append(possibleMoves, stepUp(now))
				}
			*/
		}
	} else {
		if !checkStepDown(now) {
			now = stepDown(now)
			possibleMoves = append(possibleMoves, now)
		}
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
	text := (" pawn ")
	if p.getColor() == true {
		return text
	}
	return colorText(text)
}
func (p pawn) value() int {
	return 1
}
