package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"parking_lot/internal/parking"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 { // Run program from file
		inputFile := args[0]
		commands := LoadCmd(inputFile)
		result := parking.ProcessMultipleCommand(commands)
		fmt.Println(result)
	} else { // Run program as interactive shell
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			cmd := scanner.Text()
			if cmd == "exit" {
				break
			}
			if len(cmd) == 0 {
				continue
			}
			response := parking.ProcessCommand(cmd)
			fmt.Println(response)
		}
	}
}
