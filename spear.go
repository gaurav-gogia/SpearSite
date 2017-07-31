package spear

import (
	"html/template"
	"log"
	"net/http"
)

type gameFiles struct {
	Windows string
	Linux   string
	MacOS   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.html"))

	http.Handle("/assets/",
		http.StripPrefix("/assets",
			http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/game", game)
}

func index(w http.ResponseWriter, r *http.Request) {
	var games gameFiles
	games.Windows = "https://drive.google.com/open?id=0B4GT2dEsAtCwYURWdGp4dU9SMG8"
	games.Linux = "https://drive.google.com/open?id=0B4GT2dEsAtCwQkJVVmsxX1JtX1U"
	games.MacOS = "https://drive.google.com/open?id=0B4GT2dEsAtCwM29PTzdIRHhNaUU"
	err := tpl.ExecuteTemplate(w, "index.html", games)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.URL.Path)
}

func game(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "game.html", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.URL.Path)
}
