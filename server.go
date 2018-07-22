package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/andreccosta/hooksrv/config"
)

func main() {
	addr, ok := os.LookupEnv("PORT")
	if !ok {
		addr = "8080"
	}

	cfg := config.Load("config.json")

	for url, hook := range cfg.Hooks {
		http.HandleFunc("/"+url, hook.Handler)
	}

	log.Printf("Listening on %s...", addr)
	err := http.ListenAndServe(":"+addr, logRequest(http.DefaultServeMux))

	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s / %s - %s - %s - %s", r.Host, r.Header.Get("X-Real-IP"), r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
