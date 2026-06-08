package main

import (
	"fmt"
	"strconv"
)

/*
Write a function hexToDec(s string) (int64, error)
that converts a hexadecimal string to its decimal integer
equivalent. Test it with inputs like '1E' (should return 30),
'FF' (should return 255), and 'BADF00D'.
Use only the standard library.
*/

func hexToDec(s string) (int64, error) {
	hex, err := strconv.ParseInt(s, 16, 64)
	return hex, err
}

func mai() {
	fmt.Println(hexToDec("1E"))
	fmt.Println(hexToDec("FF"))
	fmt.Println(hexToDec("BADFOOD"))

}
