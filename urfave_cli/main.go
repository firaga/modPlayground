package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

//https://cli.urfave.org/v2/getting-started/

func main() {
	app := &cli.App{
		Name:  "example",
		Usage: "Demonstrate combining short options",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "A",
				Aliases: []string{"a"},
				Usage:   "Enable option A",
			},
			&cli.BoolFlag{
				Name:    "B",
				Aliases: []string{"b"},
				Usage:   "Enable option B",
			},
			&cli.BoolFlag{
				Name:    "C",
				Aliases: []string{"c"},
				Usage:   "Enable option C",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("a") {
				fmt.Println("Option A is enabled")
			}
			if c.Bool("b") {
				fmt.Println("Option B is enabled")
			}
			if c.Bool("c") {
				fmt.Println("Option C is enabled")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
