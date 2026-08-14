package pieces

type rook struct{
	chessPiece
}

func NewRook(c string) chessPiece {
	return chessPiece{
		pieceName : "rook",
		color : c,
	}
}