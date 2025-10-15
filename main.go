package main

import (
	"log"
	"net/http"
	"net/mail"
	"os"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func writeTo(email string) {
	var fileName = "emails.txt"
	f, err := os.OpenFile("logs/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(email + "\n"); err != nil {
		log.Fatal(err)
	}
}

func applyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")

	if !valid(email) {
		http.Error(w, "Invalid email address", http.StatusBadRequest)
		return
	}

	writeTo(email)

	http.ServeFile(w, r, "html/apply.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/apply", applyHandler)

	port := ":8000"
	log.Printf("Server running on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
