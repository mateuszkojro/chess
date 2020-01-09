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

    # constructor
    def __init__(self, in_type="blank", in_position=(0, 0), in_color="no"):
        self.cur_position = in_position
        self.type = in_type
        self.color = in_color

    # function to move a piece
    def move(self, tar_position):
        if not self.check(tar_position):
            self.cur_position = tar_position
        else:
            raise ("Prohibited move", self.type, "cannot go to", self.cur_position)

    def check(self, tar_position):
        if self.type == "pawn":
            pass
        if self.type == "rook":
            pass
        if self.type == "queen":
            pass
        if self.type == "bishop":
            pass
        if self.type == "knight":
            pass

    # legal moves for every piece

    def check_pawn(self, tar_position):
        pass

    def check_queen(self, tar_position):
        pass

    def check_rook(self, tar_position):
        pass

    def check_bishop(self, tar_position):
        pass

    def check_knight(self, tar_position):
        pass

    def check_pawn(self, tar_position):
        pass
