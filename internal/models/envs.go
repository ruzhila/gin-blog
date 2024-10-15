package models

type Envs struct {
	Prefix         string `env:"PREFIX" comment:"Prefix for all routes"`
	ConsolePrefix  string `env:"CONSOLE" comment:"Prefix for console"`
	AuthPrefix     string `env:"AUTH_PREFIX" comment:"Prefix for auth"`
	Static         string `env:"STATIC" comment:"Prefix for static files"`
	ThemePath      string `env:"THEME_PATH" comment:"Path to the theme"`
	AllowRegistion bool   `env:"ALLOW_REGISTION" comment:"Allow registion"`
}

func GetEnvs() *Envs {
	return &Envs{
		Prefix:         "/",
		Static:         "/static",
		ConsolePrefix:  "/console",
		AuthPrefix:     "/auth",
		ThemePath:      "themes/default",
		AllowRegistion: true,
	}
}
