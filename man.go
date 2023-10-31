package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/", fakeHandler)
	port := 8080
	fmt.Println(`
	  ,_,   
	 {O,o}
	 /)__)
	=="="==`)

	fmt.Printf("FakeServiceDoneRight listening on port %d...\n", port)
	go func() {
		for {
			fmt.Println("I'm alive")
			time.Sleep(1 * time.Minute)
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

// Gestione concorrente delle richieste usando goroutine
func startConcurrentServer() {
	var wg sync.WaitGroup
	numRequests := 1000 // Modifica questo valore al numero desiderato di richieste concorrenti
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
