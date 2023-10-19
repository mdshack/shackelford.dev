package services

type Services struct {
	GithubClientID     string `envconfig:"SERVICES_GITHUB_CLIENT_ID"`
	GithubClientSecret string `envconfig:"SERVICES_GITHUB_CLIENT_SECRET"`
}
