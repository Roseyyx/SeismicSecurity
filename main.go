package main

import (
	"fmt"
	"log"
	"main/backend/helpers"
	"main/backend/routes"
	utilities "main/backend/utils"
	"main/frontend/app/functions"

	"github.com/gofiber/fiber/v2"
)

func main() {

	utilities.Client = utilities.ConnectDB()

	if !helpers.Debug {
		log.Println("Using File")
		functions.Setup()
	} else {
		app := fiber.New(fiber.Config{
			EnablePrintRoutes: true,
		})

		// Routes
		routes.RoutesHandler(app)

		// Get Port from .ENV file
		port := utilities.GetEnvVariable("PORT")
		// If .ENV file is not defined or PORT does not exist refer to default port
		if port == "" {
			port = "4545"
		}

		log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
	}
}
