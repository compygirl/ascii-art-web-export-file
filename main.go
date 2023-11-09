package main

import (
	ascii "asciiweb/ascii"
	asciifuncs "asciiweb/ascii/funcs"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type text struct {
	input string
	font  string
}

var info text

func get_req(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound, "")
		return
	}

	switch r.Method {
	case "GET":

		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err := t.Execute(w, nil); err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		Errors(w, http.StatusMethodNotAllowed, "")
		return
	}
}

func post_req(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		t, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := r.ParseForm(); err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		input, ok := r.Form["w3review1"]
		if !ok {
			Errors(w, http.StatusBadRequest, "Bad typing")
			return
		}
		font, ok := r.Form["font"]
		if !ok {
			Errors(w, http.StatusBadRequest, "Incorrect font")
			return
		}
		info.input = input[0]
		info.font = font[0]
		res, err := ascii.Asciiart(font[0], input[0])

		if err != nil && err.Error() == "ERROR: the string is invalid (NOT ASCII RANGE)" {
			Errors(w, http.StatusBadRequest, "Bad typing")
			return
		}
		if err != nil && err.Error() == "ERROR: not valid font or didn't choose font" {
			Errors(w, http.StatusBadRequest, "Incorrect font")
			return
		}
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		t.Execute(w, res)
	default:
		Errors(w, http.StatusMethodNotAllowed, "")
		return
	}
}

func exportHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the input and font from the form.
	r.ParseForm()
	input := info.input
	font := info.font

	if err := asciifuncs.IsValid(input, font); err != nil {
		Errors(w, http.StatusBadRequest, err.Error())
		return
	}

	// Generate the ASCII art result using the selected font.
	res, err := ascii.Asciiart(font, input)
	if err != nil {
		Errors(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Set the appropriate HTTP headers for file export.
	w.Header().Set("Content-Type", "text/csv")                                  // Set the appropriate MIME type for CSV format.
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt") // Set the filename for download.
	w.Header().Set("Content-Length", strconv.Itoa(len(res)))                    // Set the content length.

	// Write the ASCII art result to the response.
	if _, err := io.WriteString(w, res); err != nil { // w is the file now and res(ascii art result) will be written to the file.
		Errors(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func main() {
	styles := http.FileServer(http.Dir("./assets/styles"))
	http.Handle("/assets/styles/", http.StripPrefix("/assets/styles/", styles))

	http.HandleFunc("/", get_req)
	http.HandleFunc("/ascii-art", post_req)
	http.HandleFunc("/export", exportHandler)
	fmt.Printf("Starting server got testing... http://127.0.0.1:8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func Errors(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, strconv.Itoa(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}
	res := struct {
		StatusCodeAndText string
		MessageError      string
	}{
		StatusCodeAndText: strconv.Itoa(status) + " " + http.StatusText(status),
		MessageError:      message,
	}
	if err := t.Execute(w, res); err != nil {
		http.Error(w, strconv.Itoa(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}
}
