package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/enorith/enocli/internal/handlers"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "init",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "dir",
						Value:   ".",
						Usage:   "directory of project",
						Aliases: []string{"d"},
					},
				},
				Usage: "init module-name [--dir=.]",
				Action: func(c *cli.Context) error {
					dir := c.String("dir")
					realDir, e := filepath.Abs(dir)
					if e != nil {
						fmt.Println(e)
						return e
					}
					module := c.Args().First()
					if len(module) < 1 {
						fmt.Printf("usage: enocli %s\n", c.Command.Usage)
						return fmt.Errorf("invalid module name")
					}

					path := filepath.Join(realDir, module)
					fmt.Printf("init enorith project at: %s\n", path)

					return handlers.InitCommand(path, module)
				},
			},
		},
	}

	app.Run(os.Args)
}
