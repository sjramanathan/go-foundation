package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sjramanathan/go-foundation/nlp"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()

	// TODO: Run a health check
	fmt.Fprintln(w, "OK")
}

// exercise: Write a tokenizeHandler that will read the text form the requset
// body and return JSON in the format {"tokens": ["who", "on", "first"]}
func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read request", http.StatusBadRequest)
		return
	}

	tokens := nlp.Tokenize(string(data))
	resp := map[string]any{
		"tokens": tokens,
	}
	data, err = json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't encode", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	// routing
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", tokenizeHandler).Methods(http.MethodPost)
	http.Handle("/", r)

	// run server
	if err := http.ListenAndServe(":9112", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
