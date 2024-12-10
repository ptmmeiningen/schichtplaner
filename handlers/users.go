package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ptmmeiningen/schichtplaner/database"
	"github.com/ptmmeiningen/schichtplaner/models"
)

// @Summary Get all users
// @Description fetch all users
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /users [get]
func HandleAllUsers(c *fiber.Ctx) error {
	var users []models.User
	result := database.GetDB().Preload("Departments").Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
	}
	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Users successfully retrieved",
		Data:    users,
	})
}

type CreateUserDTO struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Color         string `json:"color"`
	IsAdmin       bool   `json:"is_admin"`
	DepartmentIDs []uint `json:"department_ids"`
}

// @Summary Create a user
// @Description create new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserDTO true "User to create"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /users [post]
func HandleCreateUser(c *fiber.Ctx) error {
	dto := new(CreateUserDTO)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Error:   "Invalid input",
		})
	}

	user := models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
		Color:     dto.Color,
		IsAdmin:   dto.IsAdmin,
	}

	if len(dto.DepartmentIDs) > 0 {
		var departments []models.Department
		if err := database.GetDB().Find(&departments, dto.DepartmentIDs).Error; err != nil {
			return c.Status(400).JSON(models.APIResponse{
				Success: false,
				Error:   "Invalid department IDs",
			})
		}
		user.Departments = departments
	}

	result := database.GetDB().Create(&user)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "User successfully created",
		Data:    user,
	})
}

// @Summary Get a single user
// @Description fetch user by ID
// @Tags users
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /users/{id} [get]
func HandleGetOneUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := database.GetDB().Preload("Departments").Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Error:   "User not found",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "User successfully retrieved",
		Data:    user,
	})
}

// @Summary Update a user
// @Description update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body CreateUserDTO true "User update data"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /users/{id} [put]
func HandleUpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := database.GetDB().Preload("Departments").Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Error:   "User not found",
		})
	}

	dto := new(CreateUserDTO)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Error:   "Invalid input",
		})
	}

	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Email = dto.Email
	user.Password = dto.Password
	user.Color = dto.Color
	user.IsAdmin = dto.IsAdmin

	if len(dto.DepartmentIDs) > 0 {
		var departments []models.Department
		if err := database.GetDB().Find(&departments, dto.DepartmentIDs).Error; err != nil {
			return c.Status(400).JSON(models.APIResponse{
				Success: false,
				Error:   "Invalid department IDs",
			})
		}
		user.Departments = departments
	}

	database.GetDB().Save(&user)
	return c.JSON(models.APIResponse{
		Success: true,
		Message: "User successfully updated",
		Data:    user,
	})
}

// @Summary Delete a user
// @Description delete user by ID
// @Tags users
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /users/{id} [delete]
func HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.GetDB().Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "User successfully deleted",
	})
}
