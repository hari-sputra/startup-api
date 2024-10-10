package handler

import (
	"net/http"
	"startup-api/helper"
	"startup-api/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.ErrorValidationFormatter(err)
		errMessage := gin.H{"errors": errors}

		res := helper.APIResponse("Register user failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	createUser, usrErr := h.userService.RegisterUser(input)

	if usrErr != nil {
		res := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := user.FormatterData(createUser, "token1223")
	res := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, res)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationFormatter(err)
		errMsg := gin.H{"errors": errors}

		res := helper.APIResponse("User login failed", http.StatusUnprocessableEntity, "error", errMsg)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	loginUser, usrerr := h.userService.LoginUser(input)
	if usrerr != nil {
		errMsg := gin.H{"errors": usrerr.Error()}
		res := helper.APIResponse("User login failed", http.StatusBadRequest, "error", errMsg)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := user.FormatterData(loginUser, "token1234")
	res := helper.APIResponse("User logged successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, res)
}
