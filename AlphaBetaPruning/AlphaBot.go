package AlphaBetaPruning

import (
	"fmt"
	"github.com/m/v2/CoreGameplay"
	"math"
	"strconv"
	"strings"
)

type AlphaBot struct {
	player CoreGameplay.PlayerPiece
}

func (this *AlphaBot) Init(player CoreGameplay.PlayerPiece) {
	evaluateLocationScore(CoreGameplay.NewGameBoard(), CoreGameplay.NewLocation(1, 1), player)
	this.player = player
}

func (this *AlphaBot) GetName() string {
	return "Adam's AlphaBot"
}

func (this *AlphaBot) MakeMove(board CoreGameplay.Board, c chan int) {

	//this.printBoardState(board)
	_, bestColumn := this.determineBestMoveWithDepth(board, 3)
	c <- bestColumn
	//time.Sleep(time.Millisecond * 500)
	return
}

func (this *AlphaBot) determineBestMoveWithDepth(board CoreGameplay.Board, depth int) (bestScore, bestColumn int) {
	if depth == 0 {
		return this.evaluateBoardState(board), -1
	}

	bestColumn = -1
	bestScore = math.MinInt

	for column := 0; column < CoreGameplay.NumColumns; column++ {
		newBoard := board.Copy()
		if !newBoard.AddPiece(this.player, column) {
			continue
		}
		if newBoard.IsWinningState() == this.player {
			//Instant win, do this move
			return 10000000000, column
		}

		worstPositionScore := math.MaxInt

		for enemyColumn := 0; enemyColumn < CoreGameplay.NumColumns; enemyColumn++ {
			enemyBoard := newBoard.Copy()
			enemyBoard.AddPiece(CoreGameplay.OtherPlayer(this.player), enemyColumn)
			if enemyBoard.IsWinningState() == CoreGameplay.OtherPlayer(this.player) {
				worstPositionScore = math.MinInt
				break
			}
			score, _ := this.determineBestMoveWithDepth(enemyBoard.Copy(), depth-1)
			if score < worstPositionScore {
				worstPositionScore = score
			}
		}
		if worstPositionScore > bestScore {
			bestColumn = column
			bestScore = worstPositionScore
		}
	}
	return bestScore, bestColumn
}

func (this *AlphaBot) evaluateBoardState(board CoreGameplay.Board) int {

	positiveScore := 0
	negativeScore := 0

	for y := 0; y < CoreGameplay.NumRows; y++ {
		for x := 0; x < CoreGameplay.NumColumns; x++ {
			positiveScore += evaluateLocationScore(board, CoreGameplay.NewLocation(x, y), this.player)
			negativeScore += evaluateLocationScore(board, CoreGameplay.NewLocation(x, y), CoreGameplay.OtherPlayer(this.player))
		}
	}

	return positiveScore - negativeScore
}

func (this *AlphaBot) printBoardState(board CoreGameplay.Board) {

	var sb strings.Builder

	for y := CoreGameplay.NumRows - 1; y >= 0; y-- {
		for x := 0; x < CoreGameplay.NumColumns; x++ {

			positiveScore := evaluateLocationScore(board, CoreGameplay.NewLocation(x, y), this.player)
			negativeScore := evaluateLocationScore(board, CoreGameplay.NewLocation(x, y), CoreGameplay.OtherPlayer(this.player))

			sb.WriteString("|")
			switch board.GetPieceAtLocation(CoreGameplay.NewLocation(x, y)) {
			case CoreGameplay.NoPlayer:

				sb.WriteString(strconv.Itoa(positiveScore - negativeScore))
			case CoreGameplay.Player1:
				sb.WriteString("X")
			case CoreGameplay.Player2:
				sb.WriteString("O")
			}
		}
		sb.WriteString("|\n")
	}
	fmt.Println(sb.String())
}

func evaluateLocationScore(board CoreGameplay.Board, location CoreGameplay.Location, piece CoreGameplay.PlayerPiece) int {

	if board.GetPieceAtLocation(location) != CoreGameplay.NoPlayer {
		return 0
	}

	bestScore := 0

	for direction := CoreGameplay.Up; direction <= CoreGameplay.DownLeft; direction++ {
		score := 0

		workingLocation := CoreGameplay.NewLocation(location.Column, location.Row)

		for i := 0; i < 3; i++ {
			workingLocation = CoreGameplay.MoveLocationInDirection(workingLocation, direction)
			if !workingLocation.IsValid() {
				score = 0
				break
			}
			locationPiece := board.GetPieceAtLocation(workingLocation)

			if locationPiece == piece {
				score++
			} else if locationPiece == CoreGameplay.OtherPlayer(piece) {
				score = 0
				break
			}
		}
		if score > bestScore {
			bestScore = score
		}

	}

	return scoreToScaledScore(bestScore)
}

func scoreToScaledScore(score int) int {
	switch score {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 7
	case 3:
		return 25
	case 4:
		return 10000000000
	}
	return 0
}
