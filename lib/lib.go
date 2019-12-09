package lib

import (
	"io/ioutil"
	"strconv"
	"strings"
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

//CSInts parses CSV of ints from the first line of input
func CSInts(input string) []int64 {
	var output []int64
	input = strings.TrimSpace(input)
	input = strings.Split(input, "\n")[0]
	for _, s := range strings.Split(input, ",") {
		s = strings.TrimSpace(s)
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		output = append(output, int64(v))
	}
	return output
}

//IntSlicePermutations returns all permutations of a slice of ints
func IntSlicePermutations(arr []int64)[][]int64{
	var helper func([]int64, int64)
	res := [][]int64{}

	helper = func(arr []int64, n int64){
		if n == 1{
			tmp := make([]int64, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := int64(0); i < n; i++{
				helper(arr, n - 1)
				if n % 2 == 1{
					tmp := arr[i]
					arr[i] = arr[n - 1]
					arr[n - 1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n - 1]
					arr[n - 1] = tmp
				}
			}
		}
	}
	helper(arr, int64(len(arr)))
	return res
}

//StringDigits returns a list of digits in a string omitting any non-digit characters
func StringDigits(str string) []int {
	result := make([]int, 0, len(str))
	for _, r := range str {
		i := int(r - 48)
		if i >= 0 && i <= 9 {
			result = append(result, i)
		}
	}
	return result
}
