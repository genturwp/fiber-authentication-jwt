package main

import (
	"context"
	"fiberauthenticationjwt/api"
	"fiberauthenticationjwt/repositories"
	"fiberauthenticationjwt/services"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Cannot find configuration file. Make sure .env file exist in the same directory with the binary file")
	}
}

func dbConn() *pgxpool.Pool {
	connURL, ok := viper.Get("DB_CONN").(string)
	if !ok {
		log.Fatalln("cannot read DB_CONN in config file")
	}

	config, err := pgxpool.ParseConfig(connURL)
	if err != nil {
		log.Fatalln("cannot get db configuration")
	}

	config.MinConns = 4
	config.MaxConns = 5

	connPool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalln("cannot connect to database")
	}

	return connPool
}

func main() {
	fmt.Println("Fiber authentication jwt")
	poolConn := dbConn()
	repository := repositories.NewRepository(poolConn)
	service := services.NewService(repository)

	app := fiber.New()
	app.Use(cors.New(), logger.New(), recover.New())
	api.Handler(service, app)

	app.Listen(":9797")
}
