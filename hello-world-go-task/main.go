package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

func main() {

	serviceURL := os.Getenv("SVC_URL")
	tokenUrl := os.Getenv("TOKEN_URL")
	clientSecret := os.Getenv("CONSUMER_SECRET")
	clientId := os.Getenv("CONSUMER_KEY")
	var w http.ResponseWriter
	var clientCredsConfig = clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     tokenUrl,
	}

	client := clientCredsConfig.Client(context.Background())
	response, err := client.Get(serviceURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error making request: %v", err), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
			http.Error(w, fmt.Sprintf("Error reading response body: %v", err), http.StatusInternalServerError)
			return
	}
	
	fmt.Println(string(body))
}
