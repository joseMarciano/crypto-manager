package config

type Server struct {
	Port string
}

func serverLoader(cfg *Configuration, props properties) {
	cfg.Server = Server{
		Port: props["SERVER_PORT"],
	}
}
