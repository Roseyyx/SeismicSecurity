package main

import (
	"fmt"
	"log"
	"main/routes"
	utilities "main/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	utilities.Client = utilities.ConnectDB()

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

	// Auto error checking log and listen to previous specified port
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
