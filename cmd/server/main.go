package main

import (
	"context"
	"fmt"
	db2 "goRestApi/internal/db"
)

func Run() error {
	fmt.Println("Starting our Application")
	db, err := db2.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to Database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
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
