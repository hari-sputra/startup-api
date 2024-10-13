package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"startup-api/API/user"
	"startup-api/auth"
	"startup-api/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	userService user.UserService
	authService auth.AuthService
}

func NewUserHandler(userService user.UserService, authService auth.AuthService) *userHandler {
	return &userHandler{userService, authService}
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

	token, err := h.authService.GenerateJWTToken(createUser.ID)
	if err != nil {
		res := helper.APIResponse("Create token failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := user.FormatterData(createUser, token)
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

	token, err := h.authService.GenerateJWTToken(loginUser.ID)
	if err != nil {
		res := helper.APIResponse("Create token failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := user.FormatterData(loginUser, token)
	res := helper.APIResponse("User logged successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, res)
}

func (h *userHandler) CheckAvailableEmail(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationFormatter(err)
		errMsg := gin.H{"errors": errors}

		res := helper.APIResponse("email checking failed", http.StatusUnprocessableEntity, "error", errMsg)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	isEmail, errEmail := h.userService.IsEmailAvailable(input)
	if errEmail != nil {
		errMsg := gin.H{"errors": "Server error"}
		res := helper.APIResponse("User login failed", http.StatusBadRequest, "error", errMsg)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	data := gin.H{
		"is_available": isEmail,
	}

	metaMsg := "Email has been registered"

	if isEmail {
		metaMsg = "Email is available"
	}

	res := helper.APIResponse(metaMsg, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, res)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		res := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	uniqueFileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))
	path := "storage/images/" + uniqueFileName

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		res := helper.APIResponse("Failed to save avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	currentUser := c.MustGet("Current-User").(user.User)
	userID := currentUser.ID

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		res := helper.APIResponse("Failed to save avatar path", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	data := gin.H{"is_uploaded": true}
	res := helper.APIResponse("Avatar uploaded successfully", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, res)
}
