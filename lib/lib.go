package lib

import (
	"io/ioutil"
)

//IntDigits repeats the digits from n in base 10
func IntDigits(n int) []int {
	if n < 0 {
		n = n * -1
	}
	var digs []int
	for {
		dig := n % 10
		digs = append(digs, dig)
		n = n / 10
		if n == 0 {
			break
		}
	}
	result := make([]int, len(digs))
	for i, dig := range digs {
		result[len(digs)-i-1] = dig
	}
	return result
}

func ReverseInts(input []int) []int {
	output := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		j := (len(input) - i) - 1
		output[j] = input[i]
	}
	return output
}

//RepeatInts finds repeats in input and returns a map of starting position to length
func RepeatInts(input []int) map[int]int {
	repeats := make(map[int]int, len(input))
	for i := 0; i < len(input); {
		dig := input[i]
		length := 1
		for {
			if i+length == len(input) {
				break
			}
			if input[i+length] != dig {
				break
			}
			length++
		}
		if length > 1 {
			repeats[i] = length
		}
		i += length
	}
	return repeats
}

//RepeatAtIndex returns the starting point and length of the digit repeat i is in
//if there is no repeat, returns i,1
func RepeatAtIndex(sl []int, i int) (start, length int) {
	if i >= len(sl) {
		return i, 0
	}
	for jStart, jLength := range RepeatInts(sl) {
		if i < jStart {
			continue
		}
		if i > jStart+jLength {
			continue
		}
		return jStart, jLength
	}
	return i, 1
}

func MustReadFile(file string) []byte {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return b
}
