package dto_cora

type BoletoDTO struct {
	Code              *string `json:"codigo_interno"`
	NomeDestinatario  string  `json:"nome_pagador"`
	Email             string  `json:"email"`
	Documento         string  `json:"documento"`
	NomeCobranca      string  `json:"nome_cobranca"`
	DescricaoCobranca string  `json:"descricao_cobranca"`
	Valor             int32   `json:"valor"`
	DataVencimento    string  `json:"data_vencimento"`
}

func NewBoletoDTO() *BoletoDTO {
	return &BoletoDTO{}
}

var BoletoDTOSchema string = `
{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "title": "BoletoDTO",
    "type": "object",
    "properties": {
        "codigo_interno": {
            "type": ["string", "null"],
            "description": "Código interno do boleto, opcional."
        },
        "nome_pagador": {
            "type": "string",
            "description": "Nome, obrigatório."
        },
        "email": {
            "type": "string",
            "description": "Email, obrigatório."
        },
        "documento": {
            "type": "string",
            "description": "Documento, obrigatório."
        },
        "nome_cobranca": {
            "type": "string",
            "description": "Descrição, obrigatório."
        },
		"descricao_cobranca": {
            "type": "string",
            "description": "Descrição, obrigatório."
        },
        "valor": {
            "type": "integer",
            "minimum": 0,
            "description": "Valor, obrigatório."
        },
        "data_vencimento": {
            "type": "string",
            "format": "date",
            "description": "Data de vencimento, obrigatório."
        }
    },
    "required": ["nome_pagador", "email", "documento", "nome_cobranca", "descricao_cobranca", "valor", "data_vencimento"],
    "additionalProperties": false
}`
