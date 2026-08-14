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
