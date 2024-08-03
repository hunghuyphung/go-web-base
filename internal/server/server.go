package server

import (
	"fmt"
	"go-gin-web/internal/controller"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type AppServer struct {
	AppController  controller.AppController
	UserController controller.UserController
	RoleController controller.RoleController
}

func NewAppServer(
	appController controller.AppController,
	userController controller.UserController,
	roleController controller.RoleController,
) *AppServer {
	return &AppServer{
		AppController:  appController,
		UserController: userController,
		RoleController: roleController,
	}
}

func NewServer(appServer *AppServer) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := appServer

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
