package pieces

type king struct {
	chessPiece
}

func NewKing(c string) chessPiece {
	return chessPiece{
		pieceName: "king",
		color:     c,
	}
}

func (k *king) LegalMoves(board *Board, pos Position) []Position {
	moves := []Position{}
	distance := [3]int{-1, 0, 1}
	//adds moves to array

	for r := 0; r < len(distance); r++ {
		for c := 0; c < len(distance); c++ {
			if distance[r] != 0 && distance[c] != 0 {
				//not out of bounds
				if !(pos.Row+distance[r] < 0) && !(pos.Row+distance[r] > 7) && !(pos.Col+distance[c] > 7) && !(pos.Col+distance[c] < 0) {
					//empty square
					if board.boardArray[pos.Row+distance[r]][pos.Col+distance[c]].pieceName == "empty" {
						moves = append(moves, Position{Row: pos.Row + distance[r], Col: pos.Col + distance[c]})
					}
					//enemy
					if board.boardArray[pos.Row+distance[r]][pos.Col+distance[c]].color != board.boardArray[pos.Row][pos.Col].color {
						moves = append(moves, Position{Row: pos.Row + distance[r], Col: pos.Col + distance[c]})
					}
				}
			}
		}
	}
	return moves
}
