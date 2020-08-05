package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/farhaanbukhsh/shorty/storage"
)

// URLRequest stores the information to be handled by the request
type URLRequest struct {
	LongURL string `json:"url"`
	Slug    string `json:"slug"`
}

type response struct {
	Success bool   `json:"success"`
	Code    string `json:"response"`
}

func registerHandler(svc storage.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			var urlRequest URLRequest
			err := json.NewDecoder(request.Body).Decode(&urlRequest)
			if err != nil {
				fmt.Printf("Error Occured While Processing Request")
			}
			code, err := svc.Save(urlRequest.LongURL, urlRequest.Slug)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(response{Code: code, Success: err == nil})
			if err != nil {
				log.Printf("could not encode response to output: %v", err)
			}
			defer request.Body.Close()

		}
	}
}

// StartServer helps you run server
func StartServer(port string, svc storage.Service) {
	http.HandleFunc("/register", registerHandler(svc))
	log.Fatal(http.ListenAndServe(port, nil))
}
