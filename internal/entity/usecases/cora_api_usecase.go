package usecases_cora_api

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	entity_cora_api "github.com/OswaldoRodriguesM14/BankSincGo/internal/entity/register"
	pkg "github.com/OswaldoRodriguesM14/BankSincGo/pkg/utilits"
)

type CoraAPI struct {
	Utilits pkg.UtilitsInterface
}

func CoraAPICompose(Utilits pkg.UtilitsInterface) *CoraAPI {
	return &CoraAPI{
		Utilits: Utilits,
	}
}

func configurarClienteTLS(certFileBase64, privateKeyBase64 string) (*http.Client, error) {
	// Decodificar certFileBase64 e privateKeyBase64 de volta para bytes
	certFileBytes, err := base64.StdEncoding.DecodeString(certFileBase64)
	if err != nil {
		return nil, fmt.Errorf("falha ao decodificar certFileBase64: %w", err)
	}

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("falha ao decodificar privateKeyBase64: %w", err)
	}

	// Carregar o par de chaves/certificados a partir dos bytes decodificados
	cert, err := tls.X509KeyPair(certFileBytes, privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("falha ao carregar o par de chaves/certificados: %w", err)
	}

	// Configurar o cliente HTTP com o certificado TLS
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: false, // Certifique-se de verificar os certificados do servidor
			},
		},
		Timeout: 30 * time.Second, // Definir um timeout para o cliente HTTP
	}

	return client, nil
}

func (c *CoraAPI) RetornaToken(certFileBase64, privateKeyBase64 string, clientID string) (*entity_cora_api.TokenResponse, error) {
	// Configura o cliente HTTP com o certificado TLS
	client, err := configurarClienteTLS(certFileBase64, privateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("Erro ao configurar o cliente TLS: %w", err)
	}

	urlCora := fmt.Sprintf("%s/token", pkg.Url)

	// Configura os dados da solicitação
	formData := url.Values{
		"grant_type": {"client_credentials"},
		"client_id":  {clientID},
	}
	requestData := bytes.NewBufferString(formData.Encode())

	// Cria uma solicitação POST
	req, err := http.NewRequest("POST", urlCora, requestData)
	if err != nil {
		return nil, fmt.Errorf("Erro ao criar a solicitação: %w", err)
	}

	// Configura o cabeçalho "Content-Type" como "application/x-www-form-urlencoded"
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Envia a solicitação e recebe a resposta
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Erro ao enviar a solicitação: %w", err)
	}
	defer resp.Body.Close()

	// Lê a resposta usando io.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler a resposta: %w", err)
	}

	// Analisa o JSON da resposta
	var responseData entity_cora_api.TokenResponse
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, fmt.Errorf("Erro ao analisar o JSON da resposta: %w", err)
	}

	return &responseData, nil
}

func (c *CoraAPI) ConsultaExtratoDoMesAtual(certFileBase64, privateKeyBase64, token string) (*[]entity_cora_api.EntryData, error) {
	// Configura o cliente HTTP com o certificado TLS
	client, err := configurarClienteTLS(certFileBase64, privateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("Erro ao configurar o cliente TLS: %w", err)
	}
	PrimeiroDiaMes, UltimoDiaMes := c.Utilits.PrimeiroUltimoDiaMesAtual()
	Page := 1
	PageTotal := 5000
	URLCompose := entity_cora_api.NewExtratoRequestParms(PrimeiroDiaMes, UltimoDiaMes, nil, nil, Page, PageTotal, nil)
	URL, _ := URLCompose.BuildURL()

	// Cria uma requisição GET com a URL modificada
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Envia a requisição e obtém a resposta
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição: %v", err)
	}

	// Verifica se a resposta foi bem-sucedida (código 200)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("resposta não foi bem-sucedida: código %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler a resposta: %w", err)
	}

	// Analisa o JSON da resposta
	var responseData entity_cora_api.CoraExtrato
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, fmt.Errorf("Erro ao analisar o JSON da resposta: %w", err)
	}
	Retorno := responseData.GetEntryData()

	return Retorno, nil
}

