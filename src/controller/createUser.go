package controller

import (
	"fmt"
	"github.com/MateusGuedess/crudzinho-go/src/configuration/rest_err"
	"github.com/MateusGuedess/crudzinho-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequest(fmt.Sprintf("There are some incorrect fields,  error=%s\n", err.Error()))
		context.JSON(http.StatusBadRequest, restErr)
		return
	}

	fmt.Println(userRequest)
}
