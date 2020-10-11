package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "short",
			Usage: "complete a task on the list",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "serve", Aliases: []string{"s"}},
				&cli.BoolFlag{Name: "option", Aliases: []string{"o"}},
				&cli.StringFlag{Name: "message", Aliases: []string{"m"}},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("serve:", c.Bool("serve"))
				fmt.Println("option:", c.Bool("option"))
				fmt.Println("message:", c.String("message"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
