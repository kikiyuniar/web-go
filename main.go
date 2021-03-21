package main

import (
	"log"
	"net/http"
)

func main()  {
		mux := http.NewServeMux()

		mux.HandleFunc("/hello", helloHandler)
		mux.HandleFunc("/contactus", contactusHandler)
		
		log.Println("Starting web on port 8080")

		err:= http.ListenAndServe(":8080", mux)
		log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello World, in golang web"))
}
func contactusHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("my WA 081334131617"))
}