package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt_Add(t *testing.T) {
	tests := []struct {
		name              string
		values            []int
		expectedStringRep string
		expectedHeadValue int
		expectedTailValue int
	}{
		{
			"add 10 numbers",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			"1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> nil",
			1,
			10,
		},
		{
			"add 1 numbers",
			[]int{1},
			"1 -> nil",
			1,
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intList := &Int{}
			for _, val := range tt.values {
				intList.Add(val)
			}
			assert.Equal(t, tt.expectedStringRep, intList.String())
		})
	}
}
