package main

import (
	"fmt"
	"strings"
)

func space(s string) []string {
	data := strings.Fields(s)
	return data
}

func split(s string) []string {
	
}



func main() {
	fmt.Println(space("Hello, world! How are you?"))
	fmt.Println(split("Hello , world ! How are you ?"))
}