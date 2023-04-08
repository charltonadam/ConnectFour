package CoreGameplay

import (
	"reflect"
	"testing"
)

func Test_stack_pop(t *testing.T) {
	type testCase[T any] struct {
		name    string
		s       Stack[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{"Simple test", Stack[int]{data: []int{}}, 0, true},
		{"Simple test", Stack[int]{data: []int{2, 3, 4, 1}}, 1, false},
		{"Simple test", Stack[int]{data: []int{1}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lenBefore := len(tt.s.data)
			got, err := tt.s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pop() got = %v, want %v", got, tt.want)
			}
			if err == nil && lenBefore-1 != len(tt.s.data) {
				t.Errorf("pop() didn't shirnk the size of data")
			}
		})
	}
}
