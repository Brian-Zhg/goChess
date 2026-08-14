package pieces

type king struct{
	chessPiece
}

func NewKing(c string) chessPiece {
	return chessPiece{
		pieceName : "king",
		color : c,
	}
}

func (k *king) LegalMoves(board *Board, pos Position) []Position {
    moves := []Position{}
	distance :=[3] int{-1,0,1}
    //adds moves to array
    if(k.color == "black"){
		for r := 0; r<len(distance); r++{
			for c :=0; c<len(distance); c++{
				if(distance[r] != 0 && distance[c] != 0){
					//not out of bounds
					if(!(pos.row+distance[r] <0) && !(pos.row+distance[r] >7) && !(pos.col+distance[c] >7) && !(pos.col+distance[c] <0)){
						//empty square
						if(board.boardArray[pos.row+distance[r]][pos.col+distance[c]].pieceName == "empty"){
						moves = append(moves, Position{row: pos.row+distance[r], col: pos.col+distance[c]})
						}
						//enemy
						if(board.boardArray[pos.row+distance[r]][pos.col+distance[c]].color == "white"){
						moves = append(moves, Position{row: pos.row+distance[r], col: pos.col+distance[c]})
						}
					}
				}
			}
		}
	}

    if(k.color == "white"){
		for r := 0; r<len(distance); r++{
			for c :=0; c<len(distance); c++{
				if(distance[r] != 0 && distance[c] != 0){
					//not out of bounds
					if(!(pos.row+distance[r] <0) && !(pos.row+distance[r] >7) && !(pos.col+distance[c] >7) && !(pos.col+distance[c] <0)){
						//empty square
						if(board.boardArray[pos.row+distance[r]][pos.col+distance[c]].pieceName == "empty"){
						moves = append(moves, Position{row: pos.row+distance[r], col: pos.col+distance[c]})
						}
						//enemy
						if(board.boardArray[pos.row+distance[r]][pos.col+distance[c]].color == "black"){
						moves = append(moves, Position{row: pos.row+distance[r], col: pos.col+distance[c]})
						}
					}
				}
			}
		}
	}
    
    return moves
}