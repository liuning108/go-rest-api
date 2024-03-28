package main

import (
	"fmt"
)

func Run() error {

	fmt.Println("staring up our application")
	return nil
}

func main() {
	fmt.Println("GO Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
