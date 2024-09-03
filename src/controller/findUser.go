package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindUserByID(context *gin.Context) {

}

func FindUserByEmail(context *gin.Context) {
	email := context.Query("email")
	context.String(http.StatusOK, "Hello %s %s", email)
	fmt.Println(email)
}
