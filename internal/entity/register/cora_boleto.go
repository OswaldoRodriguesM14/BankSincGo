package entity_cora_api

// Estrutura para o documento do cliente
type Document struct {
	Identity string `json:"identity"` // Número do documento do cliente (apenas números, sem traços e pontos)
	Type     string `json:"type"`     // Tipo de documento (CPF ou CNPJ)
}

// Estrutura para o endereço do cliente
type Address struct {
	Street     string  `json:"street"`           // Rua
	Number     string  `json:"number"`           // Número
	District   string  `json:"district"`         // Bairro
	City       string  `json:"city"`             // Cidade
	State      string  `json:"state"`            // Estado
	Complement string  `json:"complement"`       // Complemento
	Coutry     *string `json:"coutry,omitempty"` // País
	ZipCode    string  `json:"zip_code"`         // CEP
}

// Estrutura para o objeto Customer
type Customer struct {
	Name     string   `json:"name"`              // Nome do cliente (obrigatório, máximo 60 caracteres)
	Email    *string  `json:"email,omitempty"`   // E-mail do cliente (opcional, máximo 60 caracteres)
	Document Document `json:"document"`          // Objeto Document (obrigatório)
	Address  *Address `json:"address,omitempty"` // Objeto Address (opcional)
}

// Estrutura para os serviços
type Service struct {
	Name        string `json:"name"`        // Nome do serviço
	Description string `json:"description"` // Descrição do serviço
	Amount      int32  `json:"amount"`      // Valor do serviço
}

// Estrutura para a multa nos termos de pagamento
type Fine struct {
	Date     string `json:"date"`     // Data da multa
	Amount   int32  `json:"amount"`   // Valor da multa
	Opcional bool   `json:"opcional"` // Indica se a multa é opcional
}

// Estrutura para o desconto nos termos de pagamento
type Discount struct {
	Type  string `json:"type"`  // Tipo do desconto (PERCENT ou FIXED)
	Value int32  `json:"value"` // Valor do desconto
}

// Estrutura para os termos de pagamento
type PaymentTerms struct {
	DueDate  string    `json:"due_date"`           // Data de vencimento
	Fine     *Fine     `json:"fine,omitempty"`     // Multa em caso de atraso
	Discount *Discount `json:"discount,omitempty"` // Desconto nos termos de pagamento
}

// Estrutura para o destino das notificações
type Destination struct {
	Name  string  `json:"name"`            // Nome do destinatário das notificações
	Email string  `json:"email"`           // E-mail do destinatário
	Phone *string `json:"phone,omitempty"` // Telefone do destinatário
}

// Estrutura para as notificações
type Notifications struct {
	Channels    [1]string   `json:"channels"`    // Canais de notificação (EMAIL, SMS, etc.)
	Destination Destination `json:"destination"` // Destino das notificações
	Rules       []string    `json:"rules"`       // Regras de notificação (NOTIFY_TWO_DAYS_BEFORE_DUE_DATE, NOTIFY_WHEN_PAID, etc.)
}

// Estrutura para o JSON principal
type BoletoCora struct {
	Code          *string        `json:"code,omitempty"`          // Código opcional definido pelo usuário
	Customer      Customer       `json:"customer"`                // Detalhes do cliente
	Services      []Service      `json:"services"`                // Lista de serviços
	PaymentTerms  PaymentTerms   `json:"payment_terms"`           // Termos de pagamento
	Notifications *Notifications `json:"notifications,omitempty"` // Notificações (opcional)
	PaymentForms  []string       `json:"payment_forms"`           // Formas de pagamento (BOLETO, PIX)
}

func NewBoletoCora(
	Code *string,
	Nome string,
	Email string,
	Documento string,
	Servico string,
	Descricao string,
	Valor int32,
	DataVencimento string,

) *BoletoCora {
	var Notification [1]string
	Notification[0] = "EMAIL"
	var Type string
	if len(Documento) == 11 {
		Type = "CPF"
	}
	if len(Documento) == 14 {
		Type = "CNPJ"
	}
	return &BoletoCora{
		Code: Code,
		Customer: Customer{
			Name:  Nome,
			Email: &Email,
			Document: Document{
				Identity: Documento,
				Type:     Type,
			},
			Address: nil,
		},
		Services: []Service{
			{
				Name:        Servico,
				Description: Descricao,
				Amount:      Valor,
			},
		},
		PaymentTerms: PaymentTerms{
			DueDate:  DataVencimento,
			Fine:     nil,
			Discount: nil,
		},
		Notifications: &Notifications{
			Channels: Notification,
			Destination: Destination{
				Name:  Nome,
				Email: Email,
				Phone: nil,
			},
			Rules: []string{"NOTIFY_WHEN_PAID", "NOTIFY_FIVE_DAYS_AFTER_DUE_DATE"},
		},
		PaymentForms: []string{"BANK_SLIP", "PIX"},
	}
}

type RetornoEmissaoBoleto struct {
	ID             string      `json:"id"`
	Status         string      `json:"status"`
	CreatedAt      string      `json:"created_at"`
	TotalAmount    int         `json:"total_amount"`
	TotalPaid      int         `json:"total_paid"`
	OccurrenceDate interface{} `json:"occurrence_date"`
	Code           string      `json:"code"`
	Customer       struct {
		Name      string      `json:"name"`
		Email     string      `json:"email"`
		Telephone interface{} `json:"telephone"`
		Document  struct {
			Identity string `json:"identity"`
			Type     string `json:"type"`
		} `json:"document"`
		Address struct {
			Street     string `json:"street"`
			Number     string `json:"number"`
			District   string `json:"district"`
			City       string `json:"city"`
			State      string `json:"state"`
			Complement string `json:"complement"`
			ZipCode    string `json:"zip_code"`
		} `json:"address"`
		Code interface{} `json:"code"`
	} `json:"customer"`
	Services []struct {
		Name        string      `json:"name"`
		Description string      `json:"description"`
		Amount      int         `json:"amount"`
		Unit        string      `json:"unit"`
		Quantity    interface{} `json:"quantity"`
		TotalAmount interface{} `json:"total_amount"`
		Code        interface{} `json:"code"`
	} `json:"services"`
	PaymentTerms struct {
		DueDate string `json:"due_date"`
		Fine    struct {
			Date   string      `json:"date"`
			Amount int         `json:"amount"`
			Rate   interface{} `json:"rate"`
		} `json:"fine"`
		Interest struct {
			Rate int32 `json:"rate"`
		} `json:"interest"`
		Discount struct {
			Percent int32 `json:"percent"`
		} `json:"discount"`
	} `json:"payment_terms"`
	PaymentOptions struct {
		BankSlip struct {
			Barcode    string `json:"barcode"`
			Digitable  string `json:"digitable"`
			OurNumber  string `json:"our_number"`
			Registered bool   `json:"registered"`
			URL        string `json:"url"`
		} `json:"bank_slip"`
	} `json:"payment_options"`
	Payments      []interface{} `json:"payments"`
	Pix           interface{}   `json:"pix"`
	Notifications struct {
		ID          string   `json:"id"`
		Channels    []string `json:"channels"`
		Destination struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"destination"`
		Schedules []struct {
			Rule        string `json:"rule"`
			Status      string `json:"status"`
			ScheduledTo string `json:"scheduledTo"`
			Active      bool   `json:"active"`
		} `json:"schedules"`
	} `json:"notifications"`
}
