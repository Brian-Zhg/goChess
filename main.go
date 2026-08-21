package main

import (
    "chess/pieces"
    "net/http"
    "fmt"
)

func main() {
    testingBoard := pieces.NewBoard()
    testingBoard.ShowBoard()
    http.HandleFunc("/test", testHandler)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}


func testHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	fmt.Fprint(w, {"Hello from Go!": "hello"})
}
