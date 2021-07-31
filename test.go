package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var str = "14"
	// var input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("\ntype of value %T", input)
	// convert string to integer
	input = strings.TrimSpace(input)
	myVar, error := strconv.Atoi(input)
	fmt.Printf("\ntype of value %T", myVar)
	fmt.Println("\nerror:::::::::::", error)
	fmt.Println("\nbas 2 value %b", myVar)
	fmt.Println("number:", myVar)
}
