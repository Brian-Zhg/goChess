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
moveable := [8][2]int{{2,1},{1,2},{-2,1},{-1,2},{2,-1},{1,-2},{-2,-1},{-1,-2},}

for check := 0; check < len(moveable); check++{
	if(outBoard(pos, moveable[check][0],moveable[check][1])){
		if(board.boardArray[pos.Row + moveable[check][0]][pos.Col + moveable[check][1]].pieceName == "empty" || board.boardArray[pos.Row + moveable[check][0]][pos.Col + moveable[check][1]].color != k.color){
			moves = append(moves, Position{Row: pos.Row+moveable[check][0], Col: pos.Col+moveable[check][1]})
		}
	}
}

return moves
}
