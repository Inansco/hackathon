package main

import (
	"strings"
)

/*
Extend your file reader: now accept TWO command-line arguments 
(input file, output file). Read the input file, 
convert all text to uppercase using strings.ToUpper, and 
write the result to the output file using os.WriteFile.
Make sure to use the correct file permission flags.
*/

func ToUpper(s string) string {
	files := strings.ToUpper(s)
	return files
}