func (c *CoraAPI) ConsultaExtratoDataInicioEFim(certFileBase64, privateKeyBase64, token string, inicio, fim string) (*[]entity_cora_api.EntryData, error) {
	// Configura o cliente HTTP com o certificado TLS
	client, err := configurarClienteTLS(certFileBase64, privateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("Erro ao configurar o cliente TLS: %w", err)
	}

	Page := 1
	PageTotal := 5000
	URLCompose := entity_cora_api.NewExtratoRequestParms(inicio, fim, nil, nil, Page, PageTotal, nil)
	URL, _ := URLCompose.BuildURL()

	// Cria uma requisição GET com a URL modificada
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Envia a requisição e obtém a resposta
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição: %v", err)
	}

	// Verifica se a resposta foi bem-sucedida (código 200)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("resposta não foi bem-sucedida: código %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler a resposta: %w", err)
	}

	// Analisa o JSON da resposta
	var responseData entity_cora_api.CoraExtrato
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, fmt.Errorf("Erro ao analisar o JSON da resposta: %w", err)
	}
	Retorno := responseData.GetEntryData()

	return Retorno, nil
}

func (c *CoraAPI) ConsultaExtratoDoMesAtualType(certFileBase64, privateKeyBase64, token string, type_ string) (*[]entity_cora_api.EntryData, error) {
	// Configura o cliente HTTP com o certificado TLS
	client, err := configurarClienteTLS(certFileBase64, privateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("Erro ao configurar o cliente TLS: %w", err)
	}
	PrimeiroDiaMes, UltimoDiaMes := c.Utilits.PrimeiroUltimoDiaMesAtual()
	Page := 1
	PageTotal := 5000
	URLCompose := entity_cora_api.NewExtratoRequestParms(PrimeiroDiaMes, UltimoDiaMes, &type_, nil, Page, PageTotal, nil)
	URL, _ := URLCompose.BuildURL()

	// Cria uma requisição GET com a URL modificada
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Envia a requisição e obtém a resposta
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição: %v", err)
	}

	// Verifica se a resposta foi bem-sucedida (código 200)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("resposta não foi bem-sucedida: código %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler a resposta: %w", err)
	}

	// Analisa o JSON da resposta
	var responseData entity_cora_api.CoraExtrato
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, fmt.Errorf("Erro ao analisar o JSON da resposta: %w", err)
	}
	Retorno := responseData.GetEntryData()

	return Retorno, nil
}
func (c *CoraAPI) EmiteBoleto(certFileBase64, privateKeyBase64, token string, Boleto *entity_cora_api.BoletoCora) (*entity_cora_api.RetornoEmissaoBoleto, error) {
	// Configura o cliente HTTP com o certificado TLS
	client, err := configurarClienteTLS(certFileBase64, privateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("Erro ao configurar o cliente TLS: %w", err)
	}

	urlCora := fmt.Sprintf("%s/invoices/", pkg.Url)

	boletoJSON, err := json.Marshal(Boleto)
	if err != nil {
		return nil, fmt.Errorf("Erro ao serializar BoletoCora em JSON: %w", err)
	}

	// Crie a solicitação POST com os dados de BoletoCora em formato JSON
	req, err := http.NewRequest("POST", urlCora, bytes.NewBuffer(boletoJSON))
	if err != nil {
		return nil, fmt.Errorf("Erro ao criar requisição POST: %w", err)
	}

	// Define os cabeçalhos da requisição
	IdempotencyKey := c.Utilits.GeraIdempotencyKey()
	req.Header.Set("Idempotency-Key", IdempotencyKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Envia a solicitação e recebe a resposta
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Erro ao enviar a solicitação: %w", err)
	}
	defer resp.Body.Close()

	// Lê a resposta usando io.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler a resposta: %w", err)
	}

	// Analisa o JSON da resposta
	var responseData entity_cora_api.RetornoEmissaoBoleto
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, fmt.Errorf("Erro ao analisar o JSON da resposta: %w", err)
	}

	return &responseData, nil
}
