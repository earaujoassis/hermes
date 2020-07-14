package server

import (
    "os"

    "github.com/urfave/cli"

    "github.com/earaujoassis/hermes/server/web"
    "github.com/earaujoassis/hermes/server/tunnel"
)

func Run() {
    app := cli.NewApp()
    app.Name = "Hermes"
    app.Usage = "An application for introspected tunnels to localhost"
    app.EnableBashCompletion = true
    app.Commands = []cli.Command{
        {
            Name:    "web",
            Usage:   "Serve the web application / REST API",
            Action:  func(c *cli.Context) error {
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
