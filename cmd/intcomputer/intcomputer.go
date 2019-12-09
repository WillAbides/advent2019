package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/WillAbides/advent2019/lib"
	"github.com/WillAbides/advent2019/lib/intcomputer"
)

func errOut(message ...interface{}) {
	_, _ = fmt.Fprintln(os.Stdout, message...)
}

func getInput(prompt bool,scanner *bufio.Scanner) (int64, error) {
	if prompt {
		_, _ = fmt.Fprint(os.Stderr, "please input an integer: ")
	}
	scanner.Scan()
	scanErr := scanner.Err()
	if scanErr != nil {
		return 0, scanErr
	}
	got := strings.TrimSpace(scanner.Text())
	return strconv.ParseInt(got, 10, 64)
}

func main() {
	usePrompt := true
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		usePrompt = false
	}
	scanner := bufio.NewScanner(os.Stdin)

	inputter := func() (int64, error) {
		return getInput(usePrompt, scanner)
	}

	if len(os.Args) != 2 {
		errOut("usage: intcode <programfile>")
		os.Exit(2)
	}
	programFile := os.Args[1]
	strings.TrimSpace(string(lib.MustReadFile(programFile)))
	outputHandler := func(_ *intcomputer.IntComputer, n int64) error {
		fmt.Println(n)
		return nil
	}
	program := lib.CSInts(strings.TrimSpace(string(lib.MustReadFile(programFile))))
	c := intcomputer.NewIntComputer(program, outputHandler, inputter)
	err = c.RunOperations()
	if err != nil {
		panic(err)
	}
}
