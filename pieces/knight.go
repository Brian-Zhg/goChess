package pieces

type knight struct{
	chessPiece
}

func NewKnight(c string) chessPiece {
	return chessPiece{
		pieceName : "knight",
		color : c,
	}
}

func (k *knight) LegalMoves(board *Board, pos Position) []Position {
moves := []Position{}
moveable := [8][2]int{{3,1},{1,3},{-3,1},{-1,3},{3,-1},{1,-3},{-3,-1},{-1,-3},}

for check := 0; check < len(moveable); check++{
	if(outBoard(pos, moveable[check][0],moveable[check][1])){
		if(board.boardArray[pos.row + moveable[check][0]][pos.col + moveable[check][1]].pieceName == "empty" || board.boardArray[pos.row + moveable[check][0]][pos.col + moveable[check][1]].color != k.color){
			moves = append(moves, Position{row: pos.row+moveable[check][0], col: pos.col+moveable[check][1]})
		}
	}
}

return moves
}
