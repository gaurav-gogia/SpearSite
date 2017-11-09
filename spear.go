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
	games.Windows = "https://drive.google.com/open?id=1XqZt6T4thsDgZ0Pil7gfMem0iYYHby_4"
	games.Linux = "https://drive.google.com/open?id=1atPZq4m7w8jfOK15zUxG816AKa50mmdB"
	games.MacOS = "https://drive.google.com/open?id=1i93p55Dh_fRTBeUf-IQzJwzR-ZjTWUGu"
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
