package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-faceit/internal/structs"
)

func LoadConfig() *structs.Config {
    port := os.Getenv("PORT")
    host := os.Getenv("HOST")

    if host == "" {
        host = "0.0.0.0"
    }

    if port == "" {
        port = "3000"
    }

    return &structs.Config{
        UserToken: os.Getenv("FACEIT_USER_TOKEN"),
        ApiToken:  os.Getenv("FACEIT_API_TOKEN"),
        Port:      port,
        Host:      host,
    }
}
