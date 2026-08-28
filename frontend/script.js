const board = document.getElementById("board");
var prevPiece;

const boardState = [
    ["black-rook", "black-knight", "black-bishop", "black-queen",
        "black-king", "black-bishop", "black-knight", "black-rook"],

    ["black-pawn", "black-pawn", "black-pawn", "black-pawn",
        "black-pawn", "black-pawn", "black-pawn", "black-pawn"],

    [null, null, null, null, null, null, null, null],
    [null, null, null, null, null, null, null, null],
    [null, null, null, null, null, null, null, null],
    [null, null, null, null, null, null, null, null],

    ["white-pawn", "white-pawn", "white-pawn", "white-pawn",
        "white-pawn", "white-pawn", "white-pawn", "white-pawn"],

    ["white-rook", "white-knight", "white-bishop", "white-queen",
        "white-king", "white-bishop", "white-knight", "white-rook"]
];

const pieces = {
    "white-king": "♔",
    "white-queen": "♕",
    "white-rook": "♖",
    "white-bishop": "♗",
    "white-knight": "♘",
    "white-pawn": "♙",

    "black-king": "♚",
    "black-queen": "♛",
    "black-rook": "♜",
    "black-bishop": "♝",
    "black-knight": "♞",
    "black-pawn": "♟"
};
for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {

        const square = document.createElement("div");
        square.classList.add("square");
        square.classList.add("light");
        if ((row + col) % 2 === 0) {
            square.classList.add("light");
            square.dataset.row = row;
            square.dataset.col = col;
        } else {
            square.classList.add("dark");
            square.dataset.row = row;
            square.dataset.col = col;
        }

        const pieceName = boardState[row][col];

        if (pieceName !== null) {

            const piece = document.createElement("div");
            piece.classList.add("piece");
            piece.classList.add(pieceName);
            piece.dataset.row = row;
            piece.dataset.col = col;
            // Put the actual chess symbol inside the element
            piece.textContent = pieces[pieceName];
            piece.addEventListener("click", async function () {
                if (prevPiece != null) prevPiece.classList.remove("pressed");
                prevPiece = square
                const row = piece.dataset.row;
                const col = piece.dataset.col;
                square.classList.add("pressed");
                console.log(boardState[row][col]);

                const response = await fetch(
                    `http://127.0.0.1:8080/legal-moves?row=${row}&col=${col}`
                );

                const data = await response.text();

                console.log(data);
                //console.log(data.row);
                // Eventually you'll receive the legal moves here
            });

            square.appendChild(piece);
        }

        board.appendChild(square);
    }
}


