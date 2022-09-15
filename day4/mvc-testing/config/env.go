package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Env = loadenv()

type EnvVar struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	JWTSecret  string
}

func loadenv() *EnvVar {
	_ = godotenv.Load()

	var env EnvVar

	env.DBHost = os.Getenv("DBHOST")
	env.DBName = os.Getenv("DBNAME")
	env.DBUser = os.Getenv("DBUSER")
	env.DBPassword = os.Getenv("DBPASSWORD")
	env.JWTSecret = os.Getenv("JWT_SECRET")

	return &env
}
