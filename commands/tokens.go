package commands

import (
	"fmt"
	"github.com/gedex/ginsta/clients"
	"github.com/gedex/ginsta/utils"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	authURL      = "https://instagram.com/oauth/authorize"
	serveDomain  = "localhost:8080"
	redirectPath = "/redirect_uri"
)

var (
	cmdTokenGet = &Command{
		Callback: runTokenGet,
		Usage:    "token_get [-client-id CLIENT_ID] [-scope SCOPE]",
		Short:    "Get access_token",
		Long:     `Get access_token.`,
	}
	flagClientID, flagScope string
)

func init() {
	cmdTokenGet.Flag.StringVar(&flagClientID, "client-id", "", "CLIENT_ID")
	cmdTokenGet.Flag.StringVar(&flagScope, "scope", "", "SCOPE")
}

func home(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(authURL)
	if err != nil {
		utils.Check(err)
	}

	var clientID, scope string
	if flagClientID != "" {
		clientID = flagClientID
	} else {
		clientID = clients.DefaultClientID
	}
	if flagScope != "" {
		scope = flagScope
	} else {
		scope = clients.DefaultScope
	}

	q := u.Query()
	q.Set("client_id", clientID)
	q.Set("redirect_uri", "http://"+serveDomain+redirectPath)
	q.Set("response_type", "token")
	q.Set("scope", strings.Replace(scope, ",", " ", -1))
	u.RawQuery = q.Encode()

	http.Redirect(w, r, u.String(), http.StatusTemporaryRedirect)
}

func runTokenGet(r *Runner, cmd *Command, args []string) {
	http.HandleFunc("/", home)
	http.HandleFunc(redirectPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Your access_token is in URL fragment. Set it with '%s config access_token [ACCESS_TOKEN]'", clients.Name)
	})
	fmt.Println("Open localhost:8080 In your browser. You'll be redirected to Instagram authorization page")
	http.ListenAndServe(serveDomain, nil)

	os.Exit(0)
}
