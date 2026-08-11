package pieces

import "fmt"
type Board struct{
	boardArray [8][8]chessPiece
}

func NewBoard() Board {
    newBoard := Board{}

    for r := 0; r < 8; r++ {
        for c := 0; c < 8; c++ {
            newBoard.boardArray[r][c] = NewChessPiece()
        }
    }

    return newBoard
}

type Position struct{
	row int
	col int
}


func (b Board) ShowBoard(){
	fmt.Println("Board at 1,1: "+ b.boardArray[1][1].pieceName)
	// for r:= 0; r < 8; r++{
	// 	for c:= 0; c < 8; c++{

	// 		if(b.boardArray[r][c].pieceName == "empty"){
	// 			fmt.Print("_")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }
}