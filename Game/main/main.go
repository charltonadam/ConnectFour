package main

import (
	"github.com/m/v2/CoreGameplay"
	"github.com/m/v2/EastonBot"
	"github.com/m/v2/Game"
	"github.com/m/v2/HumanPlayer"
	"github.com/m/v2/RandoBot"
)

func main() {
	players := []CoreGameplay.Player{
		&HumanPlayer.HumanPlayer{},
		&RandoBot.RandoBot{},
		&EastonBot.Easton{},
	}

	Game.NewGame(players)
}
