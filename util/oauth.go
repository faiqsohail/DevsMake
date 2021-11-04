package util

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/go-github/github"
	"github.com/patrickmn/go-cache"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	config = oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     githuboauth.Endpoint,
		Scopes:       []string{"user:email", "repo:public_repo"},
	}

	userCache = cache.Cache{}
)

func InitUserCache() cache.Cache {
	var doOnce sync.Once

	doOnce.Do(func() {
		userCache = *cache.New(6*time.Hour, 5*time.Minute)
	})
	return userCache
}

func GetUserCache() cache.Cache {
	return userCache
}

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
	_, err := FetchAuthedUser(token)

	if err != nil {
		return false, err
	}
	return true, nil
}

func FetchAuthedUser(token string) (*github.User, error) {
	user, found := GetUserCache().Get(token)

	if !found {
		oauthClient := GetOAuthConfig().Client(context.TODO(), &oauth2.Token{AccessToken: token})
		client := github.NewClient(oauthClient)

		githubUser, _, err := client.Users.Get(context.TODO(), "")
		if err != nil {
			return nil, errors.New("unable to fetch the logged in user")
		}

		GetUserCache().Set(token, githubUser, cache.DefaultExpiration)
		return githubUser, nil
	}

	return user.(*github.User), nil
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
