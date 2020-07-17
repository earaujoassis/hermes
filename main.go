package main

import (
    "log"
    "os"
    "os/user"
    "path/filepath"

    "github.com/joho/godotenv"
    "github.com/urfave/cli"

    "github.com/earaujoassis/hermes/server"
    "github.com/earaujoassis/hermes/server/web"
    "github.com/earaujoassis/hermes/server/tunnel"
    "github.com/earaujoassis/hermes/client"
)

func loadDotenv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("> The environment file (.env) doesn't exist; skipping")
    }
}

func main() {
    var filepathConfig string

    app := cli.NewApp()
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
                loadDotenv()
                tunnel.SetupTunnel()
                return nil
            },
        },
        {
            Name:    "client",
            Usage:   "Client resposible for creating and retrieving HTTP messages",
            Flags:   []cli.Flag{
                cli.StringFlag{
                    Name:  "config, c",
                    Usage: "Load configuration from `FILE`",
                    Value: "",
                    Destination: &filepathConfig,
                },
            },
            Action:  func(c *cli.Context) error {
                if filepathConfig == "" {
                    usr, _ := user.Current()
                    filepathConfig, _ = filepath.Abs(filepath.Join(usr.HomeDir, ".hermes.config.json"))
                }
                client.SetupClient(filepathConfig)
                return nil
            },
        },
    }

    app.Run(os.Args)
}
