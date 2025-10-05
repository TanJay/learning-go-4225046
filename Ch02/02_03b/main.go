package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	// str, _ := reader.ReadString('\n')
	str, err := reader.ReadString(',')
	if err == nil {
		fmt.Println(len(str))
	}
}
