package main

import (
	"fmt"

	"./roman"
)

func main() {
	num := 1024
	result := roman.ToRoman(num)
	fmt.Println(num, "To Roman ==", result)

}
