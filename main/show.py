import chess
import chess.svg


def convert_to_their(text):
    res = []
    n = 0
    for i, letter in enumerate(text):
        if letter == 'm':
            color = 'b'
        if letter == 'p':
            color = 'w'
        if letter == 'g':
            n = n + 1
        if letter != 'g' and n > 0:
            res.append(str(n))
            n = 0
        if letter == 'a':
            res.append('k')
        if letter == 'A':
            res.append('K')
        if letter == 'b':
            res.append('q')
        if letter == 'B':
            res.append('Q')
        if letter == 'c':
            res.append('r')
        if letter == 'C':
            res.append('R')
        if letter == 'd':
            res.append('b')
        if letter == 'D':
            res.append('B')
        if letter == 'e':
            res.append('n')
        if letter == 'E':
            res.append('N')
        if letter == 'f':
            res.append('p')
        if letter == 'F':
            res.append('P')
        if (i-1) % 8 == 1:
            if n > 0:
                res.append(str(n))
                n = 0
            res.append('/')

    return res


f = open('dane', 'r')
lines = []
for line in f:
    lines.append(line)

fen_str = "".join(convert_to_their(str(lines[len(lines) - 1:])))
fen_str = fen_str[1:-1]
fen_str = fen_str + " w KQkq - 0 1"
board = chess.Board(fen_str)
f.close()
f = open('/tmp/render.svg', 'w')
f.write(str(chess.svg.board(board=board)))
