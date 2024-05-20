package routes

import (
	"main/backend/functions"
	utilities "main/backend/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RoutesHandler(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/entries", GetEntries)
	v1.Get("/entry/:username", GetEntryByName)
	v1.Post("/entry", CreateEntry)
	v1.Delete("/entry/:id", DeleteEntry)
	v1.Put("/entry/:id", UpdateEntry)
}

func GetEntries(c *fiber.Ctx) error {
	entries := functions.GetEntries()

	return c.JSON(entries)
}

func GetEntryByName(c *fiber.Ctx) error {
	username := c.Params("username")
	entry := functions.GetEntryByName(username)

	if entry.Username == "" {
		return c.Status(404).SendString("Entry not found")
	}

	return c.JSON(entry)
}

func CreateEntry(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	website := c.FormValue("website")
	notes := c.FormValue("notes")

	// Hash the password
	hashedPassword, err := utilities.HashPassword(password)
	if err != nil {
		return c.Status(500).SendString("Failed to hash the password")
	}

	functions.CreateEntry(username, hashedPassword, website, notes)

	return c.SendString("Entry created successfully")
}

func DeleteEntry(c *fiber.Ctx) error {
	id := c.Params("id")

	// covert id to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	functions.DeleteEntry(objectID)
	return c.SendString("Entry deleted successfully")
}

func UpdateEntry(c *fiber.Ctx) error {
	id := c.Params("id")
	username := c.FormValue("username")
	password := c.FormValue("password")
	website := c.FormValue("website")
	notes := c.FormValue("notes")

	// covert id to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	functions.UpdateEntry(objectID, username, password, website, notes)
	return c.SendString("Entry updated successfully")
}
