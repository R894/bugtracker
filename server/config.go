package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func FetchEnvVars() {
	rootPath := filepath.Join(basepath, ".env")

	err := godotenv.Load(rootPath)
	if err != nil {
		log.Fatal("Error loading Environment variables: ", err)
	}
}

func GetBasePath() string {
	return basepath
}
