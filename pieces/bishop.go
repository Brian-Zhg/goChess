package pieces

type bishop struct{
	chessPiece
}

func NewBishop(c string) chessPiece {
	return chessPiece{
		pieceName : "bishop",
		color : c,
	}
}
