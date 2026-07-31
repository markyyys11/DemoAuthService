package config

import (
	"DemoAuthService/internal/logger"
	"os"

	"github.com/joho/godotenv"
)

func Default() conf {
	return def
}

var env conf
var local conf
var def conf

func Load() {
	env = new()
	def = env

	if err := godotenv.Load(); err != nil {
		logger.Hint("Can't load local ENV %s", err.Error())
		return
	}

	local = new()
	def = local
}

func new() conf {
	return conf{
		Debug:       getBool("DEBUG", false),
		PostgresURL: getStr("POSTGRES_URL", ""),
		Addr:        getStr("ADDR", ":8080"),
	}
}

type conf struct {
	Debug       bool
	PostgresURL string
	Addr        string
}

func getBool(key string, def bool) bool {
	if v, res := os.LookupEnv(key); res {
		return v == "true"
	} else {
		return def
	}
}

func getStr(key string, def string) string {
	if v, res := os.LookupEnv(key); res {
		return v
	} else {
		return def
	}
}
