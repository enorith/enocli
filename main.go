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
			initCommand(),
		},
		Usage: "Enorith cli tool",
	}

	app.Run(os.Args)
}

func initCommand() *cli.Command {
	return &cli.Command{
		Name: "init",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Value:   ".",
				Usage:   "directory of project",
				Aliases: []string{"d"},
			},
			&cli.StringFlag{
				Name:    "version",
				Value:   "",
				Usage:   "version of Enorith",
				Aliases: []string{"v"},
			},
		},
		Usage: "init [--dir=. --version=master] module-name",
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
			fmt.Printf("creating Enorith project at: %s\n", path)

			e = handlers.InitCommand(path, module, c.String("version"))

			if e != nil {
				fmt.Println(e)
				return e
			}

			return nil
		},
	}
}
