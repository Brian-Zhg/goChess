const board = document.getElementById("board");
var prevPiece;
var prevSquare;
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

    square.addEventListener("click", handleSquareClick);
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
    square.appendChild(piece);
}

//when pieces are clicked
async function handlePieceClick(piece) {
    if (prevSquare != null) {
        prevSquare.classList.remove("pressed");
        prevSquare.classList.remove("take");
    }
    if (prevPiece != null) {
        prevPiece.classList.remove("pressed");
        prevPiece.classList.remove("take");
    }
    clearPreviousMoves();
    const response = await fetch(
        `http://127.0.0.1:8080/turn`
    );
    const turn = await response.json()
    if (turn == true && piece.classList[1].includes("black") || turn == false && piece.classList[1].includes("white")) {
        prevPiece = piece;
        const row = piece.dataset.row;
        const col = piece.dataset.col;
        prevSquare = piece.parentElement;
        piece.parentElement.classList.add("pressed");
        const data = await getLegalMoves(row, col);
        highlightMoves(data);
        prevMoves = data;
    }
}

async function handleSquareClick(event) {
    const square = event.currentTarget;

    // If this square is a legal move, move the piece
    if (prevMoves?.some(
        move =>
            Number(move.Row) === Number(square.dataset.row) &&
            Number(move.Col) === Number(square.dataset.col)
    )) {
        await handleMoveClick(event);
        return;
    }

    // Otherwise, check if there's a piece on this square
    const piece = square.querySelector(".piece");

    if (piece) {
        await handlePieceClick(piece);
    }
}

//deleting previous highlighted moves
function clearPreviousMoves() {
    if (prevMoves != null) {
        prevMoves.forEach(move => {
            squares[move.Row][move.Col].classList.remove("pressed");
            squares[move.Row][move.Col].classList.remove("take");
            squares[move.Row][move.Col].removeEventListener('click', handleMoveClick);
        });
    }
}

//returns legal moves of piece
async function getLegalMoves(row, col) {
    const response = await fetch(
        `http://127.0.0.1:8080/legal-moves?row=${row}&col=${col}`
    );
    const data = await response.json();
    return data;
}

//highlight the moves that the piece can make
function highlightMoves(data) {
    data.forEach(move => {
        if (squares[move.Row][move.Col].querySelector(".piece")) {
            squares[move.Row][move.Col].classList.add("take");
        }
        else squares[move.Row][move.Col].classList.add("pressed");
        squares[move.Row][move.Col].addEventListener("click", handleMoveClick);
    });
}

//what happens when you click the highlighted square
async function handleMoveClick(event) {
    const square = event.currentTarget;
    const row = square.dataset.row;
    const col = square.dataset.col;
    const response = await fetch(
        `http://127.0.0.1:8080/move?row=${row}&col=${col}&moveR=${prevPiece.dataset.row}&moveC=${prevPiece.dataset.col}`
    );
    const data = await response.json();
    if (data == true) {
        movePiece(prevPiece.dataset.row, prevPiece.dataset.col, row, col);
    }
}

function movePiece(movingR, movingC, newR, newC) {
    const oldSquare = squares[movingR][movingC];
    const newSquare = squares[newR][newC];
    // Get the actual piece
    oldSquare.classList.remove("pressed");
    oldSquare.classList.remove("take");
    const piece = oldSquare.querySelector(".piece");
    const existingPiece = newSquare.querySelector(".piece");
    if (existingPiece) {
        existingPiece.remove();
    }
    // Move the piece visually
    newSquare.appendChild(piece);
    // Update the piece's coordinates
    piece.dataset.row = newR;
    piece.dataset.col = newC;
    clearPreviousMoves();
}