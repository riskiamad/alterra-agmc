package config

import (
	"os"
	"regexp"

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
	projectName := regexp.MustCompile(`^(.*` + "mvc-testing" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	_ = godotenv.Load(string(rootPath) + `/.env`)

	var env EnvVar

	env.DBHost = os.Getenv("DBHOST")
	env.DBName = os.Getenv("DBNAME")
	env.DBUser = os.Getenv("DBUSER")
	env.DBPassword = os.Getenv("DBPASS")
	env.JWTSecret = os.Getenv("JWT_SECRET")

	return &env
}
