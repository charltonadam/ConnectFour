package CoreGameplay

type Player interface {
	Init(piece PlayerPiece)
	GetName() string
	MakeMove(Board, chan int)
}
