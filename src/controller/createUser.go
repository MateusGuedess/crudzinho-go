package controller

import (
	"github.com/MateusGuedess/crudzinho-go/src/configuration/logger"
	"github.com/MateusGuedess/crudzinho-go/src/configuration/validation"
	"github.com/MateusGuedess/crudzinho-go/src/controller/model/request"
	"github.com/MateusGuedess/crudzinho-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CreateUser(context *gin.Context) {
	logger.Info("Init Create User controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		context.JSON(http.StatusBadRequest, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	context.JSON(http.StatusCreated, response)
}
