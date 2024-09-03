package controller

import (
	"fmt"
	"github.com/MateusGuedess/crudzinho-go/src/configuration/validation"
	"github.com/MateusGuedess/crudzinho-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		context.JSON(http.StatusBadRequest, restErr)
		return
	}

	fmt.Println(userRequest)
}
