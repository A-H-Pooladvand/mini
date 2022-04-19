package bootstraps

import (
	"github.com/joho/godotenv"
	"kraken/internal/application"
	"log"
)

type Env struct{}

// Boot loads the environment file.
func (c *Env) Boot(app *application.App) *application.App {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	return app
}
