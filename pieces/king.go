package pieces


type king struct {
	chessPiece
}

func NewKing(c string) chessPiece {
	return chessPiece{
		pieceName: "king",
		color:     c,
	}
}

func (k *king) LegalMoves(board *Board, pos Position) []Position {
	moves := []Position{}
	moveable := [8][2]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	far := 1
	for i := 0; i < len(moveable); i++ {
		row := pos.Row + moveable[i][0]*far
		col := pos.Col + moveable[i][1]*far
		if outBoard(pos, moveable[i][0]*far, moveable[i][1]*far) {
			if board.boardArray[row][col].pieceName == "empty" || board.boardArray[row][col].color != k.color {
				moves = append(moves, Position{
					Row: row,
					Col: col,
				})
			}
		}
	}

	return moves
}


