package main

import (
	"context"
	"encoding/json"
	client "github.com/wooobo/go-upbit-client/pkg/public"
	"log"
	"net/http"
)

func marketsHandler(writer http.ResponseWriter, request *http.Request) {
	upbit := client.NewClient(client.Config{
		BaseUrl: "https://api.upbit.com",
		Version: "/v1",
	})

	markets, err := upbit.GetMarkets(context.Background(), false)
	if err != nil {
		log.Println("Error fetching markets:", err)
		http.Error(writer, "Failed to fetch markets", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	marketsJSON, err := json.Marshal(markets)
	if err != nil {
		log.Println("Error marshaling markets to JSON:", err)
		http.Error(writer, "Failed to process markets data", http.StatusInternalServerError)
		return
	}
	if _, err := writer.Write(marketsJSON); err != nil {
		log.Println("Error writing response:", err)
		http.Error(writer, "Failed to write response", http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/markets", marketsHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "examples/markets/main.html")
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
