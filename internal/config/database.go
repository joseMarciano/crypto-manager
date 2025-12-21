package config

type Database struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

func databaseLoader(cfg *Configuration, props properties) {
	cfg.Database = Database{
		Host:         props["DB_HOST"],
		Port:         props["DB_PORT"],
		Username:     props["DB_USERNAME"],
		Password:     props["DB_PASSWORD"],
		DatabaseName: props["DB_NAME"],
	}
}
