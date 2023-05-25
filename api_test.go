package gots_sdk_test

import (
	"net/http"
	"testing"

	"github.com/Vaziria/gots_sdk"
)

type PayloadData struct {
	Name string
}

func TestGenerateTs(t *testing.T) {
	api := gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "/user/data/create",
		Payload:      &PayloadData{},
	}

	tsfunc := api.GenerateTs()
	t.Log(tsfunc)
}
