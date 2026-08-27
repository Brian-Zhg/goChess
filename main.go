package main

import (
	"chess/pieces"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
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

	fmt.Println("Server running on http://localhost:8080")

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
    fmt.Fprint(w, pieces.ShowMoves(&testingBoard, row, col))
	moves := pieces.ShowMoves(&testingBoard, row, col)

	jsonData, err := json.Marshal(moves)

	fmt.Fprint(w, jsonData)
	// For now, just verify that Go received them
	fmt.Fprintf(w, "Row: %d, Column: %d", row, col)
}
