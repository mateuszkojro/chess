class Position:
    pos_x = None
    pos_y = None

    def __init__(self, x, y):
        self.pos_x = x
        self.pos_y = y


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
        if self.check(tar_position):
            self.cur_position = tar_position
        else:
            pass  # raise ("Prohibited move", self.type, "cannot go to", self.cur_position)

    def check(self, tar_position):
        if self.type == "pawn":
            self.check_pawn(tar_position)
        if self.type == "rook":
            self.check_rook(tar_position)
        if self.type == "queen":
            self.check_queen(tar_position)
        if self.type == "bishop":
            self.check_bishop(tar_position)
        if self.type == "knight":
            self.check_knight(tar_position)

    # legal moves for every piece

    def check_pawn(self, tar_position):
        if not tar_position[0] == self.cur_position[0] + 1:
            return True
        else:
            return False
        pass

    def check_queen(self, tar_position):
        # zmiana w x i w y taka sama albo zmiana tylko w jednym
        pass

    def check_rook(self, tar_position):
        pass

    def check_bishop(self, tar_position):
        pass

    def check_knight(self, tar_position):
        pass

    def check_pawn(self, tar_position):
        pass
