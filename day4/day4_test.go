package day4

import (
	"fmt"
	"testing"

	"github.com/WillAbides/advent2019/lib"

	"github.com/stretchr/testify/assert"
)

func consecutiveRule(dgts []int) bool {
	for i := 1; i < len(dgts); i++ {
		if dgts[i] == dgts[i-1] {
			return true
		}
	}
	return false
}

func consecutiveRule2(dgts []int) bool {
	for i := 1; i < len(dgts); i++ {
		if dgts[i] == dgts[i-1] {
			if !isInTrip(i, dgts) {
				return true
			}
		}
	}
	return false
}

func isInTrip(i int, dgts []int) bool {
	_, repeatLength := lib.RepeatAtIndex(dgts, i)
	return repeatLength > 2
}

func Test_isInTrip(t *testing.T) {
	dgts := lib.IntDigits(123444)
	assert.True(t, isInTrip(5, dgts))
	assert.True(t, isInTrip(4, dgts))
	assert.True(t, isInTrip(3, dgts))
	assert.False(t, isInTrip(2, dgts))
	assert.False(t, isInTrip(1, dgts))
	assert.False(t, isInTrip(0, dgts))
}

func neverDecreaseRule(dgts []int) bool {
	for i := 1; i < len(dgts); i++ {
		if dgts[i] < dgts[i-1] {
			return false
		}
	}
	return true
}

func Test_neverDecreaseRule(t *testing.T) {
	assert.True(t, neverDecreaseRule([]int{1, 2, 3, 4, 5, 6}))
	assert.True(t, neverDecreaseRule([]int{1, 2, 2, 4, 5, 6}))
	assert.False(t, neverDecreaseRule([]int{1, 2, 2, 4, 5, 4}))
}

func Test_consecutiveRule(t *testing.T) {
	assert.False(t, consecutiveRule([]int{1, 2, 3, 4, 5, 6}))
	assert.True(t, consecutiveRule([]int{1, 2, 2, 4, 5, 6}))
	assert.True(t, consecutiveRule([]int{1, 2, 2, 4, 5, 4}))
}

func calcCombo(n int) bool {
	dgts := lib.IntDigits(n)
	if !consecutiveRule(dgts) {
		return false
	}
	if !neverDecreaseRule(dgts) {
		return false
	}
	return true
}

func calcCombo2(n int) bool {
	dgts := lib.IntDigits(n)
	if !consecutiveRule2(dgts) {
		return false
	}
	if !neverDecreaseRule(dgts) {
		return false
	}
	return true
}

func TestEx1(t *testing.T) {
	var count int
	for i := 134792; i <= 675810; i++ {
		if calcCombo(i) {
			count++
		}
	}
	fmt.Println(count)
}

func TestEx2(t *testing.T) {
	var count int
	for i := 134792; i <= 675810; i++ {
		if calcCombo2(i) {
			count++
		}
	}
	fmt.Println(count)
}

func TestCalcCombo2(t *testing.T) {
	assert.True(t, calcCombo2(112233))
	assert.False(t, calcCombo2(123444))
	assert.True(t, calcCombo2(111122))
}

func TestCalcCombo(t *testing.T) {
	assert.True(t, calcCombo(111111))
	assert.False(t, calcCombo(223450))
	assert.False(t, calcCombo(123789))
}
