package entity_cora_api

import (
	"fmt"
	"net/url"

	pkg "github.com/OswaldoRodriguesM14/BankSincGo/pkg/utilits"
)

type CoraExtrato struct {
	Start struct {
		Date    string `json:"date"`
		Balance int    `json:"balance"`
	} `json:"start"`
	Entries []struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		Amount      int    `json:"amount"`
		CreatedAt   string `json:"createdAt"`
		Transaction struct {
			ID           string `json:"id"`
			Type         string `json:"type"`
			Description  string `json:"description"`
			CounterParty struct {
				Name     string `json:"name"`
				Identity string `json:"identity"`
			} `json:"counterParty"`
		} `json:"transaction"`
	} `json:"entries"`
	End struct {
		Date    string `json:"date"`
		Balance int    `json:"balance"`
	} `json:"end"`
	Aggregations struct {
		CreditTotal int `json:"creditTotal"`
		DebitTotal  int `json:"debitTotal"`
	} `json:"aggregations"`
	Header struct {
		BusinessName     string `json:"businessName"`
		BusinessDocument string `json:"businessDocument"`
	} `json:"header"`
}

type EntryData struct {
	ID          string `json:"id_cora"`
	Type        string `json:"tipo"`
	Amount      int    `json:"valor"`
	Description string `json:"descricao"`
	CreatedAt   string `json:"data_criacao"`
	Name        string `json:"nome"`
	Identity    string `json:"documento"`
}

func (ce *CoraExtrato) GetEntryData() *[]EntryData {
	var entryDataList []EntryData

	for _, entry := range ce.Entries {
		var entryData EntryData

		// Verifica o tipo de entrada (CREDIT ou DEBIT) e define o tipo correspondente
		switch entry.Type {
		case "CREDIT":
			entryData = EntryData{
				ID:          entry.ID,
				Type:        "ENTRADA",
				Amount:      entry.Amount,
				Description: entry.Transaction.Description,
				CreatedAt:   entry.CreatedAt,
				Name:        entry.Transaction.CounterParty.Name,
				Identity:    entry.Transaction.CounterParty.Identity,
			}
		case "DEBIT":
			entryData = EntryData{
				ID:          entry.ID,
				Type:        "SAIDA",
				Amount:      entry.Amount,
				Description: entry.Transaction.Description,
				CreatedAt:   entry.CreatedAt,
				Name:        entry.Transaction.CounterParty.Name,
				Identity:    entry.Transaction.CounterParty.Identity,
			}
		}

		// Adiciona o EntryData à lista
		entryDataList = append(entryDataList, entryData)
	}

	return &entryDataList
}

type ExtratoRequestParmsExtrato struct {
	BaseURL         string  // URL base para a requisição
	Start           string  // Data início, no formato YYYY-MM-DD
	End             string  // Data final, no formato YYYY-MM-DD
	Type            *string // Forma de transação no extrato
	TransactionType *string // Tipo da transação
	Page            int     // Número da página
	PerPage         int     // Número de itens por página
	Aggr            *bool   // Permite incluir ou omitir o objeto Aggregations na resposta
}

func NewExtratoRequestParms(
	start string,
	end string,
	type_ *string,
	transactionType *string,
	page int,
	perPage int,
	aggr *bool,
) *ExtratoRequestParmsExtrato {
	return &ExtratoRequestParmsExtrato{
		BaseURL:         fmt.Sprintf("%s/bank-statement/statement", pkg.Url),
		Start:           start,
		End:             end,
		Type:            type_,
		TransactionType: transactionType,
		Page:            page,
		PerPage:         perPage,
		Aggr:            aggr,
	}
}

// BuildURL constrói a URL com base nos parâmetros fornecidos na struct
func (rp *ExtratoRequestParmsExtrato) BuildURL() (string, error) {
	// Cria um objeto URL com o baseURL fornecido
	u, err := url.Parse(rp.BaseURL)
	if err != nil {
		return "", err
	}

	// Cria um objeto query para armazenar os parâmetros de consulta
	query := u.Query()

	// Adiciona os parâmetros opcionais à consulta se eles forem fornecidos
	if rp.Start != "" {
		query.Set("start", rp.Start)
	}
	if rp.End != "" {
		query.Set("end", rp.End)
	}
	if rp.Type != nil {
		query.Set("type", *rp.Type)
	}
	if rp.TransactionType != nil {
		query.Set("transaction_type", *rp.TransactionType)
	}
	if rp.Page != 0 {
		query.Set("page", fmt.Sprintf("%d", rp.Page))
	}
	if rp.PerPage != 0 {
		query.Set("perPage", fmt.Sprintf("%d", rp.PerPage))
	}
	if rp.Aggr != nil {
		query.Set("aggr", fmt.Sprintf("%t", *rp.Aggr))
	}

	// Define a consulta (query) da URL com os parâmetros
	u.RawQuery = query.Encode()

	// Retorna a URL completa como uma string
	return u.String(), nil
}
