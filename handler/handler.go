package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// lanjutan cara ke 2 route untuk function
func HomeHandler(w http.ResponseWriter, r *http.Request){
	log.Printf(r.URL.Path)
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to Home"))
	tmpl, err := template.ParseFiles(path.Join("views","index.html"))
	err = tmpl.Execute(w, nil)
	if err != nil {
		// menapilkan error untuk developer
		log.Println(err)
		// menampilkan error untuk browser/user
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
	
	data := map[string]interface{}{
		"title": "I'm Learning Golang Web",
		"content": "I'm Learning Golang Web with Kiki Yuniar K",
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		// menapilkan error untuk developer
		log.Println(err)
		// menampilkan error untuk browser/user
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
	
}
func HelloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world in golang webserver"))
}
func ContactusHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("my WA 081334131617"))
}
func ProductHandler(w http.ResponseWriter, r *http.Request){
	// menangkap ID
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// menampilkan ID yang ditangkap
	fmt.Fprintf(w, "Product page : %d", idNumb)
}