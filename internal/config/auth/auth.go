package auth

type Auth struct {
	AppKey  string `envconfig:"AUTH_APP_KEY"`
	AdminID string `envconfig:"AUTH_ADMIN_ID"`
}
