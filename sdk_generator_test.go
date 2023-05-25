package gots_sdk_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/gots_sdk"
)

type PayloadDataDD struct {
	Name string
}

func TestCreateSDK(t *testing.T) {
	sdk := gots_sdk.NewApiSdk(gin.Default())

	save := sdk.GenerateSdkFunc("sdk.ts")

	sdk.Register(&gots_sdk.Api{
		Payload:      PayloadDataDD{},
		Method:       http.MethodPost,
		RelativePath: "/users",
	}, func(ctx *gin.Context) {

	})

	sdk.RegisterGroup("/product", func(group *gin.RouterGroup, register gots_sdk.RegisterFunc) {
		register(&gots_sdk.Api{
			Payload:      PayloadDataDD{},
			Method:       http.MethodPost,
			RelativePath: "/create",
		})
	})

	save()
}
