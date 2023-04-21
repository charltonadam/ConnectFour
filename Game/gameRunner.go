package Game

import (
	"bufio"
	"context"
	"fmt"
	"github.com/m/v2/CoreGameplay"
	"os"
	"strconv"
	"time"
)

type gameState struct {
	board            CoreGameplay.Board
	player1, player2 CoreGameplay.Player
	turn             int
}

func NewGame(players []CoreGameplay.Player) {
	game := &gameState{
		board:   CoreGameplay.NewGameBoard(),
		player1: selectPlayer(players),
		player2: selectPlayer(players),
		turn:    0,
	}
	game.player1.Init(CoreGameplay.Player1)
	game.player2.Init(CoreGameplay.Player2)

	for game.turn < CoreGameplay.NumRows*CoreGameplay.NumColumns {
		player := takeTurn(game)
		if player == CoreGameplay.NoPlayer {

			game.turn++
			fmt.Print(game.board.ToString() + "\n\n")
		}
		if player == CoreGameplay.Player1 {
			fmt.Printf("Player 1: %s, WINS!!!\n", game.player1.GetName())
			fmt.Print(game.board.ToString() + "\n\n")
			return
		}
		if player == CoreGameplay.Player2 {
			fmt.Printf("Player 2: %s, WINS!!!\n", game.player2.GetName())
			fmt.Print(game.board.ToString() + "\n\n")
			return
		}

	}
	fmt.Println("TIE GAME!!!")

}

func selectPlayer(players []CoreGameplay.Player) CoreGameplay.Player {
	reader := bufio.NewReader(os.Stdin)

	for true {
		listPlayers(players)
		fmt.Print("Select player: ")
		text, _ := reader.ReadString('\n')
		selection, err := strconv.Atoi(text[:len(text)-1])
		if err == nil && selection < len(players) && selection >= 0 {
			return players[selection]
		}
		fmt.Print("Invalid Selection!")
	}
	return nil
}

func listPlayers(players []CoreGameplay.Player) {
	for i, player := range players {
		fmt.Printf("%d: %s\n", i, player.GetName())
	}
}

func takeTurn(state *gameState) CoreGameplay.PlayerPiece {
	if state.turn%2 == 0 {
		//player1
		state.board.AddPiece(CoreGameplay.Player1, playerMove(state.player1, state.board))
	} else {
		//player2
		state.board.AddPiece(CoreGameplay.Player2, playerMove(state.player2, state.board))
	}
	return state.board.IsWinningState()
}

func playerMove(player CoreGameplay.Player, board CoreGameplay.Board) int {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	playerMoveChannel := make(chan int, 1)
	go func() {
		player.MakeMove(board, playerMoveChannel)
		cancel()
	}()
	move := -1
	func() {
		for {
			select {
			case <-ctx.Done():
				return
			case move = <-playerMoveChannel:
				continue
			}
		}
	}()

	if move != -1 && board.CanAddPieceAtColumn(move) {
		return move
	}

	//else, go down the list until you find a valid move
	for i := 0; i < CoreGameplay.NumColumns; i++ {
		if board.CanAddPieceAtColumn(i) {
			return i
		}
	}
	return -1

}
