package entity_cora_db

// Certificado é uma estrutura para armazenar os certificados no MongoDB.
type AccesToken struct {
	IDContaCORA   string `bson:"conta_id" json:"conta_id"`
	JWTAccesToken string `bson:"jwt" json:"jwt"`
	CertFile      string `bson:"cert_file" json:"cert_file"`
	PrivateKey    string `bson:"private_key" json:"private_key"`
}

func NewAccesToken(certFile, privateKey, idConta, JWTAccesToken string) *AccesToken {
	return &AccesToken{
		IDContaCORA:   idConta,
		JWTAccesToken: JWTAccesToken,
		CertFile:      certFile,
		PrivateKey:    privateKey,
	}
}
