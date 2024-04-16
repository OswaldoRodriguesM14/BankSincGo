package pkg

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateBody(r *http.Request, schema_input_json string) (error, []byte) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("falha ao ler o corpo da requisição"), nil
	}

	bodyString := string(body)
	schemaLoader := gojsonschema.NewStringLoader(schema_input_json)
	documentLoader := gojsonschema.NewStringLoader(bodyString)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return errors.New("erro ao efetuar a validação - Contate o Suporte"), nil
	}
	if result.Valid() {
		return nil, body
	} else {
		var errorStrings []string
		for _, desc := range result.Errors() {
			errorStrings = append(errorStrings, desc.String())
		}
		// Concatenando todos os erros em uma única string, separados por ponto e vírgula
		return errors.New(strings.Join(errorStrings, "; ")), nil
	}

}
