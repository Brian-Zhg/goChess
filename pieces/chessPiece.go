package pieces

type chessPiece struct {
	pieceName string
	color     string
	pos       Position
}

func NewChessPiece() chessPiece {
	return chessPiece{
		pieceName : "empty",
		color : "empty",
	}
}

func outBoard(currPos Position, row int, col int) bool{
	if(currPos.row +row > 7 || currPos.row +row < 0 || currPos.col +col> 7 || currPos.col +col < 0  ) {
		return false
	}
	return true
}