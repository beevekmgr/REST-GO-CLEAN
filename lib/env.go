package lib

import "os"

// Env has environment stored
type Env struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	ServerPort string
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.DBUsername = os.Getenv("DB_USERNAME")
	env.DBPassword = os.Getenv("DB_PASSWORD")
	env.DBHost = os.Getenv("HOST_NAME")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBName = os.Getenv("HOST_NAME")
	env.DBName = os.Getenv("SERVER_PORT")

}
