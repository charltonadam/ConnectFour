package HumanPlayer

import (
	"bufio"
	"fmt"
	"github.com/m/v2/CoreGameplay"
	"os"
	"strconv"
	"strings"
)

type HumanPlayer struct {
	player CoreGameplay.PlayerPiece
	name   string
}

func (this *HumanPlayer) Init(player CoreGameplay.PlayerPiece) {
	this.player = player
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Player Name: ")
	text, _ := reader.ReadString('\n')
	this.name = text
}

func (this *HumanPlayer) GetName() string {
	if this.name == "" {
		return "Human Player"
	}
	return this.name
}

func (this *HumanPlayer) MakeMove(_ CoreGameplay.Board, c chan int) {
	fmt.Print("Enter column 0-6: ")
	reader := bufio.NewReader(os.Stdin)

	for true {
		text, _ := reader.ReadString('\n')
		i, err := strconv.Atoi(strings.Trim(text, "\n"))
		if err == nil && i >= 0 && i <= CoreGameplay.NumColumns {
			c <- i
			break
		} else {
			fmt.Print("Invalid column, please try again: ")
		}
	}
}
