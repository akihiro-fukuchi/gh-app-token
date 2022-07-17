package main

import (
	"crypto/rsa"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type installationAccessToken struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

type appConfig struct {
	id                      string
	installationId          string
	privateKey              *rsa.PrivateKey
	jwtToken                string
	installationAccessToken *installationAccessToken
}

func NewAppConfig(appId, installationId, privateKey string, now time.Time) (*appConfig, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return nil, fmt.Errorf("private key is invalid. %+v", err)
	}
	c := &jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now.Add(-60 * time.Second)),
		ExpiresAt: jwt.NewNumericDate(now.Add(10 * 60 * time.Second)),
		Issuer:    appId,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, c).SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create jwt token. %+v", err)
	}

	return &appConfig{
		id:             appId,
		installationId: installationId,
		privateKey:     key,
		jwtToken:       token,
	}, nil
}

var (
	appId          = flag.String("app-id", "", "GitHub app id")
	installationId = flag.String("installation-id", "", "GitHub app installation id")
	privateKey     = flag.String("private-key", "", "GitHub app private key")

	apiBaseURL = "https://api.github.com"
)

func main() {
	flag.Parse()
	if *appId == "" || *installationId == "" || *privateKey == "" {
		flag.PrintDefaults()
		exitError(fmt.Errorf("missing flags"))
	}
	fmt.Print(apiBaseURL)
}

func exitError(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
	os.Exit(1)
}
