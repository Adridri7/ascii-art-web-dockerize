package handlers

import (
	"ascii"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var Text string

func RenderTemplate(w http.ResponseWriter, tmpl string, s string) {
	t, err := template.ParseFiles("./web/templates/" + tmpl + ".html")
	if err != nil {
		// En cas d'erreur lors du parsing du template
		RenderTemplate(w, "error-page", "Error "+strconv.Itoa(http.StatusBadRequest)+" : Oops, are you looking for a ghost..?")
		log.Printf("HTTP Response Code : %v", (http.StatusBadRequest))
		return
	}

	if err := t.Execute(w, s); err != nil {
		RenderTemplate(w, "error-page", "Error "+strconv.Itoa(http.StatusBadRequest)+" : Oops, are you looking for a ghost..?")
		log.Printf("HTTP Response Code : %v", (http.StatusBadRequest))
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderTemplate(w, "error-page", "Error "+strconv.Itoa(http.StatusNotFound)+" : Oops, are you looking for a ghost..?")
		log.Printf("HTTP Response Code : %v", (http.StatusNotFound))
		return
	}

	switch r.Method {
	case "GET":
		RenderTemplate(w, "home", "")
		log.Printf("HTTP Response Code : %v", (http.StatusOK))
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		text, theme := r.FormValue("input"), r.FormValue("themes")
		lines := ascii.ThemeToLines(theme)
		input, err := ascii.GetTextInput(text)

		if err != nil {
			RenderTemplate(w, "error-page", "Error "+strconv.Itoa(http.StatusInternalServerError)+" : Oops, looks like you chose a bad character...")
			log.Printf("HTTP Response Code : %v", (http.StatusInternalServerError))
		} else {
			Text = ascii.PrintAsciiArt(input, lines)
			SaveOutput(w, Text)
			http.Redirect(w, r, "/ascii-art", http.StatusSeeOther)
		}

	default:
		RenderTemplate(w, "error-page", "Error "+strconv.Itoa(http.StatusMethodNotAllowed)+" : Oops, looks like you're not allowed to do this..")
		log.Printf("HTTP Response Code : %v", (http.StatusMethodNotAllowed))
	}
}

func DisplayResult(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		RenderTemplate(w, "error-page", "Error "+strconv.Itoa(http.StatusNotFound)+" : Oops, are you looking for a ghost..?")
		return
	}

	switch r.Method {
	case "GET":
		log.Printf("HTTP Response Code : %v", (http.StatusOK))
		RenderTemplate(w, "ascii-art", Text)

		//w.Header().Set("Content-Disposition", "attachment; filename=res.txt")
		//w.Write()
	default:
		RenderTemplate(w, "error-page", "Error "+strconv.Itoa(http.StatusMethodNotAllowed)+" : Oops, looks like you're not allowed to do this..")
		log.Printf("HTTP Response Code : %v", (http.StatusMethodNotAllowed))
	}
}

func SaveOutput(w http.ResponseWriter, text string) {
	output := CreateFile(w, "./web/download/res.txt")
	output.WriteString(text)
}

func CreateFile(w http.ResponseWriter, path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		RenderTemplate(w, "error-page", "Error "+strconv.Itoa(http.StatusInternalServerError)+" : Oops, looks like you chose a bad character...")
		log.Printf("HTTP Response Code : %v", (http.StatusInternalServerError))
	}
	return file
}
