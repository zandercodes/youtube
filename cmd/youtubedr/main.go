package main

import (
	"fmt"
	"os"
)

func main() {
	Execute()
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
