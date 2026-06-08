package main

import (
	"fmt"
	"strconv"
)

func binToDec(s string) (int64, error) {
	hex, err := strconv.ParseInt(s, 2, 64)
	return hex, err
}

func mn() {
	fmt.Println(binToDec("10"))
	fmt.Println(binToDec("1010"))
	fmt.Println(binToDec("11111111"))

}