package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntDigits(t *testing.T) {
	for _, td := range []struct {
		input int
		want  []int
	}{
		{0, []int{0}},
		{123, []int{1, 2, 3}},
		{-123, []int{1, 2, 3}},
		{65432, []int{6, 5, 4, 3, 2}},
	} {
		t.Run(fmt.Sprintf("%d", td.input), func(t *testing.T) {
			assert.Equal(t, td.want, IntDigits(td.input))
		})
	}
}

func TestRepeatInts(t *testing.T) {
	for _, td := range []struct {
		input []int
		want  map[int]int
	}{
		{[]int{}, map[int]int{}},
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1}, map[int]int{}},
		{[]int{0, 1, 2, 3, 4, 5, 5, 4, 3, 2, 1, 0}, map[int]int{5: 2}},
		{[]int{0, 1, 2, 3, 4, 5, 5, 4, 3, 0, 0, 0}, map[int]int{5: 2, 9: 3}},
	} {
		t.Run(fmt.Sprintf("%d", td.input), func(t *testing.T) {
			assert.Equal(t, td.want, RepeatInts(td.input))
		})
	}
}

func TestRepeatAtIndex(t *testing.T) {
	type args struct {
		sl []int
		i  int
	}
	tests := []struct {
		name       string
		args       args
		wantStart  int
		wantLength int
	}{
		{
			args: struct {
				sl []int
				i  int
			}{sl: []int{1, 2, 3, 3, 3, 4}, i: 4},
			wantLength: 3,
			wantStart:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotLength := RepeatAtIndex(tt.args.sl, tt.args.i)
			if gotStart != tt.wantStart {
				t.Errorf("RepeatAtIndex() gotStart = %v, want %v", gotStart, tt.wantStart)
			}
			if gotLength != tt.wantLength {
				t.Errorf("RepeatAtIndex() gotLength = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}
