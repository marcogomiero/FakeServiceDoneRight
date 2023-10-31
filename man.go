package main

import (
	"fmt"
	"net/http"
	"sync"
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

// Routine di test per sviluppare carico
// Da usare in caso di necessità
func startConcurrentServer() {
	var wg sync.WaitGroup
	numRequests := 1000
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://localhost:8080")
			if err != nil {
				fmt.Printf("Errore nella richiesta: %s\n", err)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("Status code: %d\n", resp.StatusCode)
		}()
	}
	wg.Wait()
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I'm alive and healthy!"))
}
