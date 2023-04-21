package main

import (
	"github.com/m/v2/AlphaBetaPruning"
	"github.com/m/v2/CoreGameplay"
	"github.com/m/v2/Game"
	"github.com/m/v2/HumanPlayer"
	"github.com/m/v2/RandoBot"
)

func main() {
	players := []CoreGameplay.Player{
		&HumanPlayer.HumanPlayer{},
		&RandoBot.RandoBot{},
		&RandoBot.RandoPlus{},
		&AlphaBetaPruning.AlphaBot{},
	}

	Game.NewGame(players)
}
