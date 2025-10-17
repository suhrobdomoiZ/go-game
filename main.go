package main

import (
	"bufio"
	"fmt"
	. "go-game/functions"
	"os"
)

func main() {
	InitGame()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()
		if command == "выход" {
			break
		}
		result := HandleCommand(command)
		fmt.Println(result)
	}

}
