package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()
	todos := []Todo{}

	// Get a Todo
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Create a Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := new(Todo) // Gunakan pointer ke struct

		// Parse body JSON ke struct Todo
		if err := c.BodyParser(todo); err != nil {
			return c.Status(422).JSON(fiber.Map{"error": "Invalid input", "details": err.Error()})
		}

		// Validasi field yang dibutuhkan
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		// Set ID secara dinamis
		todo.ID = len(todos) + 1
		todos = append(todos, *todo) // Tambahkan ke list `todos`

		// Kembalikan todo yang baru dibuat
		return c.Status(201).JSON(todo)
	})

	//Update a Todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found!"})
	})

	//Delete a Todo
	app.Delete("api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found!"})
	})

	log.Fatal(app.Listen(":4000"))
}
