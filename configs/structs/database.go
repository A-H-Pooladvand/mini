package structs

type DB struct {
	Mysql struct {
		Database string `env:"DB_DATABASE"`
		Host     string `default:"127.0.0.1" env:"DB_HOST"`
		Port     uint   `default:"3306" env:"DB_PORT"`
		Username string `default:"root" env:"DB_USERNAME"`
		Password uint   `default:"" env:"DB_PASSWORD"`
	}
}
