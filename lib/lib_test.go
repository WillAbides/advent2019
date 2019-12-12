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


func TestReduceRatio(t *testing.T) {
	for _, td := range []struct {
		numerator, denominator, wantNumerator, wantDenominator int
	}{
		{
			numerator:       2,
			denominator:     6,
			wantNumerator:   1,
			wantDenominator: 3,
		},
		{
			numerator:       4,
			denominator:     6,
			wantNumerator:   2,
			wantDenominator: 3,
		},
		{
			numerator:       24,
			denominator:     36,
			wantNumerator:   2,
			wantDenominator: 3,
		},
		{
			numerator:       -24,
			denominator:     36,
			wantNumerator:   -2,
			wantDenominator: 3,
		},
		{
			numerator:       24,
			denominator:     -36,
			wantNumerator:   2,
			wantDenominator: -3,
		},
		{
			numerator:       1,
			denominator:     -36,
			wantNumerator:   1,
			wantDenominator: -36,
		},
		{
			numerator:       0,
			denominator:     -36,
			wantNumerator:   0,
			wantDenominator: -1,
		},
		{
			numerator:       -24,
			denominator:     0,
			wantNumerator:   -1,
			wantDenominator: 0,
		},
	} {
		t.Run(fmt.Sprintf("%d/%d", td.numerator, td.denominator), func(t *testing.T) {
			gotN, gotD := ReduceRatio(td.numerator, td.denominator)
			assert.Equal(t, td.wantNumerator, gotN)
			assert.Equal(t, td.wantDenominator, gotD)
		})
	}
}

func TestPrimeFactors(t *testing.T) {
	for _, td := range []struct {
		input uint
		want  []uint
	}{
		{
			input: 2,
			want:  []uint{2},
		},
		{
			input: 7,
			want:  []uint{7},
		},
		{
			input: 23,
			want:  []uint{23},
		},
		{
			input: 12,
			want:  []uint{2, 2, 3},
		},
		{
			input: 360,
			want:  []uint{2, 2, 2, 3, 3, 5},
		},
		{
			input: 1067,
			want:  []uint{97, 11},
		},
	} {
		t.Run(fmt.Sprintf("%d", td.input), func(t *testing.T) {
			got := PrimeFactors(td.input)
			assert.ElementsMatch(t, td.want, got)
		})
	}
}
