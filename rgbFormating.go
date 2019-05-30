package main

import (
	"fmt"
	"math"
	"strconv"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func formatRGB(charArray []rune) {
	for j := 0; j < len(charArray); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, charArray[j])
	}
	fmt.Printf(" ")
}

func print(userInitial string, output string, intNumber int) {
	// string response and no integer response
	charuserInitialArray := []rune(userInitial)
	formatRGB(charuserInitialArray)

	if output != "" && intNumber == -1 {
		// convert string to charArray
		charOutputArray := []rune(output)

		formatRGB(charOutputArray)
	} else {
		intNumberToChar := strconv.Itoa(intNumber)
		charNumber := []rune(intNumberToChar)
		formatRGB(charNumber)
	}

	fmt.Println()
}
