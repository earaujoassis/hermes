package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/urfave/cli"

    "github.com/earaujoassis/hermes/server"
    "github.com/earaujoassis/hermes/server/web"
    "github.com/earaujoassis/hermes/server/tunnel"
)

func loadDotenv() {
    err := godotenv.Load()
    if err != nil {
        log.Printf("> The environment file (.env) doesn't exist; skipping\n")
    }
}

func main() {    app := cli.NewApp()
    app.Name = "Hermes"
    app.Usage = "An application for introspected tunnels to localhost"
    app.EnableBashCompletion = true
    app.Commands = []cli.Command{
        {
            Name:    "web",
            Usage:   "Serve the web application / REST API",
            Action:  func(c *cli.Context) error {
                loadDotenv()
                server.RepositoryStart()
                web.SetupWeb()
                return nil
            },
        },
        {
            Name:    "tunnel",
            Usage:   "Serve the tunnel server",
            Action:  func(c *cli.Context) error {
                tunnel.SetupTunnel()
                return nil
            },
        },
    }

    app.Run(os.Args)
}
