package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Response struct {
	IP        string    `json:"ip"`
	Timestamp time.Time `json:"timestamp"`
}

func getIp(req *http.Request) string {
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	} else {
		return strings.Split(req.RemoteAddr, ":")[0]
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		res, _ := json.Marshal(Response{
			IP:        getIp(r),
			Timestamp: time.Now().UTC(),
		})

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		w.Write(res)

	})

	http.HandleFunc("/plain", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(getIp(r)))

	})

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
