package dayTwo

import (
	"bufio"
	"fmt"
	"os"
)

func Solution() int {
	input, err := readInput("dayTwo/inputDayTwo.txt")
	if err != nil {
		fmt.Println("error reading file", err)
		return 0
	}
	fmt.Println(input)
	return 0
}

func readInput(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		// No data in the file
		return "", fmt.Errorf("file is empty")
	}

	return scanner.Text(), nil
}
