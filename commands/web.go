package commands

import (
	"aviasales/web"
	"github.com/urfave/cli"
)

var Web = cli.Command{
	Name:    "web",
	Aliases: []string{"c"},
	Usage:   "starts web server",
	Action: func(c *cli.Context) error {
		web.StartUp(c.GlobalString("conf"))
		return nil
	},
}
