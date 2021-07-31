package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(" welcome to simon says! :) ")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	chkstring := "simon says"

	for count = 1; count <= 5; count++ {
		fmt.Print("Enter the command you  want me to do : ")
		reader := bufio.NewReader(os.Stdin)
		input, error := reader.ReadString('\n')
		getinput := strings.ToLower(input)

		if error != nil {
			fmt.Println("There is an issue reading the input", error)
			return
		}

		if strings.Contains(getinput, chkstring) {
			command := strings.SplitN(getinput, "simon says", -1)
			fmt.Println(command)
			fmt.Println("the count is : ", count)
		} else {
			fmt.Println("Don't see simon says, so wont do TRY AGAIN........... ")
			count -= 1

		}

	}
	fmt.Println("CONGRATULATIONS!!!!!!  YOU  HAVE WON THE GAME........")
	fmt.Println("***************************************************")
}

//https://git-scm.com/downloads
