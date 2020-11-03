package main

import (
	"eximgolog/tools"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template
var mongodb tools.MongoDB

type IndexData struct {
	PageTitle string
	Tipos     []string
	Result    []tools.LogLine
}

func ReadMainlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	teste := tools.ReadLog("mainlog")
	for _, t := range teste {
		mongodb.InsertLogLine(t)
	}
	mongodb.CloseConnection()
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", retornaIndexHandler("", nil))
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dadosForm := tools.FindForm{
		Data:    r.PostForm.Get("data"),
		Horario: r.PostForm.Get("horario"),
		Mailid:  r.PostForm.Get("mailid"),
		Tipo:    r.PostForm.Get("tipo"),
	}
	mongodb.ConnectMongoDb()
	mongodb.FindLogLine(dadosForm)
	mongodb.CloseConnection()
	http.Redirect(w, r, "/", 302)
}

func main() {

	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/readmainlog", ReadMainlogHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func retornaIndexHandler(title string, result []tools.LogLine) IndexData {
	if title == "" {
		title = "Bem-Vindo"
	}

	tipos := []string{"Enviado", "Recebido", "Redirecionado", "EntregaFailed", "EntregaAdiada", "EntregaSuprimida", "Roteada", "EmailForwarder", "Outro"}

	dados := IndexData{
		PageTitle: title,
		Tipos:     tipos,
		Result:    nil,
	}

	return dados
}
