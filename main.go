package main

import (
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	Value       []byte
	ContentType string
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept-Encoding,Authorization,X-Forwarded-For,Content-Type,Origin,Server")

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	w.Header().Set("Content-Type", s.ContentType)
	w.Write(s.Value)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %v\n", port)
	server := Server{
		ContentType: os.Getenv("PUBLISH_CONTENT_TYPE"),
		Value:       []byte(os.Getenv("PUBLISH_VALUE")),
	}
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), server)
	fmt.Printf("Server error: %v\n", err)
}
