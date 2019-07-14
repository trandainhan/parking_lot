package utils

import (
	"bufio"
	"os"
)

func LoadCmd(txtFile string) []string {
	var parkingCmd []string
	file, err := os.Open(txtFile)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parkingCmd = append(parkingCmd, line)
	}
	return parkingCmd
}
