package dayOne

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PasswordEntry struct {
	password       int
	startingNumber int
}

func (p *PasswordEntry) Add(numToAdd int) {
	p.startingNumber += numToAdd
	if p.startingNumber >= 100 {
		p.startingNumber -= 100
	}
}

func (p *PasswordEntry) Subtract(numToSubtract int) {
	p.startingNumber -= numToSubtract
	if p.startingNumber < 0 {
		p.startingNumber += 100
	}
}

func isCurrentRotationZero(passwordEntry PasswordEntry) bool {
	return passwordEntry.startingNumber == 0
}
func Solution() int {
	var passwordEntry PasswordEntry
	passwordEntry.startingNumber = 50
	rotations := getRotations()
	password := executeRotations(rotations, passwordEntry)
	return password
}

func executeRotations(rotations []string, passwordEntry PasswordEntry) int {
	for _, rotation := range rotations {
		switch rotation[0] {
		case 82: //Right
			numToAdd, _ := strconv.Atoi(rotation[1:])
			if numToAdd > 100 {
				numToAdd = numToAdd % 100
			}
			passwordEntry.Add(numToAdd)
			if isCurrentRotationZero(passwordEntry) {
				passwordEntry.password++
			}
		case 76: //Left
			numToSubtract, _ := strconv.Atoi(rotation[1:])
			if numToSubtract > 100 {
				numToSubtract = numToSubtract % 100
			}
			passwordEntry.Subtract(numToSubtract)
			if isCurrentRotationZero(passwordEntry) {
				passwordEntry.password++
			}
		default:
			fmt.Println("Invalid rotation:", rotation)
		}
	}
	return passwordEntry.password
}

func getRotations() []string {
	rotations, err := readInput("dayOne/input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return nil
	}
	return rotations
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil

}
