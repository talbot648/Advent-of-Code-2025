package dayOne

import (
	"fmt"
	"strconv"
)

func (pt *PasswordEntry) AddPTwo(numToAdd int) {
	pt.startingNumber += numToAdd
	for pt.startingNumber >= 100 {
		pt.startingNumber -= 100
		pt.password++
	}
	fmt.Println("Number added, new total:", pt.startingNumber)
}

func (pt *PasswordEntry) SubtractPTwo(numToSubtract int) {
	pt.startingNumber -= numToSubtract
	for pt.startingNumber < 0 {
		pt.startingNumber += 100
		pt.password++
	}
	fmt.Println("Number subtracted, new total:", pt.startingNumber)
}

func SolutionPTwo() int {
	passwordEntry := &PasswordEntry{startingNumber: 50}
	rotations := getRotationsPTwo()
	password := executeRotationsPTwo(rotations, passwordEntry)
	return password
}

func executeRotationsPTwo(rotations []string, passwordEntry *PasswordEntry) int {
	for _, rotation := range rotations {
		switch rotation[0] {
		case 82: //Right
			numToAdd, _ := strconv.Atoi(rotation[1:])
			passwordEntry.AddPTwo(numToAdd)
		case 76: //Left
			numToSubtract, _ := strconv.Atoi(rotation[1:])
			passwordEntry.SubtractPTwo(numToSubtract)
		default:
			fmt.Println("Invalid rotation:", rotation)
		}
	}
	return passwordEntry.password
}

func getRotationsPTwo() []string {
	rotations, err := readInput("dayOne/input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return nil
	}
	return rotations
}
