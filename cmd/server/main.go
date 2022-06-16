package main

import (
	"fmt"
	"goRestApi/internal/db"
)

func Run() error {
	fmt.Println("Starting our Application")
	_, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the Database")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
