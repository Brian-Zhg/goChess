package pieces
type pawn struct{
	chessPiece
}



func (p *pawn) LegalMoves(board *Board, pos Position) []Position {
    moves := []Position{}
    //adds moves to array
    if(p.color == "black"){
        moves = append(moves, Position{row: pos.row-1, col: pos.col})
        //checks to see if there is an enemy piece 1 square diagonal to the current pawn
        if(board.boardArray[pos.row-1][pos.col+1].pieceName != "empty" && board.boardArray[pos.row-1][pos.col+1].color != p.color){
            moves = append(moves, Position{row: pos.row-1, col: pos.col+1})
        }
        if(board.boardArray[pos.row-1][pos.col-1].pieceName != "empty" && board.boardArray[pos.row-1][pos.col+1].color != p.color){
            moves = append(moves, Position{row: pos.row-1, col: pos.col-1})
        }
    }

    if(p.color == "white"){
        moves = append(moves, Position{row: pos.row-1, col: pos.col})
        //checks to see if there is an enemy piece 1 square diagonal to the current pawn
        if(board.boardArray[pos.row+1][pos.col+1].pieceName != "empty" && board.boardArray[pos.row+1][pos.col+1].color != p.color){
            moves = append(moves, Position{row: pos.row+1, col: pos.col+1})
        }
        if(board.boardArray[pos.row+1][pos.col-1].pieceName != "empty" && board.boardArray[pos.row+1][pos.col-1].color != p.color){
            moves = append(moves, Position{row: pos.row+1, col: pos.col-1})
        }
    }
    
    return moves
}

