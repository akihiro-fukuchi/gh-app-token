package main

import (
	"crypto/rsa"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type GitHubHTTPClient struct {
	client     *http.Client
	apiBaseURL string
}

func NewGitHubHTTPClient(c *http.Client, apiBaseURL string) *GitHubHTTPClient {
	return &GitHubHTTPClient{
		client:     c,
		apiBaseURL: apiBaseURL,
	}
}

func (c *GitHubHTTPClient) GetInstallationAccessToken(appCfg *appConfig) error {
	url := fmt.Sprintf("%s/app/installations/%s/access_tokens", c.apiBaseURL, appCfg.installationId)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", appCfg.jwtToken))
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("received non 2xx response status %q when fetching %v", resp.Status, req.URL)
	}
	return json.NewDecoder(resp.Body).Decode(&appCfg.installationAccessToken)
}

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
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		exitError(fmt.Errorf("missing flags"))
	}
	now := time.Now()
	appCfg, err := NewAppConfig(*appId, *installationId, *privateKey, now)
	if err != nil {
		exitError(err)
	}
	client := NewGitHubHTTPClient(&http.Client{}, apiBaseURL)
	if err = client.GetInstallationAccessToken(appCfg); err != nil {
		exitError(err)
	}

	fmt.Print(appCfg.installationAccessToken.Token)
}

func exitError(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
	os.Exit(1)
}
