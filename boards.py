from pieces import *
import os


class Board:
    n_move = 0
    cur_color = "white"
    cur_state = []

    def __init__(self):

        for x in range(8):
            temp = []
            for y in range(8):
                temp.append(Piece(in_position=(x, y)))
            self.cur_state.append(temp)

        self.cur_state[0] = [Piece("rook", (0, 0), "black"), Piece("knight", (0, 1), "black"),
                             Piece("bishop", (0, 2), "black"),
                             Piece("king", (0, 3), "black"),
                             Piece("queen", (0, 4), "black"), Piece("bishop", (0, 5), "black"),
                             Piece("knight", (0, 6), "black"),
                             Piece("rook", (0, 7), "black")]

        self.cur_state[1] = [Piece("pawn", (1, 0), "black"), Piece("pawn", (1, 1), "black"),
                             Piece("pawn", (1, 2), "black"),
                             Piece("pawn", (1, 3), "black"),
                             Piece("pawn", (1, 4), "black"), Piece("pawn", (1, 5), "black"),
                             Piece("pawn", (1, 6), "black"),
                             Piece("pawn", (1, 7), "black")]

        self.cur_state[6] = [Piece("pawn", (6, 0), "white"), Piece("pawn", (6, 1), "white"),
                             Piece("pawn", (6, 2), "white"),
                             Piece("pawn", (6, 3), "white"),
                             Piece("pawn", (6, 4), "white"), Piece("pawn", (6, 5), "white"),
                             Piece("pawn", (6, 6), "white"),
                             Piece("pawn", (6, 7), "white")]

        self.cur_state[7] = [Piece("rook", (7, 0), "white"), Piece("knight", (7, 1), "white"),
                             Piece("bishop", (7, 2), "white"),
                             Piece("king", (7, 3), "white"),
                             Piece("queen", (7, 4), "white"), Piece("bishop", (7, 5), "white"),
                             Piece("knight", (7, 6), "white"),
                             Piece("rook", (7, 7), "white")]

    def move(self, cur_pos, end_pos):
        moved_piece = self.cur_state[cur_pos[0]][cur_pos[1]]
        if not self.check_color(moved_piece):
            return False
        if not moved_piece.check(end_pos):
            return False
        moved_piece.move(end_pos)
        self.cur_state[cur_pos[0]][cur_pos[1]] = Piece()
        self.cur_state[end_pos[0]][end_pos[1]] = moved_piece
        return True

    def check_color(self, piece):
        if self.cur_color == piece.color:
            return True
        else:
            return False

    def show(self):
        os.system('clear')

        for i in range(8):
            print(str(i).rjust(6), end="  ")
        print()
        i = 0
        for x in self.cur_state:
            print()
            for y in x:
                print(y.type.rjust(6), end="  ")
            print(" ", i)
            i = i + 1
