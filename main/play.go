package main

func convertToTxt(board state) string {
	var res string

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			res = res + board.addr(x, y).letterRepresent()
		}
	}
	return res
}

func convertFromTxt(text string) state {
	var res = state{0, 0, [64]piece{}, true}
	if text[0] == 'm' {
		res.player = false
	} else {
		res.player = true
	}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			var p piece
			switch string(text[x+(y*8)+1]) {
			case empty{true}.letterRepresent():
				p = empty{true}
			case empty{false}.letterRepresent():
				p = empty{false}
			case king{true}.letterRepresent():
				p = king{true}
			case king{false}.letterRepresent():
				p = king{false}
			case queen{true}.letterRepresent():
				p = queen{true}
			case queen{false}.letterRepresent():
				p = queen{false}
			case rook{true}.letterRepresent():
				p = rook{true}
			case rook{false}.letterRepresent():
				p = rook{false}
			case bishop{true}.letterRepresent():
				p = bishop{true}
			case bishop{false}.letterRepresent():
				p = bishop{false}
			case knight{true}.letterRepresent():
				p = knight{true}
			case knight{false}.letterRepresent():
				p = knight{false}
			case pawn{true, false}.letterRepresent():
				p = pawn{true, false}
			case pawn{false, false}.letterRepresent():
				p = pawn{false, false}
			}
			res = res.set(p, x, y)
		}
	}

	return res
}
