package CoreGameplay

import "strings"

type PlayerPiece byte

const (
	NoPlayer PlayerPiece = iota
	Player1
	Player2

	NumRows    int = 6
	NumColumns int = 7
)

func OtherPlayer(player PlayerPiece) PlayerPiece {
	if player == Player1 {
		return Player2
	} else if player == Player2 {
		return Player1
	} else {
		return NoPlayer
	}
}

type Board [][]PlayerPiece

func (this Board) GetNumColumns() int {
	return NumColumns
}

func (this Board) GetPieceAtLocation(location Location) PlayerPiece {
	return this[location.Row][location.Column]
}

func NewGameBoard() Board {
	board := make([][]PlayerPiece, NumRows)
	for i := range board {
		board[i] = make([]PlayerPiece, NumColumns)
	}
	return board
}

func (this Board) AddPiece(player PlayerPiece, column int) (Board, bool) {
	for i := 0; i < NumRows; i++ {
		if this[i][column] == NoPlayer {
			this[i][column] = player
			return this, true
		}
	}
	return this, false
}

func (this Board) IsWinningState() PlayerPiece {
	for i := 0; i < NumRows; i++ {
		for j := 0; j < NumColumns; j++ {
			player := IsPositionWinningMove(this, NewLocation(i, j))
			if player != NoPlayer {
				return player
			}
		}
	}
	return NoPlayer
}

func (this Board) CanAddPieceAtColumn(column int) bool {
	return this.GetPieceAtLocation(NewLocation(NumRows-1, column)) == NoPlayer
}

func IsBoardWinningState(board Board) PlayerPiece {
	for i := 0; i < NumRows; i++ {
		for j := 0; j < NumColumns; j++ {
			player := IsPositionWinningMove(board, NewLocation(i, j))
			if player != NoPlayer {
				return player
			}
		}
	}
	return NoPlayer
}

func IsPositionWinningMove(board Board, loc Location) PlayerPiece {
	player := board.GetPieceAtLocation(loc)
	if player == NoPlayer {
		return NoPlayer
	}
	for i := up; i <= downLeft; i++ {
		if isWinningInDirection(board, loc, i, player, 0) {
			return player
		}
	}
	return NoPlayer
}

func isWinningInDirection(board Board, loc Location, dir Direction, player PlayerPiece, amount int) bool {
	if !loc.IsValid() {
		return false
	}

	if board.GetPieceAtLocation(loc) != player {
		return false
	}

	if amount == 3 {
		return true
	}
	loc = MoveLocationInDirection(loc, dir)

	return isWinningInDirection(board, loc, dir, player, amount+1)
}

func (this Board) ToString() string {
	var sb strings.Builder
	for i := NumRows - 1; i >= 0; i-- {
		for j := 0; j < NumColumns; j++ {
			sb.WriteString("|")
			switch this.GetPieceAtLocation(NewLocation(i, j)) {
			case NoPlayer:
				sb.WriteString(" ")
			case Player1:
				sb.WriteString("X")
			case Player2:
				sb.WriteString("O")
			}
		}
		sb.WriteString("|\n")
	}
	return sb.String()
}
