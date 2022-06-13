package main

import "fmt"

func Run() error {
	fmt.Println("Starting our Application")
	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
