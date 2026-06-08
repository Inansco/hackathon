package main

import (
	"fmt"
	"os"

)

/*Extend your file reader:
now accept TWO command-line arguments (input file, output file).
Read the input file, convert all text to uppercase using strings.
ToUpper, and write the result to the output file using os.WriteFile.
Make sure to use the correct file permission flags.
*/

func man() {
	if len(os.Args) < 2 {
		fmt.Println("use: go run . sandard.txt result.txt")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error")
		return 
	}

	UpData := Allfuncs(string(data))

	err = os.WriteFile(output,[]byte(UpData), 0644)
	if err != nil{
		return 
	
	}
	fmt.Println("done")
}