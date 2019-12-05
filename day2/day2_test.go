package day2

import (
	"fmt"
	"testing"

	"advent2019"

	"github.com/stretchr/testify/assert"
)

func procCodes2(codes []int) []int {
	c := &advent2019.IntComputer{
		Memory: codes,
	}
	c.RunOperations()
	return c.Memory
}

func calcNounVerb(noun, verb int) (int, error) {
	codes := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 2, 13, 23, 27, 2, 27, 13, 31, 2, 31, 10, 35, 1, 6, 35, 39, 1, 5, 39, 43, 1, 10, 43, 47, 1, 5, 47, 51, 1, 13, 51, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 13, 63, 67, 1, 6, 67, 71, 1, 71, 10, 75, 2, 13, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 6, 87, 91, 1, 91, 13, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 5, 107, 2, 107, 10, 111, 1, 5, 111, 115, 1, 2, 115, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0}
	codes[1] = noun
	codes[2] = verb
	c := &advent2019.IntComputer{
		Memory: codes,
	}
	c.RunOperations()
	return c.ReadPosition(0), nil
}

func TestStage2(t *testing.T) {
	noun := 0
	verb := 0
	want := 19690720
	result := 0
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
		got, err := calcNounVerb(i, 0)
		if err != nil {
			fmt.Print("error: ")
		}
		fmt.Println(got)
	}
}

func TestMoveVerbs(t *testing.T) {
	for i := 0; i < 10; i++ {
		got, err := calcNounVerb(2, i)
		if err != nil {
			fmt.Print("error: ")
		}
		fmt.Println(got)
	}
}

func Test_go(t *testing.T) {
	input := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 2, 13, 23, 27, 2, 27, 13, 31, 2, 31, 10, 35, 1, 6, 35, 39, 1, 5, 39, 43, 1, 10, 43, 47, 1, 5, 47, 51, 1, 13, 51, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 13, 63, 67, 1, 6, 67, 71, 1, 71, 10, 75, 2, 13, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 6, 87, 91, 1, 91, 13, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 5, 107, 2, 107, 10, 111, 1, 5, 111, 115, 1, 2, 115, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0}
	input[1] = 12
	input[2] = 2
	codes := procCodes2(input)
	fmt.Println(codes[0])
}

func Test_procCodes(t *testing.T) {
	for _, td := range []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 4, 4, 0, 99},
			[]int{198, 4, 4, 0, 99},
		},
		{
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		},
		{
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		}, {
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		}, {
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	} {
		t.Run("", func(t *testing.T) {
			got := procCodes2(td.input)
			assert.Equal(t, td.want, got)
		})
	}
}
