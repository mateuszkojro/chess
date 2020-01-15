class Position:
    cur_pos = (None, None)
    tar_pos = (None, None)

    def __init__(self, x, y):
        self.cur_pos = x, y


class Piece:
    cur_position = None
    type = None
    color = None
    n_moves = None

    # constructor
    def __init__(self, in_type="*", in_position=(0, 0), in_color="no"):
        self.cur_position = in_position
        self.type = in_type
        self.color = in_color

    # function to move a piece
    def move(self, tar_position):
        self.cur_position = tar_position

    def check_moves(self, tar_position):
        if self.type == "pawn":
            return self.check_pawn(tar_position)
        if self.type == "rook":
            return self.check_rook(tar_position)
        if self.type == "queen":
            return self.check_queen(tar_position)
        if self.type == "bishop":
            return self.check_bishop(tar_position)
        if self.type == "knight":
            return self.check_knight(tar_position)

    def check(self, tar_position):
        if self.is_chceck(tar_position):
            return False
        if not self.check_moves(tar_position):
            return False
        if not self.check_player(tar_position):
            return False
        return True

    # legal moves for every piece

    def check_pawn(self, tar_position):

        return True  # FIXME zwracam true zamiast sprawdzac

        if not tar_position[0] == self.cur_position[0] + 1:
            return True
        else:
            return False

    def check_queen(self, tar_position):
        # zmiana w x i w y taka sama albo zmiana tylko w jednym
        return True

    def check_rook(self, tar_position):
        return True

    def check_bishop(self, tar_position):
        return True

    def check_knight(self, tar_position):
        return True

    def check_pawn(self, tar_position):
        return True

    def is_chceck(self, tar_position):
        return False

    def check_player(self, tar_position):
        return True
