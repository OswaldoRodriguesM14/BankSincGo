package pkg

import "net/http"

type UtilitsInterface interface {
	ValidateBody(r *http.Request, schema_input_json string) ([]byte, error)
	MarshalValidationError(err error) []byte
	PrimeiroUltimoDiaMesAtual() (startDate string, endDate string)
	GeraIdempotencyKey() string
}

type Utilits struct{}

func UtilitsCompose() *Utilits {
	return &Utilits{}
}

func (u *Utilits) ValidateBody(r *http.Request, schema_input_json string) ([]byte, error) {
	err, body := ValidateBody(r, schema_input_json)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (u *Utilits) MarshalValidationError(err error) []byte {
	return MarshalValidationError(err)
}

func (u *Utilits) PrimeiroUltimoDiaMesAtual() (startDate string, endDate string) {
	return PrimeiroUltimoDiaMesAtual()
}

func (u *Utilits) GeraIdempotencyKey() string {
	return GeraIdempotencyKey()
}
