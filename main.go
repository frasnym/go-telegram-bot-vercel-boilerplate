package main

import (
	"fmt"
	"net/http"

	handler "github.com/frasnym/go-telegram-bot-vercel-boilerplate/api"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/config"
)

func main() {
	cfg := config.GetConfig()

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/webhook", handler.WebhookHandler)
	fmt.Printf("Server is running on port %s...\n", cfg.Port)
	http.ListenAndServe(fmt.Sprint(":", cfg.Port), nil)
}
