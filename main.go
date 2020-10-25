package main

import (
	eximgolog "eximgolog/tools"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

type IndexData struct {
	PageTitle string
	Tipos     []string
}

func ReadMainlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//var mylog eximgolog.LogLine
	//eximgolog.InsertLogLine(mylog)
	teste := eximgolog.ReadLog("mainlog")
	for _, t := range teste {
		eximgolog.InsertLogLine(t)
	}
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	dados := IndexData{
		PageTitle: "Bem-Vindo",
		Tipos:     []string{"Enviado", "Recebido", "Redirecionado", "EntregaFailed", "EntregaAdiada", "EntregaSuprimida", "Roteada", "EmailForwarder", "Desconhecido"},
	}

	templates.ExecuteTemplate(w, "index.html", dados)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dadosForm := eximgolog.FindForm{
		Data:    r.PostForm.Get("data"),
		Horario: r.PostForm.Get("horario"),
		Mailid:  r.PostForm.Get("mailid"),
		Tipo:    r.PostForm.Get("tipo"),
	}

	eximgolog.FindLogLine(dadosForm)
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
