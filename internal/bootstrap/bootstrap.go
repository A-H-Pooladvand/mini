package bootstrap

import (
	"kraken/internal/application"
	"kraken/internal/bootstrap/bootstraps"
)

var boots = []bootstraps.Bootstrapper{
	&bootstraps.Config{},
	&bootstraps.Env{},
}

func Boot() *application.App {
	app := &application.App{}

	for _, bootable := range boots {
		bootable.Boot(app)
	}

	return app
}
