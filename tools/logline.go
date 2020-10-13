package eximgolog

// LogLine - estrutura dos dados de cada linha encontrada no mainlog
type LogLine struct {
	Data       string `json:"data"`
	Horario    string `json:"horario"`
	Mailid     string `json:"mailid"`
	Redirectid string `json:"redirectid"`
	Tipo       string `json:"tipo"`
	Email      string `json:"email"`
	ErroMsg    string `json:"erromsg"`
}

// EnumType - serve para tipar os tipos de mensagem, assim facilitando a filtragem
type EnumType string

const (
	Enviado          EnumType = "=>"
	Recebido         EnumType = "<="
	Redirecionado    EnumType = "<>"
	EntregaFailed    EnumType = "**"
	EntregaAdiada    EnumType = "=="
	EntregaSuprimida EnumType = "*>"
	Roteada          EnumType = ">>"
	EmailForwarder   EnumType = "->"
	Desconhecido     EnumType = "Informação desconhecida."
)
