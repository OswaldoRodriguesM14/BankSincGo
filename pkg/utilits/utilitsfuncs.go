package pkg

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// MarshalValidationError converte uma mensagem de erro em um JSON.
func MarshalValidationError(err error) []byte {
	type ValidationError struct {
		Error string `json:"error"`
	}
	errMsg := &ValidationError{Error: err.Error()}
	json, _ := json.Marshal(errMsg)
	return json
}

func PrimeiroUltimoDiaMesAtual() (startDate string, endDate string) {
	currentTime := time.Now()
	firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	nextMonth := firstDayOfMonth.AddDate(0, 1, 0)
	lastDayOfMonth := nextMonth.AddDate(0, 0, -1)
	start := firstDayOfMonth.Format("2006-01-02") // Formata a data para o formato YYYY-MM-DD
	end := lastDayOfMonth.Format("2006-01-02")    // Formata a data para o formato YYYY-MM-DD
	return start, end

}

func GeraIdempotencyKey() string {
	return uuid.New().String()
}
