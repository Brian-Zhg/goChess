package pieces
type pawn struct{
	position int
	color string
}
func (p *pawn) LegalMoves(board *Board, pos Position) []Position {
    moves := []Position{}

    // imagine these are legal
    moves = append(moves, Position{Row: 5, Col: 4})
    moves = append(moves, Position{Row: 4, Col: 4})

    return moves
}