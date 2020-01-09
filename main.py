import pieces
import boards


def move(cur_pos, end_pos):
    moved_piece = board.cur_state[cur_pos[0]][cur_pos[1]]
    # TODO sprawdzic czy przesuniecie sie powiodlo
    moved_piece.move(end_pos)
    board.cur_state[cur_pos[0]][cur_pos[1]] = pieces.Piece()
    board.cur_state[end_pos[0]][end_pos[1]] = moved_piece


pawn = pieces.Piece("pawn", (0, 0), "white")

board = boards.Board()

move(cur_pos=(0, 0), end_pos=(4, 4))

board.show()
