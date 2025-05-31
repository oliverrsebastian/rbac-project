package main

import (
	"context"
	"fmt"
	"rbac-project/container"
	"runtime/debug"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(ctx, "Server panicked: ", string(debug.Stack()))
		}
	}()

	server, err := container.InitServer(ctx)
	if err != nil {
		panic(err)
	}

	if err := server.Start(":9595"); err != nil {
		panic(err)
	}

	if err := server.Close(); err != nil {
		panic(err)
	}
}
