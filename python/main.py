import pieces
import boards
import ui

board = boards.Board()

result = board.move(cur_pos=(7, 7), end_pos=(5, 5))

board.show()

while True:
    (p_x, p_y), (e_x, e_y) = ui.my_input()

    board.move((int(p_x), int(p_y)), (int(e_x), int(e_y)))
    board.show()
