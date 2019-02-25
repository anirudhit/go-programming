package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Printing a word
	var st string
	fmt.Println("Enter word:")
	fmt.Scanln(&st)
	fmt.Println(st)

	//Printing a message
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter message:")
	str, _ := reader.ReadString('\n')
	fmt.Printf(str)

	//Parsing a string to float 64
	fmt.Println("Enter number:")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", f)
	}
}
