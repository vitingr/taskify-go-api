package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Print("SERVIDOR FUNCIONANDO...")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []Todo{}

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		fmt.Print(todo)

		todo.ID = len(todos) + 1
		todo.Done = false

		fmt.Print(todo)

		todos = append(todos, *todo)

		fmt.Print(todos)

		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c * fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("ID Inválido")
		}

		// I = Index || T = Task to do
		for i, t := range todos {
			if t.ID == id {

				if todos[i].Done == true {
					todos[i].Done = false
					break
				} 
				if todos[i].Done == false {
					todos[i].Done = true
					break
				}

			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c * fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Listen(":4000")

}