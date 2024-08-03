package main

import (
	"fmt"
	"go-gin-web/internal/injector"
	"go-gin-web/internal/server"
)

func main() {
	app, initializationErr := injector.InitializeApp()

	if initializationErr != nil {
		panic(fmt.Sprintf("Injecting dependencies failed %s", initializationErr))
	}

	appServer := server.NewServer(app.AppServer)

	err := appServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
