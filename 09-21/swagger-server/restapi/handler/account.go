package handler
import (
	"swaggertest/models"
	"swaggertest/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)
func ConfigGetAccount(params operations.GetAccountAllParams) middleware.Responder {
	var result []*models.AccountEntry
	result = make([]*models.AccountEntry, 0)

   // 계정 정보 추가
	result = append(result, &models.AccountEntry{
		UserID:   "1",
		Password: "password",
		Email:    "test@test.io",
	})

	return operations.NewGetAccountAllOK().WithPayload(&operations.GetAccountAllOKBody{Attr: result})
}