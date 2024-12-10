package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ptmmeiningen/schichtplaner/database"
	"github.com/ptmmeiningen/schichtplaner/models"
)

// @Summary Get all departments
// @Description fetch all departments
// @Tags departments
// @Accept */*
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /departments [get]
func HandleAllDepartments(c *fiber.Ctx) error {
	var departments []models.Department
	result := database.GetDB().Preload("Users").Find(&departments)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
	}
	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Departments successfully retrieved",
		Data:    departments,
	})
}

type CreateDepartmentDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

// @Summary Create a department
// @Description create new department
// @Tags departments
// @Accept json
// @Produce json
// @Param department body CreateDepartmentDTO true "Department to create"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /departments [post]
func HandleCreateDepartment(c *fiber.Ctx) error {
	department := new(models.Department)
	if err := c.BodyParser(department); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Error:   "Invalid input",
		})
	}

	result := database.GetDB().Create(&department)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Department successfully created",
		Data:    department,
	})
}

// @Summary Get a single department
// @Description fetch department by ID
// @Tags departments
// @Param id path int true "Department ID"
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /departments/{id} [get]
func HandleGetOneDepartment(c *fiber.Ctx) error {
	id := c.Params("id")

	var department models.Department
	if err := database.GetDB().Preload("Users").Where("id = ?", id).First(&department).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Error:   "Department not found",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Department successfully retrieved",
		Data:    department,
	})
}

// @Summary Update a department
// @Description update department by ID
// @Tags departments
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Param department body CreateDepartmentDTO true "Department update data"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /departments/{id} [put]
func HandleUpdateDepartment(c *fiber.Ctx) error {
	id := c.Params("id")

	var department models.Department
	if err := database.GetDB().Where("id = ?", id).First(&department).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Error:   "Department not found",
		})
	}

	if err := c.BodyParser(&department); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Error:   "Invalid input",
		})
	}

	database.GetDB().Save(&department)
	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Department successfully updated",
		Data:    department,
	})
}

// @Summary Delete a department
// @Description delete department by ID
// @Tags departments
// @Param id path int true "Department ID"
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /departments/{id} [delete]
func HandleDeleteDepartment(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.GetDB().Where("id = ?", id).Delete(&models.Department{})
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Department successfully deleted",
	})
}
