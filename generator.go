package gots_sdk

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

type SdkGenerator struct {
	ClientScript string
	Model        *typescriptify.TypeScriptify
}

var classTemplate = `
export class ClientSdk {
    host!: string
    client!: AxiosInstance

// ##endfunction##
}
`

var tsFuncTemplate = `
	async ##name##(query: ##query####payloadparam##): Promise<##response##> {
		let res = await this.client.##method##<any, AxiosResponse<##response##, any>, ##payload##>('##url##',##payloadinsert## {
			params: query
		});
		return res.data;
	}
`

func tsFuncName(method string, u string) string {
	funcname := strcase.ToCamel(strcase.ToSnake(method))

	urls := strings.Split(u, "/")

	for _, u := range urls {
		funcname += strcase.ToCamel(u)
	}
	return funcname
}

func (gen *SdkGenerator) CreateTsFunc(u string, method string, query interface{}, body interface{}, res interface{}) {
	gen.Model.Add(query).
		Add(res)

	funcname := tsFuncName(method, u)
	resname := reflect.TypeOf(res).Name()
	queryname := reflect.TypeOf(query).Name()

	funcstr := strings.Replace(tsFuncTemplate, `##name##`, funcname, 1)
	funcstr = strings.ReplaceAll(funcstr, `##response##`, resname)
	funcstr = strings.ReplaceAll(funcstr, `##url##`, u)

	funcstr = strings.ReplaceAll(funcstr, `##query##`, queryname)
	funcstr = strings.ReplaceAll(funcstr, `##method##`, strings.ToLower(method))

	if body == nil {
		funcstr = strings.ReplaceAll(funcstr, `##payload##`, `any`)
		funcstr = strings.ReplaceAll(funcstr, `##payloadparam##`, "")
		funcstr = strings.ReplaceAll(funcstr, `##payloadinsert##`, "")
	} else {
		payloadname := reflect.TypeOf(body).Name()
		funcstr = strings.ReplaceAll(funcstr, `##payload##`, payloadname)
		funcstr = strings.ReplaceAll(funcstr, `##payloadparam##`, `, body: `+payloadname)
		funcstr = strings.ReplaceAll(funcstr, `##payloadinsert##`, " body,")

		gen.Model.Add(body)
	}

	full := funcstr + "\n\n// ##endfunction##"
	gen.ClientScript = strings.Replace(gen.ClientScript, `// ##endfunction##`, full, 1)

}

func (gen *SdkGenerator) Generate(fname string) {
	cwdpath, _ := os.Getwd()
	basepath := filepath.Join(cwdpath, fname)

	os.Remove(basepath)

	f, err := os.OpenFile(basepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	model, _ := gen.Model.Convert(map[string]string{})
	scriptImport := `import { AxiosInstance, AxiosResponse } from 'axios'` + "\n\n"
	f.Write([]byte(scriptImport))
	f.Write([]byte(model))
	f.Write([]byte("\n\n"))
	f.Write([]byte(gen.ClientScript))
}

func NewSdkGenerator() *SdkGenerator {

	gen := SdkGenerator{
		ClientScript: classTemplate,
		Model:        typescriptify.New(),
	}

	return &gen
}
