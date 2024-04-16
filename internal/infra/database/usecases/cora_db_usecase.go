package usecases_cora_db

import (
	"context"
	"errors"

	entity_cora_db "github.com/OswaldoRodriguesM14/BankSincGo/internal/infra/database/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CoraDB struct {
	Session *mongo.Client
}

func CoraDBCompose(Session *mongo.Client) *CoraDB {
	return &CoraDB{
		Session: Session,
	}
}

// SalvaCertificado salva um certificado no banco de dados MongoDB.
func (c *CoraDB) SalvaCertificadoCora(certificado *entity_cora_db.CertificadoDB) error {
	collection := c.Session.Database("SincBank").Collection("certificados_cora")

	// Insere o documento na coleção
	_, err := collection.InsertOne(context.Background(), certificado)
	if err != nil {
		return err
	}

	return nil
}

// #CDBUS35
// ConsultaCertificado consulta um certificado do banco de dados MongoDB.
func (c *CoraDB) ConsultaCertificado(id string) (*entity_cora_db.CertificadoDB, error) {
	collection := c.Session.Database("SincBank").Collection("certificados_cora")

	// Consulta o documento pelo ID
	var certificado entity_cora_db.CertificadoDB
	err := collection.FindOne(context.Background(), bson.M{"conta_integra_id": id}).Decode(&certificado)
	if err != nil {
		return nil, errors.New("Falha ao Consultar Certificado - #CDBUS35")
	}

	return &certificado, nil
}

func (c *CoraDB) SalvaAccessToken(certificado *entity_cora_db.AccesToken) error {
	collection := c.Session.Database("SincBank").Collection("access_token_cora")

	filter := bson.M{"conta_id": certificado.IDContaCORA}

	opts := options.Replace().SetUpsert(true)
	_, err := collection.ReplaceOne(context.Background(), filter, certificado, opts)
	if err != nil {
		return err
	}

	return nil
}

// #CDBUS35
// ConsultaCertificado consulta um certificado do banco de dados MongoDB.
func (c *CoraDB) ConsultaAccessToken(id string) (*entity_cora_db.AccesToken, error) {
	collection := c.Session.Database("SincBank").Collection("access_token_cora")

	// Consulta o documento pelo ID
	var certificado entity_cora_db.AccesToken
	err := collection.FindOne(context.Background(), bson.M{"conta_id": id}).Decode(&certificado)
	if err != nil {
		return nil, errors.New("Falha ao Consultar Certificado - #CDBUS35")
	}

	return &certificado, nil
}
