package RandoBot

import (
	"github.com/m/v2/CoreGameplay"
	"math/rand"
	"time"
)

type RandoBot struct {
	player CoreGameplay.PlayerPiece
}

func (this *RandoBot) Init(player CoreGameplay.PlayerPiece) {
	this.player = player
}

func (this *RandoBot) GetName() string {
	return "RandoBot!"
}

func (this *RandoBot) MakeMove(board CoreGameplay.Board, c chan int) {

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
