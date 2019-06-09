package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

const version = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "cli-test"
	app.Version = version
	app.Usage = "test cli tool"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "global, g",
			Value: "default-global",
			Usage: "global option",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "foo",
			Usage: "usage for foo",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "sub, s",
					Value: "default-sub",
					Usage: "sub option",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Printf("this is foo (global=%s, sub=%s)\n", c.GlobalString("global"), c.String("sub"))
				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
