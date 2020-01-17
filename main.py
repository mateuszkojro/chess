import pieces
import boards
import ui


def move(cur_pos, end_pos):
    moved_piece = board.cur_state[cur_pos[0]][cur_pos[1]]
    # TODO sprawdzic czy przesuniecie sie powiodlo
    moved_piece.move(end_pos)
    board.cur_state[cur_pos[0]][cur_pos[1]] = pieces.Piece()
    board.cur_state[end_pos[0]][end_pos[1]] = moved_piece


pawn = pieces.Piece("pawn", (0, 0), "white")

board = boards.Board()

move(cur_pos=(1, 4), end_pos=(2, 4))

board.show()

while (True):
    (p_x, p_y), (e_x, e_y) = ui.my_input() # cords of figure u'd like to move (X from, Y from), (X to, Y to)

    move((int(p_x), int(p_y)), (int(e_x), int(e_y))) #move
    board.show()
    print()
    print(board.convert_to_transfer())
