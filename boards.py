from pieces import *
import os


class Board:
    cur_state = []

    def __init__(self):

        self.cur_state = [
            [Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white"),
             Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white"),
             Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white")],

            [Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white"),
             Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white"),
             Piece("pawn", (0, 0), "white"), Piece("pawn", (0, 0), "white")]
        ]

        for x in range(6):
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

    def show(self):
        os.system('clear')
        for x in self.cur_state:
            print()
            for y in x:
                print(y.type.rjust(6), end="  ")
            print()
