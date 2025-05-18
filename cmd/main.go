package main

import (
	"os"
	"pos-kasir/config"
	httpHandler "pos-kasir/internal/delivery/http"
	"pos-kasir/internal/repository"
	"pos-kasir/internal/usecase"
	"pos-kasir/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectPostgres()
	r := gin.Default()

	hasher := utils.NewHasher()
	jwtSecret := os.Getenv("JWT_SECRET")
	jwt := utils.NewJWT(jwtSecret)

	repo := repository.NewUserRepo(db)
	usecase := usecase.NewUserUsecase(repo, hasher, jwt)
	httpHandler.NewUserHandler(r, usecase)

	r.Run(":8080")
}
