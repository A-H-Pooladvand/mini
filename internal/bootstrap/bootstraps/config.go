package bootstraps

import (
	"github.com/jinzhu/configor"
	"kraken/internal/application"
	fileManager "kraken/pkg/file"
	configs "kraken/pkg/utils/config"
	"kraken/pkg/utils/structs"
	"log"
	"reflect"
	"strings"
)

type Config struct{}

var config = &configs.Config{}

// Boot Loads configuration files.
func (c *Config) Boot(app *application.App) *application.App {

	fields := structs.FieldNames(config)
	files := fileManager.Filenames("configs")

	for _, file := range files {
		filename := fileManager.Filename(file)
		filename = strings.ToLower(filename)

		for _, field := range fields {
			if strings.ToLower(field) == filename {
				r := reflect.ValueOf(config)
				f := reflect.Indirect(r).FieldByName(field)

				err := configor.Load(f.Addr().Interface(), "configs/"+file)

				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	app.Config = config

	return app
}
