package controllers

import (
	"net/http"
	"startup/app"
	"startup/app/helpers"
	"startup/app/services"
	"startup/app/structs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type roleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *roleController {
	return &roleController{roleService}
}

func (h *roleController) Index(c *gin.Context) {
	roles, err := h.roleService.GetRoles()
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.RoleSummaryResponse(roles)
	res := helpers.ResponseAPI("List of roles.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *roleController) Store(c *gin.Context) {
	var request structs.RoleStoreRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ResponseAPI("Something wrong with the request.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = app.RegisterRoleNameAvailableValidation(validate, h.roleService)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		errors := helpers.FormatMessageValidationErrors(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Failed to create a role.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	_, err = h.roleService.CreateRole(request)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseAPI("Role successfully created.", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, res)
}