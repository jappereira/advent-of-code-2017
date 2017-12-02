package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//func main() {
//	fmt.Println(compute2())
//}

func compute2() int {
	input, err := ioutil.ReadFile("2.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Print(string(input))
	return 1
}
