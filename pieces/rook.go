package pieces

type rook struct {
	chessPiece
}

func NewRook(c string) chessPiece {
	return chessPiece{
		pieceName: "rook",
		color:     c,
	}
}

func (r *rook) LegalMoves(board *Board, pos Position) []Position {
	moves := []Position{}
	var left, right, top, bottom, stop bool
	change := 1
	for !stop {
		if !left {
			if !(pos.col-change < 0) && !(pos.col-change > 8) {
				if board.boardArray[pos.row][pos.col-change].pieceName == "empty" {
					moves = append(moves, Position{row: pos.row, col: pos.col - change})
				} else if board.boardArray[pos.row][pos.col-change].color != r.color {
					moves = append(moves, Position{row: pos.row, col: pos.col - change})
					left = true
				} else {
					left = true
				}
			}
		}
		if !right {
			if !(pos.col+change < 0) && !(pos.col+change > 8) {
				if board.boardArray[pos.row][pos.col+change].pieceName == "empty" {
					moves = append(moves, Position{row: pos.row, col: pos.col + change})
				} else if board.boardArray[pos.row][pos.col+change].color != r.color {
					moves = append(moves, Position{row: pos.row, col: pos.col + change})
					right = true
				} else {
					right = true
				}
			}
		}
		if !bottom {
			if !(pos.row-change < 0) && !(pos.row-change > 8) {
				if board.boardArray[pos.row-change][pos.col].pieceName == "empty" {
					moves = append(moves, Position{row: pos.row-change, col: pos.col})
				} else if board.boardArray[pos.row-change][pos.col].color != r.color {
					moves = append(moves, Position{row: pos.row-change, col: pos.col })
					bottom = true
				} else {
					bottom = true
				}
			}
		}
		if !top {
			if !(pos.row+change < 0) && !(pos.row+change > 8) {
				if board.boardArray[pos.row+change][pos.col].pieceName == "empty" {
					moves = append(moves, Position{row: pos.row+change, col: pos.col})
				} else if board.boardArray[pos.row+change][pos.col].color != r.color {
					moves = append(moves, Position{row: pos.row+change, col: pos.col })
					top = true
				} else {
					top = true
				}
			}
		}
		
		change++
		stop = top && bottom && left && right
	}

	return moves
}
