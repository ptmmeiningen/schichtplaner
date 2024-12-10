package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ptmmeiningen/schichtplaner/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	// setup the todos group
	todos := app.Group("/todos")
	todos.Get("/", handlers.HandleAllTodos)
	todos.Post("/", handlers.HandleCreateTodo)
	todos.Put("/:id", handlers.HandleUpdateTodo)
	todos.Get("/:id", handlers.HandleGetOneTodo)
	todos.Delete("/:id", handlers.HandleDeleteTodo)

	// setup the users group
	users := app.Group("/users")
	users.Get("/", handlers.HandleAllUsers)
	users.Post("/", handlers.HandleCreateUser)
	users.Get("/:id", handlers.HandleGetOneUser)
	users.Put("/:id", handlers.HandleUpdateUser)
	users.Delete("/:id", handlers.HandleDeleteUser)

	// setup the departments group
	departments := app.Group(("/departments"))
	departments.Get("/", handlers.HandleAllDepartments)
	departments.Post("/", handlers.HandleCreateDepartment)
	departments.Get("/:id", handlers.HandleGetOneDepartment)
	departments.Put("/:id", handlers.HandleUpdateDepartment)
	departments.Delete("/:id", handlers.HandleDeleteDepartment)
}
