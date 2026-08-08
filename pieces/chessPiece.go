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