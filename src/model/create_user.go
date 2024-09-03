package model

import (
	"fmt"
	"github.com/MateusGuedess/crudzinho-go/src/configuration/logger"
	"github.com/MateusGuedess/crudzinho-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (userDomain *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain)
	return nil
}
