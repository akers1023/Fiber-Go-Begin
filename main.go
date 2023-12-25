package main

import (
	db "fiber-begin/database"
	"fiber-begin/handler"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("DB_PORT must be a valid integer")
	}

	sql := &db.Sql{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Password: os.Getenv("DB_PASS"),
		UserName: os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("DB_NAME"),
	}

	db, _ := sql.Connect()
	repo := &handler.Repository{
		DB: db,
	}

	// utils.LogErrorf("Error: ")
	// log.WithFields(log.Fields{"animal": "walrus"}).Error("A walrus appears")
	app := fiber.New()
	repo.SetupBookRoutes(app)
	app.Listen(":3000")
	defer sql.Close()
}
