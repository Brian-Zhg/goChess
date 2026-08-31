const board = document.getElementById("board");
var prevPiece;
var prevMoves;


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

const squares = Array.from({ length: 8 }, () => Array(8));

for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {
        const square = document.createElement("div");

        createSquare(square, row, col);
        createPiece(square, row, col);

        board.appendChild(square);
    }
}

//creates individual pieces
function createSquare(square, row, col) {
    square.classList.add("square");
    if ((row + col) % 2 === 0) {
        square.classList.add("light");
        square.dataset.row = row;
        square.dataset.col = col;
    } else {
        square.classList.add("dark");
        square.dataset.row = row;
        square.dataset.col = col;
    }
    squares[row][col] = square;
}

//add the actual piece
function createPiece(square, row, col) {
    const pieceName = boardState[row][col];
    if (pieceName == null) {
        return;
    }
    const piece = document.createElement("div");
    piece.classList.add("piece");
    piece.classList.add(pieceName);
    piece.dataset.row = row;
    piece.dataset.col = col;
    piece.textContent = pieces[pieceName];
    piece.addEventListener("click", () => {
        handlePieceClick(piece, square);
    });
    square.appendChild(piece);
}

//when pieces are clicked
async function handlePieceClick(piece, square) {
    if (prevPiece != null) prevPiece.classList.remove("pressed");
    clearPreviousMoves();
    prevPiece = square;
    const row = piece.dataset.row;
    const col = piece.dataset.col;
    square.classList.add("pressed");
    const data = await getLegalMoves(row, col);
    highlightMoves(data);
    prevMoves = data;
}

//deleting previous highlighted moves
function clearPreviousMoves() {
    if (prevMoves != null) {
        prevMoves.forEach(move => {
            squares[move.Row][move.Col].classList.remove("pressed");
            squares[move.Row][move.Col].removeEventListener('click',handleMoveClick);
        });
    }
}

async function getLegalMoves(row, col) {
    const response = await fetch(
        `http://127.0.0.1:8080/legal-moves?row=${row}&col=${col}`
    );
    const data = await response.json();
    return data;
}

function highlightMoves(data) {
    data.forEach(move => {
        squares[move.Row][move.Col].classList.add("pressed");
        squares[move.Row][move.Col].addEventListener("click", handleMoveClick)
    });
}

async function handleMoveClick(event) {
    const square = event.currentTarget;
    const row = square.dataset.row;
    const col = square.dataset.col;
    const response = await fetch(
        `http://127.0.0.1:8080/move?row=${row}&col=${col}&moveR=${prevPiece.Row}&moveC${prevPiece.Col}`
    );
    
    console.log("Row:" + row + " Col:" + col);
}