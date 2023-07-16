package main

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
)

const VERSION = "1.0"

func main() {
	var (
		bindAddress string
		port        int
		region      string
		user        string
		password    string
		silent      bool
		healthPath  string
		limit       string
	)

	app := &cli.App{
		Name:    "kibana-proxy",
		Usage:   "kibana-proxy [options] <aws-es-cluster-endpoint>",
		Version: VERSION,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "bind-address",
				Usage:       "the ip address to bind to",
				DefaultText: "127.0.0.1",
				Aliases:     []string{"b"},
				Destination: &bindAddress,
			},
			&cli.IntFlag{
				Name:        "port",
				Usage:       "the port to bind to",
				DefaultText: "9200",
				Aliases:     []string{"p"},
				Destination: &port,
			},
			&cli.StringFlag{
				Name:        "region",
				Usage:       "the region of the Elasticsearch cluster",
				Aliases:     []string{"r"},
				Destination: &region,
			},
			&cli.StringFlag{
				Name:        "user",
				Usage:       "the username to access the proxy",
				DefaultText: os.Getenv("USER"),
				Aliases:     []string{"u"},
				Destination: &user,
			},
			&cli.StringFlag{
				Name:        "password",
				Usage:       "the password to access the proxy",
				Aliases:     []string{"a"},
				Destination: &password,
			},
			&cli.BoolFlag{
				Name:        "silent",
				Usage:       "remove figlet banner",
				DefaultText: "false",
				Aliases:     []string{"s"},
				Destination: &silent,
			},
			&cli.StringFlag{
				Name:        "health-path",
				Usage:       "URI path for health check",
				Aliases:     []string{"H"},
				Destination: &healthPath,
			},
			&cli.StringFlag{
				Name:        "limit",
				Usage:       "request limit",
				Aliases:     []string{"l"},
				DefaultText: "10000kb",
				Destination: &limit,
			},
		},
		Action: func(c *cli.Context) error {
			endpoint, err := setEndpoint(c)
			if err != nil {
				cli.ShowAppHelp(c)
				return cli.Exit(err, 1)
			}

			region, err := setRegion(region, endpoint)
			if err != nil {
				cli.ShowAppHelp(c)
				return cli.Exit(err, 1)
			}

			target := setTarget(endpoint)
			fmt.Printf("%v", target)

			run(bindAddress, port, region, user, password, silent, healthPath, limit, endpoint)

			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setEndpoint(c *cli.Context) (string, error) {
	endpoint := os.Getenv("ENDPOINT")

	if endpoint == "" {
		if c.NArg() > 0 && c.Args().Get(0) != "" {
			endpoint = c.Args().Get(0)
		} else {
			return "", errors.New("not set endpoint")
		}
	}

	return endpoint, nil
}

func setRegion(region, endpoint string) (string, error) {
	if region == "" {
		rep := regexp.MustCompile(`\.([^.]+)\.es\.amazonaws\.com\.?`)
		m := rep.FindAllStringSubmatch(endpoint, -1)
		if len(m) != 0 {
			region = m[0][1]
		} else {
			return "", errors.New("region cannot be parsed from endpoint address, either the endpoint must end in .<region>.es.amazonaws.com or --region should be provided as an argument")
		}
	}

	return region, nil
}

func setTarget(target string) string {
	rep := regexp.MustCompile(`^https?:\/\/`)
	if !rep.MatchString(target) {
		target = "https://" + target
	}

	return target
}

func run(
	bindAddress string,
	port int,
	region string,
	user string,
	password string,
	silent bool,
	healthPath string,
	limit string,
	endpoint string) {
	e := echo.New()

	if healthPath != "" {
		e.GET(healthPath, func(c echo.Context) error {
			c.Response().Header().Set("Content-Type", "text/plain")
			return c.JSON(http.StatusOK, "ok")
		})
	}

	if user != "" && password != "" {
		e.Use(middleware.BasicAuth(func(username, pass string, c echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(username), []byte(user)) == 1 &&
				subtle.ConstantTimeCompare([]byte(pass), []byte(password)) == 1 {
				return true, nil
			}
			return false, nil
		}))
	}

	url, err := url.Parse(endpoint)
	if err != nil {
		e.Logger.Fatal(err)
	}
	target := []*middleware.ProxyTarget{
		{
			URL: url,
		},
	}

	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(target)))
	e.Logger.Fatal(e.Start(bindAddress + ":" + strconv.Itoa(port)))
}
