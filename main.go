package main

import (
    "log"

    "github.com/joho/godotenv"
    "github.com/earaujoassis/hermes/server"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Printf("> The environment file (.env) doesn't exist; skipping\n")
    }
}

func main() {
    server.RepositoryStart()
}
