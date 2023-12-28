package main

import (
	"log"
	"os"
	"strconv"

	"github.com/akers1023/handlers"
	"github.com/akers1023/models"

	"github.com/akers1023/routes"
	"github.com/akers1023/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	sql := &storage.Sql{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Password: os.Getenv("DB_PASS"),
		UserName: os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("DB_NAME"),
	}

	db, _ := sql.Connect()
	defer sql.Close()

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	repo := &handlers.Repository{
		DB: db,
	}

	// print(db)
	app := fiber.New()
	routes.SetupBooksRoutes(app, repo)
	routes.SetupUsersRoutes(app, repo)
	app.Listen(":3000")

}
