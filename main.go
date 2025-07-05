package main

import (
	"fmt"
	"context"
	"github.com/DanYT2/orders-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start server", err)
	}
}
