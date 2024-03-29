package main

import (
	"context"
	"fmt"
	"go-rest-api/internal/comment"
	"go-rest-api/internal/db"
	transportHttp "go-rest-api/internal/transport/http"
)

func Run() error {

	fmt.Println("staring up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	cmtSerive := comment.NewServer(db)

	httpHandler := transportHttp.NewHandler(cmtSerive)
	if err := httpHandler.Serve(); err != nil {
		return err

	}
	return nil
}

func main() {
	fmt.Println("GO Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
