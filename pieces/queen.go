package pieces

type queen struct{
	chessPiece
}

func NewQueen(c string) chessPiece {
	return chessPiece{
		pieceName : "queen",
		color : c,
	}
}