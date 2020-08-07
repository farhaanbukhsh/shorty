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
	Success  bool   `json:"success"`
	Response string `json:"response"`
}

func saveURL(req URLRequest, svc storage.Service) (response, int) {
	responseCode := http.StatusOK
	if req.LongURL == "" {
		res := response{Response: "Empty URL", Success: false}
		responseCode = http.StatusBadRequest
		return res, responseCode
	}
	code, err := svc.Save(req.LongURL, req.Slug)
	res := response{Response: code, Success: err == nil}
	if err != nil {
		responseCode = http.StatusConflict
		if err.Error() == "Slug Already Exists" {
			res = response{Response: err.Error(), Success: false}
		}
	}
	return res, responseCode
}

func registerHandler(svc storage.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			var urlRequest URLRequest
			var res response
			responseCode := http.StatusOK
			w.Header().Set("Content-Type", "application/json")
			err := json.NewDecoder(request.Body).Decode(&urlRequest)
			if err != nil {
				res = response{Response: "Bad Request", Success: false}
				responseCode = http.StatusBadRequest
			} else {
				res, responseCode = saveURL(urlRequest, svc)
			}

			w.WriteHeader(responseCode)
			err = json.NewEncoder(w).Encode(res)
			if err != nil {
				log.Fatalf("Could not encode response to output: %v", err)
			}
			defer request.Body.Close()

		}
	}
}

func redirectHandler(svc storage.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		code := request.URL.Path[1:]
		url, err := svc.Load(code)
		fmt.Printf("%s, %s", url, err)
		if err != nil {
			fmt.Fprintf(w, "Error getting URL, %s", err)
		} else {
			http.Redirect(w, request, url, 307)
		}
	}
}

func faviconHandler(svc storage.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		code := request.URL.Path[1:]
		log.Printf("Favicon: %s", code)
	}
}

// StartServer helps you run server
func StartServer(port string, svc storage.Service) {
	http.HandleFunc("/register", registerHandler(svc))
	http.HandleFunc("/", redirectHandler(svc))
	http.HandleFunc("/favicon.ico", faviconHandler(svc))
	log.Fatal(http.ListenAndServe(port, nil))
}
