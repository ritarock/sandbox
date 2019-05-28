package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

var Version string = "0.1.0"

var Commands = []cli.Command{
	commandGet,
	commandPost,
}

var commandGet = cli.Command{
	Name:   "get",
	Usage:  "get method",
	Action: doGet,
}

var commandPost = cli.Command{
	Name:   "post",
	Usage:  "post method",
	Action: doPost,
}

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "cli_test"
	app.Usage = "cli test"
	app.Version = Version
	app.Commands = Commands
	return app
}

func doGet(c *cli.Context) error {
	url := c.Args().Get(0)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	fmt.Println(string(body))

	return nil
}

func doPost(c *cli.Context) error {
	fmt.Println("later")
	return nil
}
