package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "12a"

	i, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(i)
}
