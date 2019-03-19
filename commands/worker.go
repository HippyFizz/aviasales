package commands

import (
	"aviasales/worker"
	"github.com/urfave/cli"
)

var Worker = cli.Command{
	Name:    "worker",
	Aliases: []string{"c"},
	Usage:   "run worker process",
	Action: func(c *cli.Context) error {
		worker.StartUp(c.GlobalString("conf"))
		return nil
	},
}
