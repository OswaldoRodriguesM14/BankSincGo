package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	dto_cora "github.com/OswaldoRodriguesM14/BankSincGo/internal/dto/CoraDTO"
	interfaces_cora_api "github.com/OswaldoRodriguesM14/BankSincGo/internal/entity/interfaces"
	entity_cora_api "github.com/OswaldoRodriguesM14/BankSincGo/internal/entity/register"
	entity_cora_db "github.com/OswaldoRodriguesM14/BankSincGo/internal/infra/database/entity"
	interface_db "github.com/OswaldoRodriguesM14/BankSincGo/internal/infra/database/interface"
	pkg "github.com/OswaldoRodriguesM14/BankSincGo/pkg/utilits"
	"github.com/go-chi/chi"
)

type CoraHandler struct {
	CoraDB  interface_db.CoraDB
	CoraAPI interfaces_cora_api.CoraAPIInterface
	Utilits pkg.UtilitsInterface
}

func CoraHandlerComposer(coraDB interface_db.CoraDB, coraAPI interfaces_cora_api.CoraAPIInterface, utilits pkg.UtilitsInterface) *CoraHandler {
	return &CoraHandler{
		CoraDB:  coraDB,
		CoraAPI: coraAPI,
		Utilits: utilits,
	}
}

func (c *CoraHandler) SalvaCertificadoCora(w http.ResponseWriter, r *http.Request) {

	body, err := c.Utilits.ValidateBody(r, dto_cora.SchemaJsonCertificadoDTO)
	if err != nil {
		Json := c.Utilits.MarshalValidationError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Json)
		return
	}

	//#CH34
	var CertificadoDTO dto_cora.CertificadoDTO
	err = json.Unmarshal(body, &CertificadoDTO)
	if err != nil {
		Json := c.Utilits.MarshalValidationError(errors.New("Falha ao Deserealizar o JSON - #CH34"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	NewCertificado := entity_cora_db.NewCertificado(
		CertificadoDTO.CertFile,
		CertificadoDTO.PrivateKey,
		CertificadoDTO.IDContaCORA,
		CertificadoDTO.IDContaINTEGRA,
		CertificadoDTO.EmpresaID,
	)

	err = c.CoraDB.SalvaCertificadoCora(NewCertificado)
	if err != nil {
		Json := c.Utilits.MarshalValidationError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (c *CoraHandler) GeraAccesToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Url := chi.URLParam(r, "id")

	Certificado, err := c.CoraDB.ConsultaCertificado(Url)
	if err != nil {
		Json := c.Utilits.MarshalValidationError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	TokenCora, err := c.CoraAPI.RetornaToken(Certificado.CertFile, Certificado.PrivateKey, Certificado.IDContaCORA)
	if err != nil {
		Json := c.Utilits.MarshalValidationError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	err = c.CoraDB.SalvaAccessToken(
		entity_cora_db.NewAccesToken(
			Certificado.CertFile,
			Certificado.PrivateKey,
			Certificado.IDContaCORA,
			TokenCora.AccessToken,
		))
	if err != nil {
		Json := c.Utilits.MarshalValidationError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (c *CoraHandler) ConsultaExtrato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	AccesToken, _ := c.CoraDB.ConsultaAccessToken("int-7hbUt8b6hNNihkOnRdHEhS")
	if AccesToken == nil {
		Json := c.Utilits.MarshalValidationError(errors.New("Falha ao Consultar o Access Token - #CE35"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	Extrato, err := c.CoraAPI.ConsultaExtratoDoMesAtual(AccesToken.CertFile, AccesToken.PrivateKey, AccesToken.JWTAccesToken)
	if err != nil {
		Json := c.Utilits.MarshalValidationError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	json.NewEncoder(w).Encode(Extrato)
	w.WriteHeader(http.StatusOK)
}

func (c *CoraHandler) ConsultaExtratoType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	URL := chi.URLParam(r, "id")

	AccesToken, _ := c.CoraDB.ConsultaAccessToken("int-7hbUt8b6hNNihkOnRdHEhS")
	if AccesToken == nil {
		Json := c.Utilits.MarshalValidationError(errors.New("Falha ao Consultar o Access Token - #CE35"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	Extrato, err := c.CoraAPI.ConsultaExtratoDoMesAtualType(AccesToken.CertFile, AccesToken.PrivateKey, AccesToken.JWTAccesToken, URL)
	if err != nil {
		Json := c.Utilits.MarshalValidationError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	json.NewEncoder(w).Encode(Extrato)
	w.WriteHeader(http.StatusOK)
}

func (c *CoraHandler) EmiteBoleto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	NewBoletoDTO := dto_cora.NewBoletoDTO()

	AccesToken, _ := c.CoraDB.ConsultaAccessToken("int-7hbUt8b6hNNihkOnRdHEhS")
	if AccesToken == nil {
		Json := c.Utilits.MarshalValidationError(errors.New("Falha ao Consultar o Access Token - #CE35"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	body, err := c.Utilits.ValidateBody(r, dto_cora.BoletoDTOSchema)
	if err != nil {
		Json := c.Utilits.MarshalValidationError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Json)
		return
	}

	err = json.Unmarshal(body, &NewBoletoDTO)
	if err != nil {
		Json := c.Utilits.MarshalValidationError(errors.New("Falha ao Deserealizar o JSON - #CE35"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	NewBoleto := entity_cora_api.NewBoletoCora(
		NewBoletoDTO.Code,
		NewBoletoDTO.NomeDestinatario,
		NewBoletoDTO.Email,
		NewBoletoDTO.Documento,
		NewBoletoDTO.NomeCobranca,
		NewBoletoDTO.DescricaoCobranca,
		NewBoletoDTO.Valor,
		NewBoletoDTO.DataVencimento,
	)

	BoletoEmitido, err := c.CoraAPI.EmiteBoleto(
		AccesToken.CertFile,
		AccesToken.PrivateKey,
		AccesToken.JWTAccesToken,
		NewBoleto,
	)

	if err != nil {
		Json := c.Utilits.MarshalValidationError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Json)
		return
	}

	json.NewEncoder(w).Encode(BoletoEmitido)
	w.WriteHeader(http.StatusOK)

}
