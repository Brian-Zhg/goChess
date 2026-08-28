package pieces
type pawn struct{
	chessPiece
}

func NewPawn(c string) chessPiece {
	return chessPiece{
		pieceName : "pawn",
		color : c,
	}
    
}

func (p *pawn) LegalMoves(board *Board, pos Position) []Position {
    moves := []Position{}
    //adds moves to array
    if(p.color == "black"){
        if(p.firstMove == false && board.boardArray[pos.Row-1][pos.Col].pieceName == "empty" && board.boardArray[pos.Row-2][pos.Col].pieceName == "empty"){
            moves = append(moves, Position{Row: pos.Row-2, Col: pos.Col})
        }
        if(board.boardArray[pos.Row-1][pos.Col].pieceName == "empty"){
            moves = append(moves, Position{Row: pos.Row-1, Col: pos.Col})
        }
        //checks to see if there is an enemy piece 1 square diagonal to the current pawn
        if(board.boardArray[pos.Row-1][pos.Col+1].pieceName != "empty" && board.boardArray[pos.Row-1][pos.Col+1].color != p.color){
            moves = append(moves, Position{Row: pos.Row-1, Col: pos.Col+1})
        }
        if(board.boardArray[pos.Row-1][pos.Col-1].pieceName != "empty" && board.boardArray[pos.Row-1][pos.Col+1].color != p.color){
            moves = append(moves, Position{Row: pos.Row-1, Col: pos.Col-1})
        }
    }

    if(p.color == "white"){
        if(p.firstMove == false && board.boardArray[pos.Row+1][pos.Col].pieceName == "empty" && board.boardArray[pos.Row+2][pos.Col].pieceName == "empty"){
            moves = append(moves, Position{Row: pos.Row+2, Col: pos.Col})
        }
        if(board.boardArray[pos.Row+1][pos.Col].pieceName == "empty"){
            moves = append(moves, Position{Row: pos.Row+1, Col: pos.Col})
        }
        //checks to see if there is an enemy piece 1 square diagonal to the current pawn
        if(board.boardArray[pos.Row+1][pos.Col+1].pieceName != "empty" && board.boardArray[pos.Row+1][pos.Col+1].color != p.color){
            moves = append(moves, Position{Row: pos.Row+1, Col: pos.Col+1})
        }
        if(board.boardArray[pos.Row+1][pos.Col-1].pieceName != "empty" && board.boardArray[pos.Row+1][pos.Col-1].color != p.color){
            moves = append(moves, Position{Row: pos.Row+1, Col: pos.Col-1})
        }
    }
    
    return moves
}

