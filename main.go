package main

import (
	dayOne "AdventOfCode/dayOne"
	"AdventOfCode/dayTwo"
	"fmt"
)

func main() {
	dayOneAnswer := dayOne.Solution()
	fmt.Println(dayOneAnswer)
	dayOnePartTwoAnswer := dayOne.SolutionPTwo()
	fmt.Println(dayOnePartTwoAnswer)
	dayTwoAnswer := dayTwo.Solution()
	fmt.Println(dayTwoAnswer)
}
