package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-faceit/internal/structs"
)

func LoadConfig() *structs.Config {
    return &structs.Config{
        UserToken: os.Getenv("FACEIT_USER_TOKEN"),
        ApiToken:  os.Getenv("FACEIT_API_TOKEN"),
        Port:      os.Getenv("PORT"),
    }
}
