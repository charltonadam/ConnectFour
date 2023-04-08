package CoreGameplay

import "errors"

type Player interface {
	Init(piece PlayerPiece)
	GetName() string
	MakeMove(Board, chan int)
}

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Add(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var defaultVal T
	if len(s.data) == 0 {
		return defaultVal, errors.New("Stack empty")
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, nil
}
