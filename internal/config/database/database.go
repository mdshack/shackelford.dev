package database

type Database struct {
	Database string `envconfig:"DATABASE_NAME"`
	Host     string `envconfig:"DATABASE_HOST"`
	Username string `envconfig:"DATABASE_USERNAME"`
	Password string `envconfig:"DATABASE_PASSWORD"`
}
