package main

import (
	"bufio"
	"fmt"
	"kata/calculator"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Input:\n")
	for scan.Scan() {
		if len(scan.Text()) > 0 {
			val, err := calculator.Calculate(scan.Text())
			if err != nil {
				panic(err)
			}
			fmt.Printf("Output:\n %s", val)
			// break
		}

	}

}
