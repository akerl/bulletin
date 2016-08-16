package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "bulletin"
	app.Version = "0.0.1"
	app.Usage = "Service for exposing dynamic information from other services for collection and parsing"
	app.Action = setup

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host, H",
			Value: "::",
			Usage: "Bind to `HOST`",
		},
		cli.IntFlag{
			Name:  "port, P",
			Value: 2626,
			Usage: "Bind to `PORT`",
		},
		cli.StringFlag{
			Name:  "dir, d",
			Value: "/var/lib/bulletin",
			Usage: "Load data from `DIR`",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Print debug information",
		},
	}

	app.Run(os.Args)
}

func setup(c *cli.Context) error {
	fmt.Println("Hello world")
	return nil
}
