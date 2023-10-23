package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateJWT(pemPath string, appID string) (string, error) {
	// Read the PEM file
	pemBytes, err := ioutil.ReadFile(pemPath)
	if err != nil {
		return "", err
	}

	// Parse the private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pemBytes)
	if err != nil {
		return "", err
	}

	// Define the token claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": int64(time.Now().Unix() - 60),       // Issued at time, 60 seconds in the past to account for clock drift
		"exp": int64(time.Now().Unix() + (10 * 60)), // Expiration time (10 minute maximum)
		"iss": appID,                                // GitHub App's identifier
	})

	// Sign the token
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	// Read environment variables
	pemPath := os.Getenv("PRIV_KEY_PATH")
	appID := os.Getenv("GH_APP_ID")

	// Validate environment variables
	if pemPath == "" {
		log.Fatal("Please set the PRIV_KEY_PATH environment variable.")
	}
	if appID == "" {
		log.Fatal("Please set the GH_APP_ID environment variable.")
	}

	jwt, err := generateJWT(pemPath, appID)
	if err != nil {
		log.Fatalf("Failed to generate JWT: %v", err)
	}

	log.Println("Generated JWT:", jwt)
}

