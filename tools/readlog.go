package tools

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// ReadLog - Funcao que le o arquivo que foi dado como argumento e retorna um slice de cada linha lida
func ReadLog(fileName string) []LogLine {
	var partner = regexp.MustCompile(`(?m)(([0-9-]+) ([0-9:]+) ([-0-9a-zA-Z]{10,})) ([<*\-=>]+|[a-z.]+) ([<>=]+|[\[\]0-9.]+|[a-z0-9@.A-Z\:]+|[0-9a-zA-Z@. <>]+) (R=[0-9a-zA-Z-]+|[a-zA-Z ]+|U=[a-zA-Z0-9]+|H=[\(\[a-zA-Z0-9.\]\)]+) (.*)`)

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("falha ao abrir arquivo: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var loglines []string
	var loglineList []LogLine

	for scanner.Scan() {
		loglines = append(loglines, scanner.Text())
	}

	file.Close()

	for _, eachline := range loglines {
		for _, match := range partner.FindAllString(eachline, -1) {
			_, msg := CheckData(strings.Split(match, " "))
			if msg == Enviado || msg == Recebido {
				linha := LogLine{}
				linha.Data = strings.Split(match, " ")[0]
				linha.Horario = strings.Split(match, " ")[1]
				linha.Mailid = strings.Split(match, " ")[2]
				var msgType string
				var email string
				if msg == Enviado {
					msgType = "Enviado"
					email = strings.Trim(strings.Split(match, " ")[5], "<>")
				} else {
					msgType = "Recebido"
					email = strings.Split(match, " ")[4]
				}
				linha.Tipo = msgType
				linha.Email = email
				loglineList = append(loglineList, linha)
			} else if msg == Redirecionado {
				linha := LogLine{}
				linha.Data = strings.Split(match, " ")[0]
				linha.Horario = strings.Split(match, " ")[1]
				linha.Mailid = strings.Split(match, " ")[2]

				linha.Tipo = "Redirecionado"
				linha.Redirectid = strings.Split(match, " ")[5]
				loglineList = append(loglineList, linha)
			} else if msg == EntregaFailed {
				linha := LogLine{}
				linha.Data = strings.Split(match, " ")[0]
				linha.Horario = strings.Split(match, " ")[1]
				linha.Mailid = strings.Split(match, " ")[2]
				linha.Tipo = "EntregaFailed"
				linha.ErroMsg = strings.Split(match, " ")[4]
				loglineList = append(loglineList, linha)
			} else if msg == EntregaAdiada {
				linha := LogLine{}
				linha.Data = strings.Split(match, " ")[0]
				linha.Horario = strings.Split(match, " ")[1]
				linha.Mailid = strings.Split(match, " ")[2]
				linha.Tipo = "EntregaAdiada"
				linha.ErroMsg = strings.Split(match, " ")[6]
				loglineList = append(loglineList, linha)
			} else {
				msgError := ""
				linha := LogLine{}
				linha.Data = strings.Split(match, " ")[0]
				linha.Horario = strings.Split(match, " ")[1]
				linha.Mailid = strings.Split(match, " ")[2]
				linha.Tipo = "Outro"
				for i := range strings.Split(match, " ") {
					if i > 2 {
						msgError += strings.Split(match, " ")[i] + " "
					}
				}
				linha.ErroMsg = msgError
				loglineList = append(loglineList, linha)
			}
		}
	}

	return loglineList

}

// CheckData - verifica se a informação dada, esta com os dados corretos e retorna um true, caso não retorna false
func CheckData(dado []string) (bool, EnumType) {
	var tipoMensagem EnumType
	var enumTipos EnumType
	enumTipos = EnumType(dado[3])

	switch enumTipos {
	case Enviado:
		tipoMensagem = Enviado
	case Recebido:
		tipoMensagem = Recebido
		if EnumType(dado[4]) == "<>" {
			tipoMensagem = Redirecionado
		}
	case EntregaFailed:
		tipoMensagem = EntregaFailed
	case EntregaAdiada:
		tipoMensagem = EntregaAdiada
	case EntregaSuprimida:
		tipoMensagem = EntregaSuprimida
	case Roteada:
		tipoMensagem = Roteada
	case EmailForwarder:
		tipoMensagem = EmailForwarder
	default:
		tipoMensagem = Desconhecido
		return false, tipoMensagem
	}

	return true, tipoMensagem
}
