package main

import (
    "fmt"
    "chess/pieces"
)

func main() {
    fmt.Print("_")
    testingBoard := pieces.Board{}
    testingBoard.ShowBoard()
}