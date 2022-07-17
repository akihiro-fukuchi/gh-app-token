package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	appId          = flag.String("app-id", "", "GitHub app id")
	installationId = flag.String("installation-id", "", "GitHub app installation id")
	privateKey     = flag.String("private-key", "", "Gitub app private key")

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
