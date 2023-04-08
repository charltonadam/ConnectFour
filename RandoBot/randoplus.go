package RandoBot

import (
	"github.com/m/v2/CoreGameplay"
	"math/rand"
	"time"
)

type RandoPlus struct {
	player CoreGameplay.PlayerPiece
}

func (this *RandoPlus) Init(player CoreGameplay.PlayerPiece) {
	this.player = player
}

func (this *RandoPlus) GetName() string {
	return "RandoPlus"
}

func (this *RandoPlus) MakeMove(board CoreGameplay.Board, c chan int) {

	columnPick := this.CheckForWinBoardStates(board)
	if columnPick != -1 {
		c <- columnPick
		return
	}

	columnPick = this.CheckForLoseBoardState(board)
	if columnPick != -1 {
		c <- columnPick
		return
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	for true {
		i := r.Intn(CoreGameplay.NumColumns)
		if board.CanAddPieceAtColumn(i) {
			c <- i
			return
		}
	}
}

func (this *RandoPlus) CheckForWinBoardStates(board CoreGameplay.Board) int {
	for i := 0; i < CoreGameplay.NumColumns; i++ {
		newBoard := board.Copy()
		valid := newBoard.AddPiece(this.player, i)
		if valid && newBoard.IsWinningState() != CoreGameplay.NoPlayer {
			return i
		}
	}
	return -1
}

func (this *RandoPlus) CheckForLoseBoardState(board CoreGameplay.Board) int {
	for i := 0; i < CoreGameplay.NumColumns; i++ {
		newBoard := board.Copy()
		valid := newBoard.AddPiece(CoreGameplay.OtherPlayer(this.player), i)
		if valid && newBoard.IsWinningState() != CoreGameplay.NoPlayer {
			return i
		}
	}
	return -1
}
