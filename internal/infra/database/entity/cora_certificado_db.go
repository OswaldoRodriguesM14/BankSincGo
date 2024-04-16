package entity_cora_db

import "go.mongodb.org/mongo-driver/bson/primitive"

// Certificado é uma estrutura para armazenar os certificados no MongoDB.
type CertificadoDB struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	IDContaCORA    string             `bson:"conta_id" json:"conta_id"`
	IDContaINTEGRA string             `bson:"conta_integra_id" json:"conta_integra_id"`
	EmpresaID      string             `bson:"empresa_id" json:"empresa_id"`
	CertFile       string             `bson:"cert_file" json:"cert_file"`
	PrivateKey     string             `bson:"private_key" json:"private_key"`
}

func NewCertificado(certFile, privateKey, idConta, idIntegra, idEmpresa string) *CertificadoDB {
	return &CertificadoDB{
		ID:             primitive.NewObjectID(),
		IDContaCORA:    idConta,
		IDContaINTEGRA: idIntegra,
		EmpresaID:      idEmpresa,
		CertFile:       certFile,
		PrivateKey:     privateKey,
	}
}
