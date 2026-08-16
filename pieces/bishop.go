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

func (b *bishop) LegalMoves(board *Board, pos Position) []Position {
moves := []Position{}

return moves
}
