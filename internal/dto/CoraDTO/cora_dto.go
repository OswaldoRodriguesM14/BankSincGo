package dto_cora

// CertificadoDTO é uma estrutura para armazenar dados extraídos de uma requisição.
type CertificadoDTO struct {
	IDContaCORA    string `json:"conta_id"`
	IDContaINTEGRA string `json:"conta_integra_id"`
	EmpresaID      string `json:"empresa_id"`
	CertFile       string `json:"cert_file"`
	PrivateKey     string `json:"private_key"`
}

var SchemaJsonCertificadoDTO string = `
{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"title": "CertificadoDTO",
	"type": "object",
	"properties": {
	  "conta_id": {
		"type": "string",
		"description": "Identificador da conta",
		"maxLength": 512
	  },
	  "conta_integra_id": {
		"type": "string",
		"description": "Identificador da conta",
		"maxLength": 512
	  },
	  "empresa_id": {
		"type": "string",
		"description": "Identificador da conta",
		"maxLength": 512
	  },
	  "cert_file": {
		"type": "string",
		"description": "Conteúdo do arquivo de certificado em formato base64",
		"format": "byte",
		"maxLength": 5120
	  },
	  "private_key": {
		"type": "string",
		"description": "Conteúdo do arquivo de chave privada em formato base64",
		"format": "byte",
		"maxLength": 5120
	  }
	},
	"required": ["conta_id", "conta_integra_id", "empresa_id", "cert_file", "private_key"],
	"additionalProperties": false
  }
  `
