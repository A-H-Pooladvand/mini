package bootstraps

import "kraken/internal/application"

type Bootstrapper interface {
	Boot(app *application.App) *application.App
}
