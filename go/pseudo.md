# ideas

- czy moges sprawdzac czy bailycz czy czrany po tym czy licznik ruchow jest parzysty czy nie

# podstawowe funkcje ruchow

```python
def is_border_up(position):
   if position + 1 < 8:
      return false
   else:
      return true
def check_step_line_next(board, position):
   return is_border_up(position) + is_my_piece(border, position)

def step_line_next(board, position):
   new_board = board
   new_board[x][position] = None
   new_board[x][position+1] = board[x][position]
   return new_board

def line_h_next(board, position):
   possible_moves = []
   while (check_step_line_next(board, position)):
      if is_next_enemy_piece():
         possible_moves.append(step_line_next(board, position))
         # rezczy zwiazane z zbiciem pionka
         break;
      else:
         board = step_line_next(board, position)
         possible_moves.append(board)
         position = position + 1

def line_h_prev(board, position):
   while (check_step_line_prev(board, position)):
      if is_prev_enemy_piece():
         possible_moves.append(step_line_prev(board, position))
         # rezczy zwiazane z zbiciem pionka
         break;
      else:
         board = step_line_prev(board, position)
         possible_moves.append(board)
         position = position - 1

   return possible_moves
```

analogicznie dla lewo, prawo, dol i wszystkich przekatnych

# ruchy figur

## krol

### specjalne problemy:

- szach
- mat

```python
step_line_h()
step_line_v()
step_cross_left()
step_cross_right()
long_castle()
short_casle()
```

## krolowa

```python
cross_right()
cross_left()
line_h()
line_v()
```

## wieza

```python
line_h()
line_v()
```

## bishop

```python
cross_right()
cross_left()
```

## kon

## pion

### specjalne problemy

- bije inaczej niz chodzi
- bicie w przelocie
- ruch o 2 w pierwszej rundzie

```python
check_step_front()
2 x check_step_front()
```

# boards

## binary

pamiet tylko gdzie stoi cokolwiek (0/1 stoi albo nie)

## total

pamieta gdzie stoi kazda figura (jej typ)

# obiekty

## board

```python
class board:
   def __init__(self):
      self.board = []
      self.pieces = []
```

## piece

```python
class piece:
   def __init__(self):
      self.color
```

po nim dziedzicza wszytskie figury i daja swoje metody do sprawdzania ruchow

# funkcje

## evaluate(board: board) => (rating: double, move: board)

```python
def evaluate(board, depth):
   for piece in board.pieces:
      for new_board in piece.moves(board):
         if (depth > 0):
            evaluate(board=new_board, depth=depth-1)
         else:
            return analyze_board(new_board)

```

## analyze_board(board: board) => rating: double

- punkty za kazda obecna figure i pionka
- punkty za past pionki
- punkty za kazde kontrolowane pole (pole na ktorym moge stanac)
- punkty za ograniczenie kontrolowanych pol przeciwnika
- punkty za strukture pionkow

```python
def analyze_board(board):
   """zasady oceniania jakosci pozycji"""

```

## piece.moves()

zmienia sie w zaleznosci od tego jaka to figura mozliwe rozwiazania:

- [x] polimorfizm
- oznaczanie i sprawdzanie jakiego typu jest konkretna figura
