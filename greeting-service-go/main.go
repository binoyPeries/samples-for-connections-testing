package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ChatRequest struct {
	UserQuery string `json:"user_query"`
}

type ChatResponse struct {
	Response string      `json:"response"`
	Chart    interface{} `json:"chart,omitempty"`
}

type ToolRequest struct {
	UserQuery string `json:"user_query"`
}

type ToolResponse struct {
	SelectedTools []string `json:"selected_tools"`
}

func main() {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/health", healthCheckHandler)
	serverMux.HandleFunc("/chat", chatHandler)
	serverMux.HandleFunc("/tools", toolsHandler)

	serverPort := 9090
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: serverMux,
	}

	go func() {
		log.Printf("Starting server on port %d\n", serverPort)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP ListenAndServe error: %v", err)
		}
		log.Println("Server stopped serving new requests.")
	}()

	// Graceful shutdown
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Shutting down the server...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Shutdown complete.")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{
		Status:  "healthy",
		Message: "Application is running",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error writing health response: %v", err)
	}
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var chatReq ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&chatReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Dummy response logic
	chatResp := ChatResponse{
		Response: fmt.Sprintf("Analysis of query: %s", chatReq.UserQuery),
		Chart:    nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(chatResp); err != nil {
		log.Printf("Error writing chat response: %v", err)
	}
}

func toolsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var toolReq ToolRequest
	if err := json.NewDecoder(r.Body).Decode(&toolReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Dummy response logic
	toolResp := ToolResponse{
		SelectedTools: []string{"Tool1", "Tool2"},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(toolResp); err != nil {
		log.Printf("Error writing tools response: %v", err)
	}
}
