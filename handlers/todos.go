package handlers

import (
	"github.com/ptmmeiningen/schichtplaner/database"
	"github.com/ptmmeiningen/schichtplaner/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get all todos.
// @Description fetch every todo available.
// @Tags todos
// @Accept */*
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /todos [get]
func HandleAllTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	result := database.GetDB().Find(&todos)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
	}
	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Todos successfully retrieved",
		Data:    todos,
	})
}

type CreateTodoDTO struct {
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

// @Summary Create a todo.
// @Description create a single todo.
// @Tags todos
// @Accept json
// @Param todo body CreateTodoDTO true "Todo to create"
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /todos [post]
func HandleCreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Error:   "Invalid input",
		})
	}

	result := database.GetDB().Create(&todo)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Todo successfully created",
		Data:    todo,
	})
}

// @Summary Update a todo.
// @Description update a single todo.
// @Tags todos
// @Accept json
// @Param todo body CreateTodoDTO true "Todo update data"
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /todos/{id} [put]
func HandleUpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	var todo models.Todo
	if err := database.GetDB().First(&todo, id).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Error:   "Todo not found",
		})
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Error:   "Invalid input",
		})
	}

	database.GetDB().Save(&todo)
	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Todo successfully updated",
		Data:    todo,
	})
}

// @Summary Get a single todo.
// @Description fetch a single todo.
// @Tags todos
// @Param id path int true "Todo ID"
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /todos/{id} [get]
func HandleGetOneTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	var todo models.Todo
	if err := database.GetDB().First(&todo, id).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Error:   "Todo not found",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Todo successfully retrieved",
		Data:    todo,
	})
}

// @Summary Delete a single todo.
// @Description delete a single todo by id.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /todos/{id} [delete]
func HandleDeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.GetDB().Delete(&models.Todo{}, id)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Todo successfully deleted",
	})
}
