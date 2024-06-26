package main

import (
	"context"
	"fmt"

	"github.com/Infael/gogoVseProject/application"
)

func main() {
	app := application.New()

	err := app.Start(context.Background())

	if err != nil {
		fmt.Printf("failed to start application: %v\n", err)
	}
}
