package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/giannisalinetti/photo-lab/tools/ev-convert/pkg/convert"
)

// ScanValue get the actual scan value from stdin
func ScanValue() (float64, string, error) {
	var input string
	var value float64
	var scanType string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose convertion")
		fmt.Printf("1) Lux to EV100\n2) EV100 to Lux\n")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "1" {
			scanType = "Lux"
		} else if choice == "2" {
			scanType = "EV100"
		} else {
			continue
		}

		fmt.Printf("Enter %s value:\n", scanType)
		scanner.Scan()
		input = scanner.Text()
		if len(input) != 0 {
			break
		} else {
			continue
		}
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, "", err
	}
	return value, scanType, nil
}

func main() {

	var res float64

	inputValue, inputType, err := ScanValue()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if inputType == "Lux" {
		res = convert.LuxToEv100(inputValue)
	} else {
		res = convert.Ev100ToLux(inputValue)
	}

	fmt.Println(res)
}
