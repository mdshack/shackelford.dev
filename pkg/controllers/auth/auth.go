package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mdshack/shackelford.dev/internal/container"
	"github.com/mdshack/shackelford.dev/pkg/templates/pages"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type AuthController struct {
	container *container.Container
	oauth     *oauth2.Config
}

type GithubUserResponse struct {
	ID int `json:"id"`
}

func New(c *container.Container) AuthController {

	return AuthController{
		container: c,
		oauth: &oauth2.Config{
			ClientID:     c.Config.Services.GithubClientID,
			ClientSecret: c.Config.Services.GithubClientSecret,
			Scopes:       []string{"read:user"},
			Endpoint:     github.Endpoint,
		},
	}
}

func (a *AuthController) Show(w http.ResponseWriter, r *http.Request) {
	pages.Login().Render(context.Background(), w)
}

func (a *AuthController) Redirect(w http.ResponseWriter, r *http.Request) {
	verifier := oauth2.GenerateVerifier()
	url := a.oauth.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func (a *AuthController) Callback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	verifier := oauth2.GenerateVerifier()

	code := r.URL.Query().Get("code")
	token, err := a.oauth.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}

	if !token.Valid() {
		log.Fatalf("Invalid token: %v", err)
		return
	}

	client := a.oauth.Client(ctx, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	var githubUser GithubUserResponse
	json.Unmarshal(body, &githubUser)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"github_id": githubUser.ID,
		"iat":       time.Now(),
		"exp":       time.Now().Add((24 * 7) * time.Hour),
	})

	s, err := jwtToken.SignedString([]byte(a.container.Config.Auth.AppKey))

	fmt.Println(s)
	fmt.Println(err)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
