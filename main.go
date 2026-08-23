package main

import (
	"chess/pieces"
	"fmt"
	// "net/http"
	// "strconv"
)

var testingBoard pieces.Board

func main() {
	var row int
	var col int 

	testingBoard = pieces.NewBoard()
	testingBoard.ShowBoard()
	fmt.Print("Enter row") 
	fmt.Scan(&row)
	fmt.Print("Enter col") 
	fmt.Scan(&col)
	fmt.Print(pieces.ShowMoves(testingBoard,row, col))
//website interface
	// testingBoard = pieces.NewBoard()
	// testingBoard.ShowBoard()
	// http.HandleFunc("/legal-moves", legalMovesHandler)

	// fmt.Println("Server running on http://localhost:8080")

	// http.ListenAndServe(":8080", nil)
}

// func legalMovesHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
// 	// Get the query parameters
// 	rowString := r.URL.Query().Get("row")
// 	colString := r.URL.Query().Get("col")

// 	// Convert them from strings to integers
// 	row, err := strconv.Atoi(rowString)
// 	if err != nil {
// 		http.Error(w, "Invalid row", http.StatusBadRequest)
// 		return
// 	}

// 	col, err := strconv.Atoi(colString)
// 	if err != nil {
// 		http.Error(w, "Invalid column", http.StatusBadRequest)
// 		return
// 	}

// 	testingBoard.PieceLocation(row, col)
// 	piece := testingBoard.ReturnPiece(row, col).Name()
//     pieces.ShowMoves(testingBoard, row, col, piece)

// 	// For now, just verify that Go received them
// 	fmt.Fprintf(w, "Row: %d, Column: %d", row, col)
// }
