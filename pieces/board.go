package pieces

import "fmt"


type Board struct {
	boardArray [8][8]chessPiece
}

func (b Board) PieceLocation(row int, col int) string {
	return b.boardArray[row][col].pieceName
}

func (b Board) ReturnPiece(row int, col int) chessPiece {
	return b.boardArray[row][col]
}

func NewBoard() Board {
	newBoard := Board{}

	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			newBoard.boardArray[r][c] = NewChessPiece()
		}
	}

	for c := 0; c < 8; c++ {
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

type Position struct {
	Row int
	Col int
}

func (b Board) ShowBoard() {
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			switch b.boardArray[r][c].pieceName {
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

func ShowMoves(b *Board, row int, col int) []Position {
	piece := b.boardArray[row][col].pieceName
	moves := []Position{}
	switch piece {
	case "pawn":
		p := pawn{chessPiece: b.ReturnPiece(row, col)}
		moves = p.LegalMoves(b, Position{Row: row, Col: col})
	case "queen":
		q := queen{chessPiece: b.ReturnPiece(row, col)}
		moves = q.LegalMoves(b, Position{Row: row, Col: col})
	case "king":
		k := king{chessPiece: b.ReturnPiece(row, col)}
		moves = k.LegalMoves(b, Position{Row: row, Col: col})
	case "bishop":
		bs := bishop{chessPiece: b.ReturnPiece(row, col)}
		moves = bs.LegalMoves(b, Position{Row: row, Col: col})
	case "rook":
		r := rook{chessPiece: b.ReturnPiece(row, col)}
		moves = r.LegalMoves(b, Position{Row: row, Col: col})
	case "knight":
		n := knight{chessPiece: b.ReturnPiece(row, col)}
		moves = n.LegalMoves(b, Position{Row: row, Col: col})
	}
	return moves
}

func Move(b *Board, piece1 Position, piece2 Position){
	b.boardArray[piece2.Row][piece2.Col]= b.boardArray[piece1.Row][piece1.Col]
	b.boardArray[piece1.Row][piece1.Col]=NewChessPiece()
}

func ConfirmMove(b *Board, piece1 Position, piece2 Position) bool{
	if(contains(ShowMoves(b,piece1.Row,piece1.Col),piece2)){
		Move(b, piece1, piece2)
		b.ShowBoard()
		return true 
	}
	return false
}


func contains(moveable []Position, check Position) bool {
    for i := 0;i < len(moveable); i++ {
        if moveable[i] == check {
            return true
        }
    }
    return false
}

func CreatePostion(row int, col int ) Position{
	return Position{Row: row, Col: col}
}