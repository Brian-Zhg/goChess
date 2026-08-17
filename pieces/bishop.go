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
moveable := [4][2]int{{1,1},{-1,1},{1,-1},{-1,-1}}
far := 1
finished := [4]bool{}
for !(finished[0] && finished[1] && finished[2] && finished[3]) {
        for i := 0; i < len(moveable); i++ {
            if !finished[i] {
                row := pos.row + moveable[i][0]*far
                col := pos.col + moveable[i][1]*far

                if !outBoard(pos, moveable[i][0]*far, moveable[i][1]*far) {
                    if board.boardArray[row][col].pieceName == "empty" {
                        moves = append(moves, Position{
                            row: row,
                            col: col,
                        })
                    } else if board.boardArray[row][col].color != b.color {
                        moves = append(moves, Position{
                            row: row,
                            col: col,
                        })
                        finished[i] = true
                    } else {
                        finished[i] = true
                    }
                } else {
                    finished[i] = true
                }
            }
        }

        far++
    }

return moves
}
