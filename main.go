package main

import (
	"chess/pieces"
	"encoding/json"
	"net/http"
	"strconv"
)

var testingBoard pieces.Board

func main() {
	//testing code
	// var pos1R int
	// var pos1C  int
	// var pos2R int
	// var pos2C  int

	// testingBoard = pieces.NewBoard()
	// testingBoard.ShowBoard()
	// for{
	// fmt.Print("Enter row")
	// fmt.Scan(&pos1R)
	// fmt.Print("Enter col")
	// fmt.Scan(&pos1C)
	// fmt.Print("Enter row")
	// fmt.Scan(&pos2R)
	// fmt.Print("Enter col")
	// fmt.Scan(&pos2C)

	// pieces.ConfirmMove(&testingBoard,pieces.CreatePostion(pos1R,pos1C),pieces.CreatePostion(pos2R,pos2C))
	// testingBoard.ShowBoard()
	// }
	// website interface
	testingBoard = pieces.NewBoard()
	testingBoard.ShowBoard()
	http.HandleFunc("/legal-moves", legalMovesHandler)
	http.HandleFunc("/move", movePiece)
	http.HandleFunc("/turn", turn)

	http.ListenAndServe(":8080", nil)
}

func legalMovesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	// Get the query parameters
	rowString := r.URL.Query().Get("row")
	colString := r.URL.Query().Get("col")

	// Convert them from strings to integers
	row, err := strconv.Atoi(rowString)
	if err != nil {
		http.Error(w, "Invalid row", http.StatusBadRequest)
		return
	}

	col, err := strconv.Atoi(colString)
	if err != nil {
		http.Error(w, "Invalid column", http.StatusBadRequest)
		return
	}

	testingBoard.PieceLocation(row, col)
	moves := pieces.ShowMoves(&testingBoard, row, col)

	jsonData, err := json.Marshal(moves)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func movePiece(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	// Get the query parameters
	rowString := r.URL.Query().Get("row")
	colString := r.URL.Query().Get("col")
	moveToR := r.URL.Query().Get("moveR")
	moveToC := r.URL.Query().Get("moveC")

	row, err := strconv.Atoi(rowString)
	if err != nil {
		http.Error(w, "Invalid row", http.StatusBadRequest)
		return
	}

	col, err := strconv.Atoi(colString)
	if err != nil {
		http.Error(w, "Invalid column", http.StatusBadRequest)
		return
	}

	mRow, err := strconv.Atoi(moveToR)
	if err != nil {
		http.Error(w, "Invalid moving row", http.StatusBadRequest)
		return
	}

	mCol, err := strconv.Atoi(moveToC)
	if err != nil {
		http.Error(w, "Invalid moving column", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(pieces.ConfirmMove(&testingBoard, pieces.Position{Row: mRow, Col: mCol}, pieces.Position{Row: row, Col: col}))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func turn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(testingBoard.ReturnTurn())
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
