package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", fakeHandler)
	http.HandleFunc("/healthz", healthzHandler)
	port := 8080
	fmt.Println(`
	  ,_,   
	 {O,o}
	 /)__)
	=="="==`)

	fmt.Printf("FakeServiceDoneRight listening on port %d...\n", port)
	// Non serve a nulla se non a loggare ogni 5 minuti che è ancora vivo
	go func() {
		for {
			fmt.Println("I'm alive")
			time.Sleep(5 * time.Minute)
		}
	}()
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Errore nel server: %s\n", err)
	}
}

func fakeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("I'm alive and healthy!"))
	if err != nil {
		return
	}
}
