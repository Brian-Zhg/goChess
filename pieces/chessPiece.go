package pieces

type chessPiece struct {
	pieceName string
	color     string
	pos       Position
	firstMove bool
}

func NewChessPiece() chessPiece {
	return chessPiece{
		pieceName: "empty",
		color:     "empty",
		firstMove: false,
	}
}

func outBoard(currPos Position, row int, col int) bool {
	if currPos.Row+row > 7 || currPos.Row+row < 0 || currPos.Col+col > 7 || currPos.Col+col < 0 {
		return false
	}
	return true
}

func (c chessPiece) Name() string{
	return c.pieceName
}