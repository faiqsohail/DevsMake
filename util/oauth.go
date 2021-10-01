package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	userInfoURL = "https://api.github.com/user"

	config = oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     githuboauth.Endpoint,
		Scopes:       []string{"user:email", "repo"},
	}
)

//GetOAuthConfig
func GetOAuthConfig() *oauth2.Config {
	return &config
}

func DoLogin(r *http.Request) middleware.Responder {
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			http.Redirect(w, r, config.AuthCodeURL(GenerateStateCookie(w)), http.StatusTemporaryRedirect)
		})
}

func DoCallback(r *http.Request) (*oauth2.Token, error) {
	authState, _ := r.Cookie("OAuthState")

	if r.URL.Query().Get("state") != authState.Value {
		return nil, fmt.Errorf("state did not match")
	}

	myClient := &http.Client{}

	parentContext := context.Background()
	ctx := oidc.ClientContext(parentContext, myClient)

	authCode := r.URL.Query().Get("code")

	oauth2Token, err := config.Exchange(ctx, authCode)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token")
	}
	return oauth2Token, nil
}

func IsAuthenticated(token string) (bool, error) {
	bearToken := "Bearer " + token
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}

	req.Header.Add("Authorization", bearToken)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("fail to get response: %v", err)
	}
	if resp.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}

func GenerateStateCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "OAuthState", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}
