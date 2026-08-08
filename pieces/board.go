package pieces

import "fmt"
type Board struct{
	boardArray[8][8] chessPiece
}

type Position struct{
	row int
	col int
}

func NewBoard() Board {
    return Board{
    }
}

func (b Board) ShowBoard(){
	for r:= 0; r < 8; r++{
		for c:= 0; c < 8; c++{
			if(b.boardArray[r][c].pieceName == "empty"){
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
}