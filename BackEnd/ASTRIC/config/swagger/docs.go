package swagger

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/swaggo/swag"
)

type s struct{}

type swaggerInfo struct {
	Host  string
	Token string
}

func (s *s) ReadDoc() string {

	var SwaggerInfo swaggerInfo

	SwaggerInfo.Host = env.GetEnvGeneral().ServerIP + ":" + env.GetEnvDoc().Port

	bytesLeidos, err := os.ReadFile("swagger.json")
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}

	doc := string(bytesLeidos)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
