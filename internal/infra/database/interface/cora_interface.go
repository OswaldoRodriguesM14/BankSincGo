package interferface_db

import (
	entity_cora_db "github.com/OswaldoRodriguesM14/BankSincGo/internal/infra/database/entity"
)

type CoraDB interface {
	SalvaCertificadoCora(certificado *entity_cora_db.CertificadoDB) error
	ConsultaCertificado(id string) (*entity_cora_db.CertificadoDB, error)
	SalvaAccessToken(certificado *entity_cora_db.AccesToken) error
	ConsultaAccessToken(id string) (*entity_cora_db.AccesToken, error)
}
