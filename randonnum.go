package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 50
	num := (rand.Intn(max-min+1) + min)
	fmt.Println(num)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(" welcome to simon says Guessing Number!!!!! ")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Guess the number between 0 to 50 .......")
	reader := bufio.NewReader(os.Stdin)
	for count := 1; count <= 10; count++ {
		fmt.Println("counter : = ", count)

		//reader := bufio.NewReader(os.Stdin)
		input, error := reader.ReadString('\n')

		input = strings.TrimSpace(input)
		intnum, err := strconv.Atoi(input)

		if error != nil {
			fmt.Println("There is issue in reading the input", error)
			return
		}
		if err != nil {
			fmt.Println("You have entered the string... Please enter the number from 1 to 50", err)

		}
		if intnum > num {
			fmt.Println("Guess the number lower ")
		} else if intnum < num {
			fmt.Println("Guess the number higher ")
		} else if intnum == num {
			fmt.Println("******Congratulations you won the game!!!!!!!!!******** ")
			return
		}

	}
	fmt.Println("You have reached the maximum number of tries....")
}
