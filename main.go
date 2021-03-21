package main

import (
	"log"
	"net/http"
	"webgo/handler"
)

func main()  {
		mux := http.NewServeMux()
// cara 1 route
		aboutHandler := func(w http.ResponseWriter, r *http.Request)  {
			w.Write([]byte("About Page"))
		}
// cara ke 2 route => yang paling sering digunakan
		mux.HandleFunc("/", handler.HomeHandler)
		mux.HandleFunc("/hello", handler.HelloHandler)
		mux.HandleFunc("/contactus", handler.ContactusHandler)
		mux.HandleFunc("/product", handler.ProductHandler)

//  lanjutan cara ke 1
		mux.HandleFunc("/about", aboutHandler)

// cara ke 3 route
		mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request){
			w.Write([]byte("Profileku"))
		})

// memberikan akses ke server dengan port 8080
		log.Println("Starting web on port 8080")

		err:= http.ListenAndServe(":8080", mux)
		log.Fatal(err)
}
