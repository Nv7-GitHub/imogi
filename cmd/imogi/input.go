package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func GetInput(input string) string {
	fmt.Print(input)
	text, _, err := reader.ReadLine()
	handle(err)
	return string(text)
}

func GetInputInt(input string) int {
	text := GetInput(input)
	val, err := strconv.Atoi(text)
	handle(err)
	return val
}
