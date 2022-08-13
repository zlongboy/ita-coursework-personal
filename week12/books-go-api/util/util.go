package util

import (
	"fmt"

	"github.com/joho/godotenv"
)

func getEnvMap() map[string]string {

	envMap, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println("Error reading .env to map")
	}
	return envMap
}
