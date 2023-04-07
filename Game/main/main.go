package main

import (
	"github.com/m/v2/CoreGameplay"
	"github.com/m/v2/Game"
	"github.com/m/v2/HumanPlayer"
	"github.com/m/v2/RandoBot"
)

func main() {
	players := make([]CoreGameplay.Player, 0)
	players = append(players, &HumanPlayer.HumanPlayer{})
	players = append(players, &RandoBot.RandoBot{})

	Game.NewGame(players)

}
