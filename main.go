package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main()  {
		mux := http.NewServeMux()
// cara 1 route
		aboutHandler := func(w http.ResponseWriter, r *http.Request)  {
			w.Write([]byte("About Page"))
		}
// cara ke 2 route
		mux.HandleFunc("/", homeHandler)
		mux.HandleFunc("/hello", helloHandler)
		mux.HandleFunc("/contactus", contactusHandler)
		mux.HandleFunc("/product", productHandler)

//  lanjutan cara ke 1
		mux.HandleFunc("/about", aboutHandler)

// cara ke 3 route
		mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request){
			w.Write([]byte("Profileku"))
		})
		
		log.Println("Starting web on port 8080")

		err:= http.ListenAndServe(":8080", mux)
		log.Fatal(err)
}
// lanjutan cara ke 2 route untuk function
func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to Home"))
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world in golang webserver"))
}
func contactusHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("my WA 081334131617"))
}
func productHandler(w http.ResponseWriter, r *http.Request){
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