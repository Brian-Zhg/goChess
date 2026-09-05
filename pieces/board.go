package pieces

import "fmt"


type Board struct {
	boardArray [8][8]chessPiece
	turn bool
	blackPieces []*chessPiece
	whitePieces []*chessPiece
	bKing *chessPiece
	wKing *chessPiece
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
	
	for i:=0; i < 8; i++{
		if(i == 0 || i ==1 || i == 7 ||i ==8){
			for r:=0; r<8; r++{
				if(i == 0 || i == 1){
					newBoard.blackPieces = append(newBoard.blackPieces, &newBoard.boardArray[i][r])
					if(newBoard.boardArray[i][r].pieceName == "king") {newBoard.bKing = &newBoard.boardArray[i][r]}
				}
				if(i == 7 || i == 8){
					newBoard.whitePieces = append(newBoard.whitePieces, &newBoard.boardArray[i][r])
					if(newBoard.boardArray[i][r].pieceName == "king") {newBoard.wKing = &newBoard.boardArray[i][r]}
				}
		}
		}
		
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
	piece := b.boardArray[row][col].pieceName
	color := b.boardArray[row][col].color
	moves := []Position{}
	if(b.turn == true && color == "white" || b.turn == false && color =="black"){
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
}
	return moves
}

func Move(b *Board, piece1 Position, piece2 Position){
	if(b.boardArray[piece1.Row][piece1.Col].firstMove== false){
		b.boardArray[piece1.Row][piece1.Col].firstMove= true
	}
	b.boardArray[piece2.Row][piece2.Col]= b.boardArray[piece1.Row][piece1.Col]
	b.boardArray[piece1.Row][piece1.Col]=NewChessPiece()
	b.turn = !b.turn
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

func (b Board) ReturnTurn() bool{
	return b.turn
}

func inCheck(board *Board, k king) bool {
	fmt.Print(k.color)
	if k.color == "black" {
		for i := 0; i < len(board.whitePieces); i++{
			if(contains(ShowMoves(board,board.whitePieces[i].pos.Row,board.whitePieces[i].pos.Col),Position{Row:k.pos.Row, Col:k.pos.Col})) {return true}
		}
	}else{
		for i := 0; i < len(board.blackPieces); i++{
			if(contains(ShowMoves(board,board.blackPieces[i].pos.Row,board.blackPieces[i].pos.Col),Position{Row:k.pos.Row, Col:k.pos.Col})) {return true}
		}
	}
	
	return false;
}

