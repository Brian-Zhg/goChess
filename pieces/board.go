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

	for c:= 0; c< 8; c++{
		newBoard.boardArray[1][c] = NewPawn("white")
		newBoard.boardArray[6][c] = NewPawn("black")
	}
	newBoard.boardArray[0][4] = NewKing("white")
	newBoard.boardArray[7][4] = NewKing("black")

	newBoard.boardArray[0][3] = NewQueen("white")
	newBoard.boardArray[7][3] = NewQueen("black")

	newBoard.boardArray[0][0] = NewRook("white")
	newBoard.boardArray[0][7] = NewRook("white")
	newBoard.boardArray[7][0] = NewRook("black")
	newBoard.boardArray[7][7] = NewRook("black")

	newBoard.boardArray[0][2] = NewBishop("white")
	newBoard.boardArray[0][5] = NewBishop("white")
	newBoard.boardArray[7][2] = NewBishop("black")
	newBoard.boardArray[7][5] = NewBishop("black")

	newBoard.boardArray[0][1] = NewKnight("white")
	newBoard.boardArray[0][6] = NewKnight("white")
	newBoard.boardArray[7][1] = NewKnight("black")
	newBoard.boardArray[7][6] = NewKnight("black")
	
    return newBoard
}

type Position struct{
	row int
	col int
}


func (b Board) ShowBoard(){
	for r:= 0; r < 8; r++{
		for c:= 0; c < 8; c++{
			switch b.boardArray[r][c].pieceName{
			case "empty":
				fmt.Print("_")
			case "pawn":
				fmt.Print("P")
			case "queen":
				fmt.Print("Q")
			case "king":
				fmt.Print("K")
			case "bishop":
				fmt.Print("B")
			case "rook":
				fmt.Print("R")
			case "knight":
				fmt.Print("N")
			}
		}
		fmt.Print("\n")
	}
}