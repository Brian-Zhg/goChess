package pieces

import "fmt"

type Board struct {
	boardArray [8][8]chessPiece
	turn       bool //white == true
	bKing      Position
	wKing      Position
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

	newBoard.wKing = Position{
		Row: 0,
		Col: 4,
	}

	newBoard.bKing = Position{
		Row: 7,
		Col: 4,
	}

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
	color := b.boardArray[row][col].color
	moves := []Position{}
	if b.turn == true && color == "white" || b.turn == false && color == "black" {
		moves = getMoves(b, row, col)
	}
	return moves
}

func getMoves(b *Board, row int, col int) []Position {
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

//actually moves the piece
func Move(b *Board, piece1 Position, piece2 Position) {
	if b.boardArray[piece1.Row][piece1.Col].firstMove == false {
		b.boardArray[piece1.Row][piece1.Col].firstMove = true
	}

	b.boardArray[piece2.Row][piece2.Col] =
		b.boardArray[piece1.Row][piece1.Col]

	// Update its position AFTER copying
	b.boardArray[piece2.Row][piece2.Col].pos = Position{
		Row: piece2.Row,
		Col: piece2.Col,
	}
	if b.boardArray[piece2.Row][piece2.Col].pieceName == "king" {
		if b.boardArray[piece2.Row][piece2.Col].color == "white" {
			b.wKing = Position{
				Row:piece2.Row,
				Col:piece2.Col,
			}
		} else {
			b.bKing = Position{
				Row:piece2.Row,
				Col:piece2.Col,
			}
		}
	}

	// Empty the old square
	b.boardArray[piece1.Row][piece1.Col] = NewChessPiece()

	b.turn = !b.turn
}

//checks to see if the move causes the king to be in check (illegal move)
func ConfirmMove(b *Board, piece1 Position, piece2 Position) bool {
	if contains(ShowMoves(b, piece1.Row, piece1.Col), piece2) {
		temp:= *b
		Move(&temp, piece1, piece2)
		if(!inCheck(&temp, !temp.turn)){
			Move(b, piece1, piece2)
			return true
		}
		b.ShowBoard()
		fmt.Print(inCheck(b, b.turn))
	}
	return false
}

func contains(moveable []Position, check Position) bool {
	for i := 0; i < len(moveable); i++ {
		if moveable[i] == check {
			return true
		}
	}
	return false
}

func CreatePostion(row int, col int) Position {
	return Position{Row: row, Col: col}
}

func (b Board) ReturnTurn() bool {
	return b.turn
}

func inCheck(board *Board, turn bool) bool {
	if turn == false {
		fmt.Print(board.bKing.Row)
		fmt.Print(board.bKing.Col)
		for row := 0; row < 8; row++ {
			for col := 0; col < 8; col++ {
				if board.boardArray[row][col].color == "white" {
					if contains(getMoves(board, row, col), board.bKing) {
						return true
					}
				}
			}
		}
	} else {
		fmt.Print(board.wKing.Row)
		fmt.Print(board.wKing.Col)
		for row := 0; row < 8; row++ {
			for col := 0; col < 8; col++ {
				if board.boardArray[row][col].color == "black" {
					if contains(getMoves(board, row, col), board.wKing) {
						return true
					}
				}
			}
		}
	}
	return false
}

