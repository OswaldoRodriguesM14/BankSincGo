package interfaces_cora_api

import entity_cora_api "github.com/OswaldoRodriguesM14/BankSincGo/internal/entity/register"

type CoraAPIInterface interface {
	RetornaToken(certFileBase64, privateKeyBase64 string, clientID string) (*entity_cora_api.TokenResponse, error)
	ConsultaExtratoDoMesAtual(certFileBase64, privateKeyBase64, token string) (*[]entity_cora_api.EntryData, error)
	ConsultaExtratoDataInicioEFim(certFileBase64, privateKeyBase64, token string, inicio, fim string) (*[]entity_cora_api.EntryData, error)
	ConsultaExtratoDoMesAtualType(certFileBase64, privateKeyBase64, token string, type_ string) (*[]entity_cora_api.EntryData, error)

	EmiteBoleto(certFileBase64, privateKeyBase64, token string, Boleto *entity_cora_api.BoletoCora) (*entity_cora_api.RetornoEmissaoBoleto, error)
}
