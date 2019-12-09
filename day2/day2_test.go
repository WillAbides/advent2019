package day2

import (
	"fmt"
	"testing"

	"github.com/WillAbides/advent2019/lib/intcomputer"

	"github.com/stretchr/testify/assert"
)

func calcNounVerb(noun, verb int64) (int64, error) {
	codes := []int64{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 2, 13, 23, 27, 2, 27, 13, 31, 2, 31, 10, 35, 1, 6, 35, 39, 1, 5, 39, 43, 1, 10, 43, 47, 1, 5, 47, 51, 1, 13, 51, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 13, 63, 67, 1, 6, 67, 71, 1, 71, 10, 75, 2, 13, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 6, 87, 91, 1, 91, 13, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 5, 107, 2, 107, 10, 111, 1, 5, 111, 115, 1, 2, 115, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0}
	codes[1] = noun
	codes[2] = verb
	c := intcomputer.NewIntComputer(codes, nil, nil)
	err := c.RunOperations()
	if err != nil {
		return 0, err
	}
	return c.ReadPosition(0), nil
}

func TestStage2(t *testing.T) {
	var noun, verb, result, want int64
	want = 19690720
	var err error
	for result != want {
		if want-result > 300000 {
			noun += 1
		} else {
			verb += 1
		}
		result, err = calcNounVerb(noun, verb)
		if err != nil {
			fmt.Printf("got error with %d, %d\n", noun, verb)
		}
	}
	fmt.Printf("%d%d\n", noun, verb)
}

func TestMoveNouns(t *testing.T) {
	for i := 0; i < 10; i++ {
		got, err := calcNounVerb(int64(i), 0)
		if err != nil {
			fmt.Print("error: ")
		}
		fmt.Println(got)
	}
}

func TestMoveVerbs(t *testing.T) {
	for i := 0; i < 10; i++ {
		got, err := calcNounVerb(2, int64(i))
		if err != nil {
			fmt.Print("error: ")
		}
		fmt.Println(got)
	}
}

func Test_go(t *testing.T) {
	input := []int64{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 2, 13, 23, 27, 2, 27, 13, 31, 2, 31, 10, 35, 1, 6, 35, 39, 1, 5, 39, 43, 1, 10, 43, 47, 1, 5, 47, 51, 1, 13, 51, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 13, 63, 67, 1, 6, 67, 71, 1, 71, 10, 75, 2, 13, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 6, 87, 91, 1, 91, 13, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 5, 107, 2, 107, 10, 111, 1, 5, 111, 115, 1, 2, 115, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0}
	input[1] = 12
	input[2] = 2
	c := intcomputer.NewIntComputer(input, nil, nil)
	assert.NoError(t, c.RunOperations())
	assert.Equal(t, int64(3790689), c.ReadPosition(0))
}
