package util

import (
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
