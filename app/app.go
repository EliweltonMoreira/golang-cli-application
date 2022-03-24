package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Get IPs and server names in the internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Get address IPs in the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			},
			Action: getIPs,
		},
	}

	return app
}

func getIPs(c *cli.Context) {
	host := c.String("host")

	IPs, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, IP := range IPs {
		fmt.Println(IP)
	}
}
