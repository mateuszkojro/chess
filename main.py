import pieces
import boards
import ui

pawn = pieces.Piece("pawn", (0, 0), "white")

board = boards.Board()

board.move(cur_pos=(1, 4), end_pos=(2, 4))

board.show()

while (True):

    (p_x, p_y), (e_x, e_y) = ui.my_input()

    board.move((int(p_x), int(p_y)), (int(e_x), int(e_y)))
    board.show()
